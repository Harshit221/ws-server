package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func init() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world!")
	})
}

func main() {
	http.HandleFunc("/ws",
		func(w http.ResponseWriter, req *http.Request) {
			s := websocket.Server{Handler: websocket.Handler(handleNewConnection)}
			s.ServeHTTP(w, req)
		})
	// http.Handle("/ws", websocket.Handler(handleNewConnection))
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
