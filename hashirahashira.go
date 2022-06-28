package main

import(
  "fmt"
  "bufio"
  "os"
  "math"
  "strconv"
  "strings"
)

// https://atcoder.jp/contests/abc040/tasks/abc040_c

func main() {
  n, a := ScanStdIn()
  fmt.Println(n, a)

  cost := CaliculateCost(n, a)

  fmt.Printf("%v\n", cost)
}

func ScanStdIn() (int, []int) {
  scanner := bufio.NewScanner(os.Stdin)
  // scan n
  var n int
  scanner.Scan()
  intN, _ := strconv.Atoi(scanner.Text())
  n = intN

  // scan ints
  var a []int
  scanner.Scan()
  for _, v := range strings.Split(scanner.Text(), " ") {
    intA, _ := strconv.Atoi(v)
    a = append(a, intA)
  }
  return n, a
}

func CaliculateCost(n int, a []int) int {
  var cost int = 0
  for i := 0; i < len(a)-1; i++ {
    if i == len(a)-2 {
      cost += int(math.Abs(float64(a[len(a)-2] - a[len(a)-1])))
      continue
    }
    nextCost := int(math.Abs(float64(a[i+1] - a[i])))
    skipCost := int(math.Abs(float64(a[i+2] - a[i])))
    if nextCost > skipCost {
      cost += skipCost
      i++
    } else {
      cost += nextCost
    }
  }
  return cost
}
