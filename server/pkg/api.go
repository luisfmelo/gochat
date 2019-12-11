package pkg

import (
	"log"
	u "luisfmelo/gochat/utils"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	u.JSON(w, http.StatusOK, "web socket chat is at /ws", nil)
}

func GetUsers(w http.ResponseWriter, r *http.Request, h Hub) {
	u.JSON(w, http.StatusOK, "Success", h.Clients)
}

func Websocket(w http.ResponseWriter, r *http.Request, h Hub) {
	// upgrade this connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	c := newClient(ws)
	h.registerClient(c)
	defer h.unregisterClient(c)

	// listen indefinitely for new messages coming through on our WebSocket connection
	h.waitForMessage(c)
}
