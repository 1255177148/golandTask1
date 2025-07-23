package main

import "fmt"

/*
给定一个字符串，判断字符串是否有效。
可以使用切片模拟Java中堆栈的形式
*/
func effectiveStr(s string) bool {
	n := len(s)
	var stack []byte
	for i := 0; i < n; i++ {
		if '(' == s[i] || '[' == s[i] || '{' == s[i] {
			// 如果是左侧括号，就入栈
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 {
				// 否则就弹出栈最顶上的值来比较
				b := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if ')' == s[i] && '(' != b {
					return false
				}
				if ']' == s[i] && '[' != b {
					return false
				}
				if '}' == s[i] && '{' != b {
					return false
				}
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	s := "{([])}"
	fmt.Println(effectiveStr(s))
}
