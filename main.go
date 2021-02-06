package main

import (
  "fmt"

  "net/http"
  "github.com/alexandreLamarre/Quantum-Chess-Backend/pkg/websocket"
  "github.com/alexandreLamarre/Quantum-Chess-Backend/pkg/quantum"
)

var RUN bool = true

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
    fmt.Println("WebSocket Endpoint Hit")
    conn, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }

    client := &websocket.Client{
        ID: "john",
        Conn: conn,
        Pool: pool,
    }

    pool.Register <- client
    client.Read()
}

func parseHTMLRoute(path string) (act string, uuid string, ok bool) {
  return "", "", false
}

func doAction(action string, uuid string) bool{
  return true
}

func RootServer(w http.ResponseWriter, r *http.Request){
  if r.URL.Path == "/" {
    // do nothing for now
  } else if act, uuid, ok := parseHTMLRoute(r.URL.Path); ok{
    // have a function handle different game room actions
    res := doAction(act, uuid)
    if !res{
      http.Error(w, "Unexpected handler error", 400)
    }
  } else {
    http.Error(w, "Not found", http.StatusNotFound)
  }
}

func setupRoutes() {
    pool := websocket.NewPool()
    go pool.Start()

    http.HandleFunc("/", RootServer)

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(pool, w, r)
    })
}

func main() {
    fmt.Println("Quantum Chess App v0.01")
    fmt.Println(quantum.TestAllQuantum())
    setupRoutes()
    if(RUN) {
      http.ListenAndServe(":8080", nil)
    }
}
