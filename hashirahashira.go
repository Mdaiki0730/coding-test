package main

import(
  "bufio"
  "os"
  "math"
  "strconv"
  "strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

const (
	initBufSize = 1
	maxBufSize  = 100000 * 10000 + 99999
)

func init() {
	buf := make([]byte, initBufSize)
	scanner.Buffer(buf, maxBufSize)
}

func main() {
  defer writer.Flush()
  n, a := ScanStdIn()

  cost := CaliculateCost(n, a)

  writer.WriteString(strconv.Itoa(cost) + "\n")
}

func ScanStdIn() (int, []int) {
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
  dpt := make([]int, n)
  dpt[1] = int(math.Abs(float64(a[1] - a[0])))
  for i := 2; i < n; i++ {
    fromMinus2 := int(math.Abs(float64(a[i] - a[i-2]))) + dpt[i-2]
    fromMinus1 := int(math.Abs(float64(a[i] - a[i-1]))) + dpt[i-1]
    if fromMinus2 < fromMinus1 {
      dpt[i] = fromMinus2
    } else {
      dpt[i] = fromMinus1
    }
  }

  return dpt[n-1]
}
