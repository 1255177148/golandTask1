package main

import "fmt"

/*
找出回文数，题中说了只有一个数字出现一次，其他的数字会出现两次。
那么可以使用异或，两个相同的数异或结果为0，而0和一个非0的数异或为非0的数
*/
func singleNumber(nums []int) int {
	num := 0
	for _, v := range nums {
		num ^= v
	}
	return num
}

func main() {
	nums := []int{1, 2, 2, 3, 4, 4, 3}
	fmt.Println(singleNumber(nums))
}
