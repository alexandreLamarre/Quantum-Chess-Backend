package websocket

import(
  "fmt"
  "math/rand"
)

var WHITE int = 0
var BLACK int = 1

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
}

func NewGamePool(id string) *GamePool {
    return &GamePool{
        ID: id,
        Register:   make(chan *GameClient),
        Unregister: make(chan *GameClient),
        Clients:    make(map[*GameClient]int),
        Broadcast:  make(chan GameMessage),
    }
}


func (pool *GamePool) StartGame() {
    for {
        select {
        case client := <-pool.Register:
            connected_id := client.ID

            fmt.Println("Size of Game Connection Pool: ", len(pool.Clients))
            fmt.Println("ID of client who joined", connected_id)

            if(len(pool.Clients) == 0){

              color:= rand.Intn(1);
              pool.Clients[client] = color;


            } else if(len(pool.Clients) == 1){
              var other_color int

              for _, color := range pool.Clients{
                if color == 1{
                  other_color = 0
                }else{
                  other_color = 1
                }
              }
              pool.Clients[client] = other_color

            }else{
              pool.Clients[client] = 2
            }

            //send messages on connect:
            var white_player string
            var black_player string
            if(len(pool.Clients) == 2){
              for player, color := range pool.Clients{
                if(color == 0){
                  white_player = player.ID
                } else if(color == 1){
                  black_player = player.ID
                }
              }
              for client, color := range pool.Clients {
                if(color == 0){
                  client.Conn.WriteJSON(GameMessage{Type:0, Pid: black_player, Color: color})
                } else if(color == 1){
                  client.Conn.WriteJSON(GameMessage{Type:0, Pid: white_player, Color: color})
                }
              }
            }
            break
        case client := <-pool.Unregister:
            delete(pool.Clients, client)
            msg := client.ID
            fmt.Println(msg);
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
