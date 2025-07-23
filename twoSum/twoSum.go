package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
loop:
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
				break loop
			}
		}
	}
	return result
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
