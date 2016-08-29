package main

import (
  "github.com/gorilla/websocket"
  r "github.com/dancannon/gorethink"
  "log"
)

type FindHandler func(string) (Handler, bool)

type Client struct {
  send         chan Message
  socket       *websocket.Conn
  findHandler  FindHandler
  session      *r.Session
  stopChannels map[int]chan bool
  id           string
  userName     string
}

func (c *Client) NewStopChannel(stopKey int) chan bool {
  c.StopForKey(stopKey)
  stop := make(chan bool)
  c.stopChannels[stopKey] = stop
  return stop
}

func (c *Client) StopForKey(key int) {
  if ch, found := c.stopChannels[key]; found {
    ch <- true
    delete(c.stopChannels, key)
  }
}

func (client *Client) Read() {
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

func (client *Client) Write() {
  for msg := range client.send {
    if err := client.socket.WriteJSON(msg); err != nil {
      break
    }
  }
  client.socket.Close()
}

func (c *Client) Close() {
  for _, ch := range c.stopChannels {
    ch <- true
  }
  close(c.send)
  // delete user
  r.Table("user").Get(c.id).Delete().Exec(c.session)
}

func NewClient(socket *websocket.Conn, findHandler FindHandler,
  session *r.Session) *Client {
  var user User
  user.Name = "anon"
  res, err := r.Table("user").Insert(user).RunWrite(session)
  if err != nil {
    log.Println(err.Error())
  }
  var id string
  if len(res.GeneratedKeys) > 0 {
    id = res.GeneratedKeys[0]
  }
  return &Client{
    send:         make(chan Message),
    socket:       socket,
    findHandler:  findHandler,
    session:      session,
    stopChannels: make(map[int]chan bool),
    id:           id,
    userName:     user.Name,
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
