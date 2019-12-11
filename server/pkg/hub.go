package pkg

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	u "luisfmelo/gochat/utils"
	"net/http"
)

// Upgrader specifies parameters for upgrading an HTTP connection to a WebSocket connection.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Hub struct {
	Clients map[string]Client
}

func (h *Hub) broadcast(emitter Client, msg string) {
	for _, c := range h.Clients {
		if c.Username == emitter.Username {
			continue
		}
		err := c.conn.WriteMessage(1, []byte(msg))
		if err != nil {
			log.Print(err)
		}
		u.LogChatMessage(emitter.Username, "everyone", msg)
	}
}

func (h *Hub) waitForMessage(c Client) {
	for {
		// read in a message
		messageType, p, err := c.conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		// print out that message for clarity
		fmt.Println(string(p), messageType)

		h.broadcast(c, string(p))
	}
}

func (h *Hub) registerClient(c Client) {
	if _, ok := h.Clients[c.Username]; !ok {
		h.Clients[c.Username] = c
		h.broadcast(c, fmt.Sprintf("%s joined.", c.Username))
	}
	c.Online = true
}

func (h *Hub) unregisterClient(c Client) {
	h.broadcast(c, fmt.Sprintf("%s left.", c.Username))
	c.Online = false
}
