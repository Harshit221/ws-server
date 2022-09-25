package main

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

var clients = make([]*client, 0)

func handleNewConnection(ws *websocket.Conn) {
	log.Println("New connection")
	newClient := &client{
		id: uuid.NewString(),
		ws: ws,
	}
	clients = append(clients, newClient)
	newClient.serve()
}
