package main

import (
  "fmt"
)

func main() {
  list := []int{2,7,6,4,1,8,5,9,6,2,4,1,5,8,3,5,2}
  windowSize := 5
  fmt.Println(Solution(list, windowSize))
}

func Solution(list []int, size int) []int {
  ret := []int{}
  for i := 0; i < len(list)-(size-1); i++ {
    window := list[i:i+size]
    ret = append(ret, Max(window))
  }
  return ret
}

func Max(list []int) int {
  max := list[0]
  for i := 1; i < len(list); i++ {
    if max < list[i] {
      max = list[i]
    }
  }
  return max
}
