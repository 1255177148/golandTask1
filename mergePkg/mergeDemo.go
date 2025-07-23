package main

import "sort"

func merge(intervals [][]int) [][]int {
	/*
		可以先对区间数组按照区间的起始位置进行排序，
		然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，
		如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
	*/
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] // 按第一个元素排序
	})
	var result [][]int
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		/*
			当前数组和下一个数组比较，看是否要合并区间
		*/
		current := intervals[i]
		pre := result[len(result)-1]
		if pre[1] >= current[0] {
			if pre[1] < current[1] {
				// 需要合并区间，将当前数组的第二个值合并进去
				pre[1] = current[1]
				result[len(result)-1] = pre
			} else {
				continue
			}
		} else if pre[len(pre)-1] > current[1] {
			// 需要合并区间，但不需要当前区间，因为当前区间就在前一个区间内
			continue
		} else {
			// 不需要合并区间，将当前区间加入切片中
			result = append(result, current)
		}
	}
	return result
}
