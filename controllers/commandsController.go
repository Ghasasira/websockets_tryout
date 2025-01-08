package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// ReadBufferSize:  1024,
	// WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connections []*websocket.Conn

func CommandsController(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	connections = append(connections, conn)
	go handleConnection(conn)
}

func handleConnection(conn *websocket.Conn) {
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		log.Printf("message sent is: %s", message)

		for _, c := range connections {
			if c != conn {
				err = c.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Println("Error writing message:", err)
					break
				}
				log.Println("wow")
			}
		}
	}
}
