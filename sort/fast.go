package sort

// 快速排序思路-双边循环
// 先获取基准元素位置（可使用随机获取），然后采用双边循环法或单边循环法进行数据交互
// 时间复杂度：O(nlogn)
// 空间复杂度：
func bFastSort(array []int, startIndex int, endIndex int) {
	if startIndex >= endIndex {
		return
	}

	// 获取基准元素位置
	pivotIndex := bilateralPartition(array, startIndex, endIndex)

	// 根据基准元素，分为两部分进行递归排序
	bFastSort(array, startIndex, pivotIndex-1)
	bFastSort(array, startIndex+1, pivotIndex)
}

// 快速排序思路-单边循环
// 先获取基准元素（pivot）位置（可使用随机获取）：
//   1、如果遍历元素大于基准元素，就继续往后遍历
//   2、如果遍历元素小于基准元素，则做两件事情，第一把mark指针右移一位，因为小于pivot的区域边界增大了1位，
//  第二让最新遍历到的元素和mark指针所在位置的元素交换位置，因为最新遍历的元素归属于小于pivot的区域
// 时间复杂度：O(nlogn)
// 空间复杂度：
func sFastSort(array []int, startIndex int, endIndex int) {
	if startIndex >= endIndex {
		return
	}

	// 获取基准元素位置
	pivotIndex := unilateralPartition(array, startIndex, endIndex)

	// 根据基准元素，分为两部分进行递归排序
	sFastSort(array, startIndex, pivotIndex-1)
	bFastSort(array, startIndex+1, pivotIndex)
}

// 分治（双边循环法）
func bilateralPartition(array []int, startIndex int, endIndex int) int {
	// 取第一个位置作为基准元素（也可以随机选择位置）
	pivot := array[startIndex]
	left := startIndex
	right := endIndex
	for left != right {

		// right 指针比较并左移
		for left < right && array[right] > pivot {
			right--
		}

		// left 指针比较并右移
		for left < right && array[left] <= pivot {
			left++
		}

		// 交换left和right指针所指向的元素
		if left < right {
			array[left], array[right] = array[right], array[left]
		}
	}
	array[startIndex], array[left] = array[left], pivot
	return left
}

// 分治（单边循环法）
func unilateralPartition(array []int, startIndex int, endIndex int) int {
	// 取第一个位置作为基准元素（也可以随机选择位置）
	pivot := array[startIndex]
	mark := startIndex
	for i := startIndex + 1; i <= endIndex; i++ {
		if array[i] < pivot{
			mark ++
			array[mark], array[i] = array[i], array[mark]
		}
	}

	array[startIndex], array[mark] = array[mark], pivot
	return mark
}
