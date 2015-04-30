package main 

import (
  "fmt"
  "os"
)

func main() {
  // to get os.open function we give it a directory path instead of a filename
  dir, err := os.Open(".")
  if err != nil {
    return
  }
  defer dir.Close()

  fileInfos, err := dir.Readdir(-1)
  if err != nil {
    return
  }
  for _, fi := range fileInfos {
    fmt.Println(fi.Name())
  }

}