package logic

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("原始数组：", arr)
	HeapSort(arr)
	fmt.Println("排序后数组：", arr)
}
