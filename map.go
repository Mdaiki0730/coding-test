package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(WordCount("I am Go ! ! !"))
}

func WordCount(s string) map[string]int {
	devidedStrings := strings.Split(s, " ")
	wcm := make(map[string]int)
	for _, v := range devidedStrings {
      wcm[v] += 1
	}
	return wcm
}
