package main

import (
  "fmt"
)

func main() {
  // nums := []int{1,3,5,7}
  nums := []int{3,2,4}
  fmt.Println(twoSum(nums, 7))
}

func twoSum(nums []int, target int) []int {
    prevMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        if index, isExist := prevMap[target-nums[i]]; isExist {
            return []int{index, i}
        }
        prevMap[nums[i]] = i
    }
    return []int{}
}
