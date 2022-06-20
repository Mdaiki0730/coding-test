package main

import (
  "fmt"
)

func main() {
  str := "feiawfje"
  // fmt.Printf("%T\n", str[:1])
  for _, v := range str {
    fmt.Println(string(v))
  }
}
