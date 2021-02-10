package websocket

import (
    "fmt"
    "log"
    "github.com/gorilla/websocket"
    "github.com/alexandreLamarre/Quantum-Chess-Backend/pkg/quantumchess"
)

type GameClient struct {
    ID   string
    Conn *websocket.Conn
    GamePool *GamePool
}

type GameMessage struct {
    Type int `json:"type"`   // 0 = player connected, 1 = board update, 2 = message, 3= opponent leave, 4= spectator join/leave
    Pid string `json:"pid"`
    Color int `json:"color"`
    GameStart bool `json:"start"`
    GameEnd bool `json:"end"`
    Message string `json:"message"`
    Move string `json:"move"`
    Board quantumchess.Board `json:"board"`
    Entanglements quantumchess.Entanglements `json:"entanglemets"`
    Pieces quantumchess.Pieces `json:"pieces"`
}

func (c *GameClient) GameRead() {
    defer func() {
        c.GamePool.Unregister <- c
        c.Conn.Close()
    }()

    for {
      message := &GameMessage{}
      err := c.Conn.ReadJSON(&message)
      if err != nil {
          log.Println(err)
          return
      }
      fmt.Println("\n mesage: \n", message)
      c.GamePool.Broadcast <- *message
      fmt.Printf("GAME Message Received: %+v\n", message)
    }
}
