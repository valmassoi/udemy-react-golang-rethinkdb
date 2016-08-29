package main

import (
  "net/http"
)

type Channel struct {
  Id string `json:"id"`
  Name string `json:"name"`
}

func main() {
  router := NewRouter()

  router.Handle("channel add", addChannel)

  http.Handle("/", router)
  http.ListenAndServe(":4000", nil)
}

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
