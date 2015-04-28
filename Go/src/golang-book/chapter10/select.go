package main

import (
    "fmt"
    "time"
)

func main() {
  // specifying the "1" creates a buffered channel with a capacity of 1
  // by default channels are synchronous, each side of a channel waiting for the other to be ready.
  // buffered channel is asynchronous sending and receiving messages will not wait unless channel is full
  c1 := make(chan string, 1)
  c2 := make(chan string, 1)

  go func() {
    for {
      c1 <- "from 1"
      time.Sleep(time.Second * 2)
    }
  }()

  go func() {
    for {
      c2 <- "from 2"
      time.Sleep(time.Second * 3)
    }
  }()

  go func() {
    for {
      // select picks the first channel that is ready and receives from it(or sends to it)
      // if more than one channel is ready it randomly picks which one to receive from
      // if no channels are ready the statement blocks until one is available
      select {
      case msg1 := <- c1:
        fmt.Println(msg1)
      case msg2 := <- c2:
        fmt.Println(msg2)
      // time.After create a channel and after the given duration will send the current time on it
      case <- time.After(time.Second):
        fmt.Println("timeout")
      // default case happens immediately if none of the channels are ready
      //default:
      //  fmt.Println("nothing ready")
      }
    }
  }()

  var input string
  fmt.Scanln(&input)
}