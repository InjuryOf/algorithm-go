package sort

func countSort(array []int) []int {
	max := array[0]
	// 获取最大元素
	for i := 0; i < len(array); i++ {
		if max < array[i] {
			max = array[i]
		}
	}
	// 遍历数组，填充统计数组
	countArray := make([]int, max+1, max+1)
	for i := 0; i < len(countArray); i++ {
		countArray[array[i]]++
	}

	// 遍历统计数组，输出结果
	index := 0
	sortArray := make([]int, len(array), len(array))
	for i := 0; i < len(countArray); i++ {
		for j := 0; j < countArray[i]; j++ {
			index++
			sortArray[index] = i
		}
	}
	return sortArray
}
