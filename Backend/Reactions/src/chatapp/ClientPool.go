package main

import "fmt"

func newPool() *Pool {
	return &Pool{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan Message),

	}
}

func (h *Pool) run() {
	fmt.Println("Client Pool started..")
	for {
		select {
		case client := <-h.register:
			h.clients[client.Clientid] = client
			fmt.Println("Client Registered..")
			var mesg Message
			mesg.Type = "OnlineUsers"
			mesg.Users = Users
			//fmt.Println("debugger 2", client.Clientid);
		//	fmt.Println("message is ",mesg);
		//	var ok bool
		//	select {
		//	case h.broadcast <- mesg:
		//		ok = true
		//	default:
		//		ok = false
			//}
			//h.broadcast <- mesg
		//	fmt.Println("message is ",ok);

		case client := <-h.unregister:
			if _, ok := h.clients[client.Clientid]; ok {
				delete(h.clients, client.Clientid)
				close(client.send)
			}
			fmt.Println("Client Unregistered..")
		case message := <-h.broadcast:
				fmt.Println("Entered broadcast channel..",message)
			for client := range h.clients {
				h.clients[client].send <- message
				fmt.Println("Sending mesg to client channel..",client)

					//close(client.send)
					//delete(h.clients, client)
				}

			}
		}
	}

