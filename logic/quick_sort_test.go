package logic

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{1, 4, 2, 9, 9}
	res := QuickSort(arr)
	fmt.Println(res)
}
