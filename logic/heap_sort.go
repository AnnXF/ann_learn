package logic

import "fmt"

// 调整堆，使其满足最大堆性质
func heapify(arr []int, n, i int) {
	largest := i     // 初始化最大元素的位置为根节点
	left := 2*i + 1  // 左子节点的位置
	right := 2*i + 2 // 右子节点的位置

	// 如果左子节点存在且大于根节点，更新最大元素位置
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// 如果右子节点存在且大于根节点，更新最大元素位置
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大元素位置不是根节点位置，交换它们并继续调整堆
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// HeapSort 堆排序函数
func HeapSort(arr []int) {
	n := len(arr)

	// 建立最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	fmt.Println("大根堆：", arr)
	// 从堆中逐个取出元素并排序
	for i := n - 1; i >= 0; i-- {
		// 将当前最大元素移到末尾
		arr[0], arr[i] = arr[i], arr[0]
		// 重新调整堆
		heapify(arr, i, 0)
	}
}
