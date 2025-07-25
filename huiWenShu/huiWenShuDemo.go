package main

import "fmt"

func isPalindromeNumber(a int) bool {
	if a < 0 {
		return false
	}
	t := a
	x := 0
	for t != 0 {
		x = (t % 10) + (x * 10)
		t = t / 10
	}
	return x == a
}

func main() {
	a := 1221
	fmt.Println(isPalindromeNumber(a))
}
