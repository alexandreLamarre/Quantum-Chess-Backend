package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

//Client represents a new websocket connection, with the user's id, and the pool it communicates through.
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

//Message target object to decode the JSON messages of the general websocket connection.
type Message struct {
	Type    int    `json:type`
	Players string `json:"players"`
	ID      string `json:"id"`
	QueueId string `json:"queue"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		message := &Message{}
		err := c.Conn.ReadJSON(&message)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(message)
		c.Pool.Broadcast <- *message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
