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
