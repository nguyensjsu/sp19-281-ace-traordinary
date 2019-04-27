package main

import (
	"github.com/gorilla/websocket"
	"time"
)

type Message struct {
	Userid    string `json:"userid"`
	Username string `json:"username"`
	Receiverid	string `json:"receiverid"`
	Message  string `json:"message"`
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
)