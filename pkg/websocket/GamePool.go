package websocket

import(
  "fmt"
  "math/rand"
)

var WHITE int = 0
var BLACK int = 1

type GameInfo struct{
    Ids []string
    Players []string
}

type Rooms struct{
  Privacy map[string]bool //false is public, true is private
  Games map[string]*GamePool
}

func NewRooms() *Rooms {
  return &Rooms{
    Privacy: make(map[string]bool),
    Games: make(map[string]*GamePool),
  }
}

type GamePool struct {
    ID string
    Register   chan *GameClient
    Unregister chan *GameClient
    Clients    map[*GameClient]int // maps to BLACK or WHITE, both cannot be the same obviously
    Broadcast  chan GameMessage
    //Timeout chan
    Start bool
    Over bool
}

func NewGamePool(id string) *GamePool {
    return &GamePool{
        ID: id,
        Register:   make(chan *GameClient),
        Unregister: make(chan *GameClient),
        Clients:    make(map[*GameClient]int),
        Broadcast:  make(chan GameMessage),
        Start: false,
        Over: false,
    }
}


func (pool *GamePool) StartGame() {
    for {
        select {
        case client := <-pool.Register:
            connectedId := client.ID

            fmt.Println("Size of Game Connection Pool: ", len(pool.Clients))
            fmt.Println("ID of client who joined", connectedId)

            if !pool.Start {
                assignInitialPlayers(pool, client)
            }

            //send messages on connect:

            break
        case client := <-pool.Unregister:
            delete(pool.Clients, client)
            msg := client.ID
            fmt.Println(msg)
            for client, _ := range pool.Clients {
              client.Conn.WriteJSON(GameMessage{Type: 3, Pid: msg})
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


func assignInitialPlayers(pool *GamePool, client *GameClient){
    if len(pool.Clients) == 0{

        color:= rand.Intn(1)
        pool.Clients[client] = color


    } else if len(pool.Clients) == 1{
        var otherColor int

        for _, color := range pool.Clients{
            if color == 1{
                otherColor = 0
            }else{
                otherColor = 1
            }
        }
        pool.Clients[client] = otherColor
        pool.Start = true

    }else{
        pool.Clients[client] = 2
    }

    // NOW SEND MESSAGES WHEN BOTH WHITE AND BLACK PLAYER HAVE CONNECTED
    var whitePlayer string
    var blackPlayer string
    if len(pool.Clients) == 2 {
        for player, color := range pool.Clients{
            if color == 0 {
                whitePlayer = player.ID
            } else if color == 1 {
                blackPlayer = player.ID
            }
        }
        for client, color := range pool.Clients {
            if color == 0 {
                client.Conn.WriteJSON(GameMessage{Type:0, Pid: blackPlayer, Color: color})
            } else if color == 1 {
                client.Conn.WriteJSON(GameMessage{Type:0, Pid: whitePlayer, Color: color})
            }
        }
    }

    //SEND MESSAGES WHEN SPECTATORS JOIN
}