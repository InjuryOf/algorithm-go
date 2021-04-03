package sort

func upAdjust(array []int) {
	childIndex := len(array) - 1
	parentIndex := (childIndex - 1) /2
	// 保存插入的叶子节点值，用于最后赋值
	temp := array[childIndex]

	for childIndex > 0 && temp < array[parentIndex] {
		array[childIndex] = array[parentIndex]
		childIndex = parentIndex
		parentIndex = (parentIndex - 1) / 2
	}
	array[childIndex] = temp
}

func downAdjust(array []int, parentIndex, length int) {
	temp := array[parentIndex]
	childIndex := 2 * parentIndex + 1
	for childIndex < length {
		if childIndex + 1 < length && array[childIndex + 1] < array[childIndex]{
			childIndex ++
		}

		// 如果父节点小于子节点，则直接跳出
		if temp <= array[childIndex]{
			break
		}

		array[parentIndex] = array[childIndex]
		parentIndex = childIndex
		childIndex = 2 * childIndex + 1
	}
	array[parentIndex] = temp
}

// 构建堆
func buildHeap(array []int) {
	for i := len(array) / 2; i >= 0; i-- {
		downAdjust(array, i, len(array))
	}
}

