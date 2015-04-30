package main 
// net/rpc remote procedure call exposes methods so they can be invoked over network
import (
  "fmt"
  "sync"
  "time"
)

func main() {
  // mutex or mutal exclusive lock locks a section of code to a single thread
  // this protects shared resources from non-atomic operations
  m := new(sync.Mutex)

  for i := 0; i < 10; i++ {
    go func(i int) {
      // when locked any other attempt to lock it will get blocked until it is unlocked
      // traditional multithreaded programming is difficult and mistakes are easy to make and hard to find
      // One of Go's biggest strenghts is the concurrency features its provides are easier to understand
      m.Lock()
      fmt.Println(i, "start")
      time.Sleep(time.Second)
      fmt.Println(i, "end")
      m.Unlock()
    }(i)
  }
}