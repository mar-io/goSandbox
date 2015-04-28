package main

import (
    "fmt"
    "time"
)

// ping infinite loop. "chan" stands for channel type. "string" is the type we are passing on the channel.
func pinger(c chan<- string) {
  //notice the channel direction above. chan<- means c can only be sent to. attempting to receive c will fail to compile.
  for i := 0; ; i++ {
    // left arrow operator is used to send and receive messages on channell
    // the usage means send "ping"
    c <- "ping"
  }
}

// using channel synchronizes goroutines. when pinger/ponger sends message it will wait until printer is ready to receive it
func ponger(c chan<- string) {
  for i := 0; ; i++ {
    c <- "pong"
  }
}

func printer(c <-chan string) {
  //notice the channel direction above. <-chan means c can only be received. if the left arrow operator is not specificed the channel is bi-directional.
  for {
    // the left arrow receives the ping message and stores to msg
    msg := <- c
    fmt.Println(msg)
    // could also write the above as "fmt.Prinln(<-c)"
    time.Sleep(time.Second * 1)
  }
}

func main() {
  var c chan string = make(chan string)

  go pinger(c)
  go ponger(c)
  go printer(c)

  var input string
  fmt.Scanln(&input)
}