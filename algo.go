package main

import (
  "fmt"
  "regexp"
)



func main() {
  req := "ooxox"
  fmt.Println(judgeWinner(req))
}

var oWin = regexp.MustCompile(`ooo`)
var xWin = regexp.MustCompile(`xxx`)

func judgeWinner(s string) (result string) {
  if oWin.MatchString(s) {
    return "o"
  } else if xWin.MatchString(s) {
    return "x"
  } else {
    return "draw"
  }
}
