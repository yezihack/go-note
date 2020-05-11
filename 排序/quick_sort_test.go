package 排序

import (
	"fmt"
	"testing"
)

func QuickSort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}
func quickSort(arr []int, left, right int) {
	pivot := arr[left]
	low, high := left, right
	for low < high {
		for low < high && arr[high] >= pivot {
			high --
		}
		arr[low] = arr[high]
		for low < high && arr[low] <= pivot {
			low ++
		}
		arr[high] = arr[low]

	}
	arr[low] = pivot
	if low - left > 1 {
		quickSort(arr, left, low -1)
	}
	if right - low > 1 {
		quickSort(arr, low +1, right)
	}
}
func ThreeMid(arr []int) int {
	start, end := 0, len(arr)-1
	mid := (start + end) >> 1
	first := arr[start]
	second := arr[mid]
	third := arr[end]
	if first > second && first < third {
		return first
	} else if first < second && second< third {
		return second
	} else {
		return third
	}
}
func TestQuickSort(t *testing.T) {
	fmt.Println(QuickSort([]int{8, 2, 1, 2, 9, 7}))
}