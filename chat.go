package main

import (
	"encoding/json"
	"fmt"
)

type message struct {
	Username string `json:"username"`
	Id       string `json:"id"`
	Message  string `json:"message"`
}

var messages = make([]*message, 0)

func (m *message) sendTo(client *client) {
	data, _ := json.Marshal(m)
	client.ws.Write(data)
}

func (m *message) broadcast() {
	for _, c := range clients {
		m.sendTo(c)
	}
	messages = append(messages, m)
}

func systemMessage(msg string, a ...any) *message {
	return &message{
		Id:       "system",
		Username: "system",
		Message:  fmt.Sprintf(msg, a...),
	}
}
