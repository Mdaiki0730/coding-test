package main

import "fmt"

func main() {
    array := []int{1, 3, 6, 9, 100, 101}
    target := 100
    p := solve(array, target)
    fmt.Println("target index is ", p)
}

func solve(array []int, target int) int {
    arrayLen := len(array)
    start := 0
    end := arrayLen - 1
    var index int
    for {
        if end < start {
            return -1
        }
        index = (start + end) / 2

        if array[index] == target {
            return index
        }

        if array[index] < target {
            start = index + 1
        } else {
            end = index - 1
        }
    }
}
