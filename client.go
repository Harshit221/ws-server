package main

import (
	"encoding/json"
	"log"

	"golang.org/x/net/websocket"
)

type client struct {
	username string
	id       string
	ws       *websocket.Conn
}

func (c *client) serve() {
	buffer := make([]byte, 512)

	for {
		n, err := c.ws.Read(buffer[0:])
		if err != nil {
			c.disconnect()
			systemMessage("%s has left the chat", c.username).broadcast()
			return
		}
		c.handleMessage(buffer[:n])
	}
}

func (c *client) disconnect() {
	for i, client := range clients {
		if c == client {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
	c.ws.Close()
}

func (c *client) handleMessage(buffer []byte) {
	msg := make(map[string]string)
	err := json.Unmarshal(buffer, &msg)
	if err != nil {
		log.Println(err)
		return
	}
	switch msg["type"] {
	case "initialise":
		c.username = msg["username"]
		msg, _ := json.Marshal(messages)
		systemMessage(string(msg)).sendTo(c)
		systemMessage("%s joined the chat", c.username).broadcast()
		break
	case "message":
		newMessage := message{
			Id:       c.id,
			Username: c.username,
			Message:  msg["message"],
		}
		newMessage.broadcast()
		break
	}
}
