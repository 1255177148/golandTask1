package main

func plus(digits []int) []int {
	var temp []int   // 暂时存放结果，不过是倒序的结果
	var result []int //最终结果
	// 标记是否有进位
	carryFlag := false
	for i := len(digits) - 1; i >= 0; i-- {
		n := digits[i]
		if i == len(digits)-1 {
			// 如果是个位，直接+1
			n += 1
		} else {
			// 如果不是个位，看当前位是否有进1
			if carryFlag {
				n += 1
			}
		}
		if n > 9 {
			// 表示需要向上进一
			carryFlag = true
			n = n % 10
		} else {
			carryFlag = false
		}
		temp = append(temp, n)
		if i == 0 && carryFlag {
			// 如果到了最高位，算完后还需要向上进1
			temp = append(temp, 1)
		}
	}
	for i := len(temp) - 1; i >= 0; i-- {
		result = append(result, temp[i])
	}
	return result
}
