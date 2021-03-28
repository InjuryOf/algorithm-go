package sort

// 鸡尾酒排序思路
// 优化思路：双向比较和交换（先左再右，交替执行），适用于特定条件的数据，可以减少排序的回合数，但是缺点是代码量增加，需要做奇数和偶数轮的判断
// 时间复杂度：O(n2)
// 空间复杂度：O(n2)
func cocktaol_sort(array []int) []int {
	for i := 0; i < len(array)/2; i++ {
		isSort := true
		// 奇数轮比较
		for j := i; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				isSort = false
			}
		}
		if isSort {
			break
		}

		isSort = true
		// 偶数轮比较
		for j := len(array) - 1 - i; j > i; j-- {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
				isSort = false
			}
		}
		if isSort {
			break
		}
	}
	return array
}
