package main

import (
	"github.com/gorilla/websocket"
	"time"
)

type Message struct {
	Type	string `json:"Type" bson:"Type,omitempty"`
	MessageId string `json:"MessageId" bson:"MessageId,omitempty"`
	Userid    string `json:"UserId" omitempty bson:"UserId,omitempty"`
	Username string `json:"Username" omitempty bson:"Username,omitempty"`
	Receiverid	string `json:"Receiverid" omitempty bson:"Receiverid,omitempty"`
	Message  string `json:"Message" omitempty bson:"Message,omitempty"`
	Time 	time.Time  `json:"Time" omitempty bson:"Time,omitempty"`
	//AddedUser string
	Status	string `json:"Status" omitempty bson:"Status,omitempty"`
	Lastupdated time.Time `json:"Lastupdated" omitempty bson:"Lastupdated,omitempty"`
	Users	map[string]bool  `json:"Users" omitempty`
	ConversationId	string	`json:"ConversationId" omitempty bson:"ConversationId,omitempty"`

}


type Client struct {
	Pool *Pool
	// The websocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	send chan Message
	//Buffered channel for incoming messages.
	recv chan []byte
	//Client id
	Clientid string
}

type Pool struct{

	clients map[string]*Client
	register chan *Client
	unregister chan *Client
	broadcast chan Message

}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512

	onlineUsersPeriod = 60 * time.Second
)
