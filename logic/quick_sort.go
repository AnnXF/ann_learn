package logic

func QuickSort[T int | uint](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)-1]
	less := make([]T, 0, len(arr)/2)
	greater := make([]T, 0, len(arr)/2)
	for _, num := range arr[:len(arr)-1] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	less = QuickSort(less)
	greater = QuickSort(greater)
	return append(append(less, pivot), greater...)
}
