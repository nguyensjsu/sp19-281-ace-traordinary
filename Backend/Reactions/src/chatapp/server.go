package main

import (
	//"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
	"github.com/satori/go.uuid"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var messagequeue = make(chan Message)           // message channel
var Users = make (map[string]bool)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()
	clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		messagequeue <- msg
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-messagequeue
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

//Read Go routine to process all buffered reads, stores messages, sends to write channel...
func (c *Client) readerProcess() {
	fmt.Println("in Reader");
	defer func() {
		c.Pool.unregister <- c
		//delete(Users, c.Clientid)
		var mesg Message
		mesg.Type = "OnlineUsers"
		mesg.Users = Users
		c.Pool.broadcast <- mesg
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))

	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	//fmt.Printf("Reader 2")
	for {
		var msg Message
		 err := c.conn.ReadJSON(&msg)
		if err != nil {
			panic(err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error: %v", err)
			}
			break
		}

		 if msg.Type == "Readreceipt" {
			seenTime := time.Now()
			 if convId, ok :=  updateSeenStatus(msg.Userid,msg.Receiverid,seenTime); ok{
				 if senderToNotify, ok := c.Pool.clients[msg.Userid]; ok {
					 var notif Message
					 notif.MessageId = msg.MessageId
					 notif.Type= "Notification"
					 notif.Userid = msg.Receiverid
					 notif.Receiverid = msg.Userid
					 notif.Status = "Seen"
					 notif.Time = seenTime
					 notif.ConversationId = convId
					 senderToNotify.notifications <- notif
					 fmt.Println("Sending sent notify, Seen")
				 }
			 }
		 }else {
			 fmt.Println("Message ", msg.Receiverid)
			 // msg.MessageId,_ := uuid.NewV4()
			 uid := uuid.NewV4()
			 msg.MessageId = uid.String()
			 msg.Status = "Sent"
			 msg.Lastupdated = time.Now()
			 msg.ConversationId = getConversationId(msg.Userid, msg.Receiverid)
			 if _, ok := storeMessage(msg); ok {
				 var notif Message
				 notif.MessageId = msg.MessageId
				 notif.Type = "Notification"
				 notif.Userid = msg.Userid
				 notif.Receiverid = msg.Receiverid
				 notif.Status = "Sent"
				 notif.Time = msg.Lastupdated
				 notif.ConversationId = msg.ConversationId
				 fmt.Println("Sending sent notify, Sent")
				 select{
					 case c.notifications <- notif:
						 fmt.Println("Sent to notif channel")
				 default:

				 }
				// c.notifications <- notif
				 fmt.Println("Sending sent notify2")
			 }
			 fmt.Println("Message ", msg.Receiverid)

			 if recvClient, ok := c.Pool.clients[msg.Receiverid]; ok {

				 msg.Type = "Text"
				 recvClient.send <- msg
			 } else {
				 //store in uread table..
				 storeUnreadMessage(msg)

			 }
		 }

	}
}

func (c *Client) writerProcess() {
	fmt.Println("in writer");
	ticker := time.NewTicker(pingPeriod)
	//tickerUsers := time.NewTicker(onlineUsersPeriod)
	defer func() {
		ticker.Stop()
		//tickerUsers.Stop()


		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			fmt.Println("in writer message is ",message.Type)
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.conn.NextWriter(websocket.TextMessage)

			//w.WriteJSON(message)
			fmt.Println("writer 1", message.Type)
			error := c.conn.WriteJSON(message)
			if error != nil {
				panic(error);
			}
			if message.Type == "Text" {
				updated_time := time.Now()
				if _, ok := updateConversations(message.MessageId, "Delivered", updated_time); ok {
					if sender, ok := c.Pool.clients[message.Userid]; ok {
						var notif Message
						notif.MessageId = message.MessageId
						notif.Type = "Notification"
						notif.Userid = message.Userid
						notif.Receiverid = message.Receiverid
						notif.Status = "Delivered"
						notif.Time = updated_time
						notif.ConversationId = message.ConversationId
						select {
							case sender.notifications <- notif:
						default:

						}
					}
				}
			}
			//c.conn.WriteMessage(websocket.TextMessage, message)
		case notifMessage, ok := <-c.notifications:

			fmt.Println("in NOTIFICATION is ",notifMessage.Type)
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.conn.NextWriter(websocket.TextMessage)

			//w.WriteJSON(message)
			fmt.Println("writer 1", notifMessage.Type)
			error := c.conn.WriteJSON(notifMessage)
			if error != nil {
				panic(error);
			}


		case <-ticker.C:
			fmt.Println("in ticker ",ticker)
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Println("writer ping error")
				panic(err)
				return
			}


		}
	}
}

func serveWs(Pool *Pool, w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var id string = params["Id"]
fmt.Println("request is ",r);
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{Pool: Pool, conn: conn, send: make(chan Message), notifications: make(chan Message), Clientid:id}
	fmt.Println("client is ",client);
	client.Pool.register <- client
	Users[client.Clientid] = true
	go client.writerProcess()
	go client.readerProcess()

	Users[client.Clientid] = true
	var mesg Message
	mesg.Type = "OnlineUsers"
	mesg.Users = Users

	client.Pool.broadcast <- mesg
	if messages ,ok := readUnreadMessages(client.Clientid); ok{

		for _, oldmesg := range messages {

			client.send <- oldmesg

		}



	}
	if err,ok :=removeUnreadMessages(client.Clientid); !ok{
		fmt.Println(err)
	}

	//check unread table and handle

}


