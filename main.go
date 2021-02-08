package main

import (
  "fmt"
  "strings"
  "net/http"
  "github.com/alexandreLamarre/Quantum-Chess-Backend/pkg/websocket"
  "github.com/alexandreLamarre/Quantum-Chess-Backend/pkg/quantum"
  "encoding/json"
  "strconv"
)


var RUN bool = true
var rooms *websocket.Rooms = websocket.NewRooms()

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
    //fmt.Println("WebSocket Endpoint Hit")
    conn, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }
    //fmt.Println("Request url", r.URL.Path)
    client := &websocket.Client{
        ID: string(r.URL.Path[4:]),
        Conn: conn,
        Pool: pool,
    }

    pool.Register <- client
    client.Read()
}

func serveGame(rooms *websocket.Rooms, w http.ResponseWriter, r *http.Request){

  fmt.Println("Game Request url", r.URL.Path)

  //create New Game Socket
  gid, cid := parseGameURL(string(r.URL.Path))
  fmt.Println("Client id: ", cid, "\n Game id: ", gid)

  var gamePool *websocket.GamePool
  if rooms.Games[gid] == nil{
    http.Error(w, "Game :" +gid, http.StatusNotFound);
    return
  } else{
    fmt.Println("Joining new Game", gid)
    gamePool = rooms.Games[gid]
  }


  fmt.Println("Game Websocket Endpoint Hit")
  conn, err := websocket.Upgrade(w,r)
  if err != nil {
    fmt.Fprintf(w, "%+v\n", err)
  }

  gameClient := &websocket.GameClient{
    ID: cid,
    Conn: conn,
    GamePool: gamePool,
  }

  gamePool.Register <- gameClient
  gameClient.GameRead()
}

func parseGameURL(url string) (string, string){

  s := strings.Split(url, "/"); // /Game / gid/ cid
  fmt.Println("url parse into", s)
  return s[2], s[3]
}


func serveList(rooms *websocket.Rooms, w http.ResponseWriter, r *http.Request){
  var ids []string
  for id, _ := range(rooms.Games){
    if(rooms.Privacy[id] == false){
      ids = append(ids, id);
      fmt.Println(id)
    }
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200);
  json.NewEncoder(w).Encode(ids)
}


func serveCreateGame(rooms *websocket.Rooms, w http.ResponseWriter, r *http.Request){
  gid, privacy := parseCreateURL(r.URL.Path)
  fmt.Println("Creating new Game", gid)
  if(rooms.Games[gid] != nil){
    fmt.Println("Game already exists")
    w.WriteHeader(409)
    return
  }
  w.WriteHeader(200)
  gamePool := websocket.NewGamePool(gid)
  rooms.Privacy[gid] = privacy
  rooms.Games[gid] = gamePool
  go gamePool.StartGame()
}

func parseCreateURL(url string) (string, bool) {
  s := strings.Split(url, "/")
  privacy, err := strconv.ParseBool(s[3])
  if(err != nil){
    fmt.Println("Parse bool error", err);
  }
  fmt.Println("created game with privacy", privacy)
  return s[2], privacy
}

func setupRoutes() {
    pool := websocket.NewPool()
    go pool.Start()



    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      if strings.HasPrefix(r.URL.Path,"/ws") {
        serveWs(pool, w, r)
      } else if strings.HasPrefix(r.URL.Path, "/game"){
        fmt.Println("Serving game websocket request")
        serveGame(rooms, w, r)
      }else if strings.HasPrefix(r.URL.Path, "/create"){
        fmt.Println("\n creating new Game")
        serveCreateGame(rooms, w, r)
      }else if strings.HasPrefix(r.URL.Path, "/listgames"){
        fmt.Println("\n fetching games in progress")
        serveList(rooms, w, r)
      }
    })
}

func main() {
    fmt.Println("Quantum Chess App v0.01")
    quantum.TestAllQuantum()
    setupRoutes()
    if(RUN) {
      http.ListenAndServe(":8080", nil)
    }
}
