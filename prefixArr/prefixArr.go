package main

import "fmt"

/*
获取一组字符串数组中的最长公共前缀
*/
func checkPrefix(arr []string) string {
	if len(arr) == 1 {
		return arr[0]
	}
	prefix := arr[0]
	var result []byte
	for i := 1; i < len(arr); i++ {
		var b []byte
		c := arr[i]
		for j := 0; j < len(prefix); j++ {
			if j >= len(c) {
				break
			}
			if prefix[j] == c[j] {
				b = append(b, prefix[j])
			} else {
				break
			}
		}
		prefix = string(b)
		result = b
	}
	return string(result)
}

func main() {
	arr := []string{"flower", "flow", "flight"}
	fmt.Println(arr)
}
