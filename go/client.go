package main

import (
  "github.com/gorilla/websocket"
  )

type FindHandler func(string) (Handler, bool)

type Message struct {
  Name string `json:"name"`
  Data interface{} `json:"data"`
}

type Client struct {
  send chan Message
  socket *websocket.Conn
  findHandler FindHandler
}

func (client *Client) Read(){
  var message Message
  for {
    if err := client.socket.ReadJSON(&message); err != nil {
      break
    }
    if handler, found := client.findHandler(message.Name); found {
      handler(client, message.Data)
    }
  }
  client.socket.Close()
}
func (client *Client) Write(){
  for msg := range client.send {
    if err := client.socket.WriteJSON(msg); err != nil {
      break
    }
  }
  client.socket.Close()
}


func NewClient(socket *websocket.Conn, findHandler FindHandler) *Client{
  return &Client{
    send: make(chan Message),
    socket: socket,
    findHandler: findHandler,
  }
}



// func (client *Client) subscribeChannels() {
//   //TODO changefeed query rethinkdb
//   for {
//     time.Sleep(r())
//     client.send <- Message{"channel add", ""}
//   }
// }
// func (client *Client) subscribeMessages() {
//   //TODO changefeed query rethinkdb
//   for {
//     time.Sleep(r())
//     client.send <- Message{"message add", ""}
//   }
// }

// func r() time.Duration {
//   return time.Millisecond * time.Duration(rand.Intn(1000))
// }


// func main() {
//   client := NewClient()
//   go client.subscribeChannels()
//   go client.subscribeMessages()
//   client.write()
// }
