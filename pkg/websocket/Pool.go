package websocket

import (
  "fmt"
  "strconv"
)

type Pool struct {
    Register   chan *Client
    Unregister chan *Client
    Clients    map[*Client]int
    Broadcast  chan Message
}

func NewPool() *Pool {
    return &Pool{
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
        Clients:    make(map[*Client]int),
        Broadcast:  make(chan Message),
    }
}

func (pool *Pool) Start() {
    for {
        select {
        case client := <-pool.Register:
            pool.Clients[client] = 0
            fmt.Println("Size of Connection Pool: ", len(pool.Clients), "\n")
            for client, _ := range pool.Clients {
                msg := strconv.Itoa(len(pool.Clients))
                fmt.Println(msg)
                client.Conn.WriteJSON(Message{Type: 0, Players: msg})
            }
            break
        case client := <-pool.Unregister:
            delete(pool.Clients, client)
            fmt.Println("Size of Connection Pool: ", len(pool.Clients), "\n")
            for client, _ := range pool.Clients {
              msg := strconv.Itoa(len(pool.Clients))
              fmt.Println(msg);
              client.Conn.WriteJSON(Message{Type: 0, Players: msg})
            }
            break
        case message := <-pool.Broadcast:
            fmt.Println("Sending message to all clients in Pool")
            for client, _ := range pool.Clients {
                if err := client.Conn.WriteJSON(message); err != nil {
                    fmt.Println(err)
                    return
                }
            }
        }
    }
}
