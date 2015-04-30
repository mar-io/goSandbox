package main 

import (
  "fmt"
  "crypto/sha1"
)

// cryptographic hash used because hard to reverse
// cryptographic hashes used to secure applications since it difficult to determine what made the hash
// we are using sha-1 here

func main() {
  // this computer 160 bit hash
  h := sha1.New()
  h.Write([]byte("test"))
  bs := h.Sum([]byte{})
  // we use a slice of 20 bytes
  fmt.Println(bs)

}