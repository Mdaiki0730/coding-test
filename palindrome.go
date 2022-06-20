package main

import (
  "fmt"
)

func main() {
  fmt.Println(isPalindrome(12321))
  fmt.Println(isPalindrome(21))
  fmt.Println(28/10)
}

func isPalindrome(x int) bool {
  var reversedNum int
  tmp := x
  for tmp > 0 {
      reversedNum = reversedNum*10 + tmp%10
      tmp =tmp/10
  }
  return x == reversedNum
}
