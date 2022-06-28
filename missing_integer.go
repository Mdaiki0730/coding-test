package main

import (
  "fmt"
  "sort"
)

func main() {
  sort.Ints(A)
  var missingInteger int = 1
  for _, v := range A {
      if v < 0 || v == missingInteger-1 {
          continue
      }
      if v == missingInteger {
          missingInteger += 1
      } else {
          break
      }
  }
  return missingInteger
}
