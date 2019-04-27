package main

import "fmt"

func newPool() *Pool {
	return &Pool{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Pool) run() {
	fmt.Println("Client Pool started..")
	for {
		select {
		case client := <-h.register:
			h.clients[client.Clientid] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client.Clientid]; ok {
				delete(h.clients, client.Clientid)
				close(client.send)
			}
		}
	}
}
