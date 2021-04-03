package sort

import "testing"

func TestBubbleSort(t *testing.T) {

	// 方式1
	array := []int{5, 8, 6, 3, 9, 2, 1, 7}
	t.Logf("冒泡排序前：%v\n", array)
	t.Logf("冒泡排序后结果：%v\n", bubble_sort(array))

	// 方式2
	array = []int{5, 8, 6, 3, 9, 2, 1, 7}
	t.Logf("优化版（v1）冒泡排序前：%v\n", array)
	t.Logf("优化版（v1）冒泡排序后结果：%v\n", bubble_sort_opt_v1(array))

	// 方式3
	arrayNew := []int{3, 4, 2, 1, 5, 6, 7, 8}
	t.Logf("优化版（v2）冒泡排序前：%v\n", arrayNew)
	t.Logf("优化版（v2）冒泡排序后结果：%v\n", bubble_sort_opt_v2(arrayNew))
}

func TestCocktail(t *testing.T) {
	array := []int{2, 3, 4, 5, 6, 7, 8, 1}
	t.Logf("鸡尾酒排序前：%v\n", array)
	t.Logf("鸡尾酒排序后结果：%v\n", cocktaol_sort(array))
}

func TestFastSort(t *testing.T) {
	// 双边循环法
	array := []int{4, 1, 6, 5, 3, 2, 8, 7}
	t.Logf("双边快速排序前：%v\n", array)
	bFastSort(array, 0, len(array)-1)
	t.Logf("双边快速排序后结果：%v\n", array)

	// 单边循环法
	array = []int{4, 7, 3, 5, 6, 2, 8, 1}
	t.Logf("单边快速排序前：%v\n", array)
	sFastSort(array, 0, len(array)-1)
	t.Logf("单边快速排序后结果：%v\n", array)
}
