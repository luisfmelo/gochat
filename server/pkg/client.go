package pkg

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	Username string `json:"username"`
	Online   bool `json:"is_online"`
}

func (c *Client) Connect() {
	c.Online = true
}

func (c *Client) Disconnect() {
	c.Online = false
}

func newClient(ws *websocket.Conn) Client {
	return Client{
		conn:     ws,
		Username: randomdata.SillyName(),
	}
}
