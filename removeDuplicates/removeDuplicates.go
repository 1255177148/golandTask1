package main

import "fmt"

func removeDup(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	num := 1
	for i := 1; i < len(nums); i++ {
		/* 下标从1开始，如果前一个数和当前数不一样，则数组的当前下标和num值是一样的，
		如果值是一样的，则num值不变，i值变大1，一直找到和(num-1)下标不一样的值的下标i为止，然后将
		下标为i的不一样的值赋给num下标
		*/
		if nums[i] != nums[i-1] {
			nums[num] = nums[i]
			num++
		}
	}
	return num
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 6, 7, 8, 8, 9, 9, 9}
	result := removeDup(nums)
	fmt.Println(result)
	fmt.Println(nums)
}
