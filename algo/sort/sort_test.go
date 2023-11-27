package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	nums := []int{3, 2, 5, 4, 8, 9, 1, 2}
	SelectionSort(nums)
	fmt.Println(nums)
}

func TestBubbleSort(t *testing.T) {
	nums := []int{3, 2, 5, 4, 8, 9, 1, 2}
	BubbleSort(nums)
	fmt.Println(nums)
}

func TestInsertionSort(t *testing.T) {
	nums := []int{3, 2, 5, 4, 8, 9, 1, 2}
	InsertionSort(nums)
	fmt.Println(nums)
}

func TestQuickSort1(t *testing.T) {
	nums := []int{3, 2, 5, 4, 8, 9, 1, 2}
	fmt.Println(QuickSort1(nums))
}

func TestQuickSort2(t *testing.T) {
	nums := []int{3, 2, 5, 4, 8, 9, 1, 2}
	QuickSort2(nums)
	fmt.Println(nums)
}

func TestMergeSort(t *testing.T) {
	nums := []int{3, 2, 5, 4, 8, 9, 1, 2}
	MergeSort(nums)
	fmt.Println(nums)
}
