package mySlice

import "sort"

func MERGE(intervalArr [][]int) [][]int {

	if len(intervalArr) == 0 {
		return nil
	}

	sort.Slice(intervalArr, func(i, j int) bool {
		return intervalArr[i][0] < intervalArr[j][0]
	})
	// 初始化结果切片，放入第一个区间
	result := [][]int{intervalArr[0]}

	// 遍历剩余区间
	for i := 1; i < len(intervalArr); i++ {
		// 获取结果切片中最后一个区间
		last := result[len(result)-1]
		// 当前区间
		current := intervalArr[i]

		// 如果当前区间的起始位置小于等于最后一个区间的结束位置，说明有重叠
		if current[0] <= last[1] {
			// 合并区间，取两个区间结束位置的最大值
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			// 没有重叠，直接添加到结果切片
			result = append(result, current)
		}
	}
	return result
}
func SUMFORTWONUM(intArr []int, target int) (int, int) {
	k2 := 1
	for i := 0; i < len(intArr); i++ {
		for k2 < len(intArr) {
			if intArr[i]+intArr[k2] == target {
				return i, k2
			} else {
				k2++
			}
		}
	}
	return 0, 0

}
