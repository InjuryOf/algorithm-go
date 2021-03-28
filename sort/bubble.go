package sort

// 冒泡排序思路
// 通过遍历将相邻两个元素进行两两比较，当一个元素大于右侧相邻元素时，交换它们的位置；当一个元素小于或等于右侧相邻元素时，位置不变
// 时间复杂度：O(n2)
// 空间复杂度：O(n2)
func bubble_sort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}


// 冒泡排序思路-优化版v1
// 优化思路：检查遍历信息，如果已经是有序的（没有进行排序移位），则直接跳出循环
// 时间复杂度：O(n2)
// 空间复杂度：O(n2)
func bubble_sort_opt_v1(array []int) []int {
	for i := 0; i < len(array); i++ {
		isSort := true
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				isSort = false
			}
		}
		if isSort{
			break
		}
	}
	return array
}

// 冒泡排序思路-优化版v2
// 优化思路：在排序过程中，确定无序区的边界，再往后就是有序区，不需要进行排序比较
// 时间复杂度：O(n2)
// 空间复杂度：O(n2)
func bubble_sort_opt_v2(array []int) []int {
	for i := 0; i < len(array); i++ {
		isSort := true
		// 无序数列的边界
		sortBorder := len(array) - 1
		for j := 0; j < sortBorder; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				isSort = false
				sortBorder = j
			}
		}
		if isSort{
			break
		}
	}
	return array
}
