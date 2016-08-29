package main

import (
  "net/http"
  r "github.com/dancannon/gorethink"
  "log"
)

func main() {
  session, err := r.Connect(r.ConnectOpts{
    Address: "localhost:28015",
    Database: "chat",
  })

  if err != nil {
    log.Panic(err.Error())
  }

  router := NewRouter(session)

  router.Handle("channel add", addChannel)
  router.Handle("channel subscribe", subscribeChannel)
  router.Handle("channel unsubscribe", unsubscribeChannel)

  router.Handle("user edit", editUser)
  router.Handle("user subscribe", subscribeUser)
  router.Handle("user unsubscribe", unsubscribeUser)

  router.Handle("message add", addChannelMessage)
  router.Handle("message subscribe", subscribeChannelMessage)
  router.Handle("message unsubscribe", unsubscribeChannelMessage)

  http.Handle("/", router)
  http.ListenAndServe(":4000", nil)
}


// type Channel struct {
//   Id string `json:"id" gorethink:"id,omitempty"`
//   Name string `json:"name" gorethink:"name"`
// }
//
// type User struct {
//   Id string `gorethink:"id,omitempty"`
//   Name string `gorethink:"name"`
// }

// func handler(w http.ResponseWriter, r *http.Request) {
//   // fmt.Fprintf(w, "hello from go")
//   // var socket *websocket.Conn
//   // var err error
//   // var socket, err = upgrader.Upgrade(w, r, nil)
//
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   for {
//     // msgType, msg, err := socket.ReadMessage()
//     // if err != nil {
//     //   fmt.Println(err)
//     //   return
//     // }
//     var inMessage Message
//     var outMessage Message
//     if err := socket.ReadJSON(&inMessage); err != nil {
//       fmt.Println(err)
//       break
//     }
//     fmt.Printf("%#v\n", inMessage)
//     switch inMessage.Name {
//     case "channel add":
//       err := addChannel(inMessage.Data)
//       if err != nil {
//         outMessage = Message{"error", err}
//         if err := socket.WriteJSON(outMessage); err != nil {
//           fmt.Println(err)
//           break
//         }
//       }
//     case "channel subscribe":
//       go subscribeChannel(socket)
//     }
//     // fmt.Println(string(msg))
//     // if err = socket.WriteMessage(msgType, msg); err != nil {
//     //   fmt.Println(err)
//     //   return
//     // }
//   }
// }
//do not communicate by sharing memory, share memory by communicating
// func addChannel(data interface{}) (error) {
//   var channel Channel
//
//   err := mapstructure.Decode(data, &channel)
//   if err != nil {
//     return err
//   }
//   channel.Id = "1"
//   // fmt.Printf("%#v\n", channel)
//   fmt.Println("add channel")
//   return nil
// }
//
// func subscribeChannel(socket *websocket.Conn) {
//   //simulate db query / changefeed
//   for {
//     time.Sleep(time.Second * 1)
//     message := Message{"channel add",
//       Channel{"1", "software support"}}
//     socket.WriteJSON(message)
//     fmt.Println("sent new channel")
//   }
// }
