package main 

import (
  "fmt"
  "os"
  "path/filepath"
)

func main() {
  // walk function will recursively read the folders contents and all the sub folder and sub-sub folder
  // . is the current directory the script is running in
  filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    fmt.Println(path)
    return nil
  })
}