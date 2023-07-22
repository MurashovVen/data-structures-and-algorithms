package sort

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low, high int) {
	if high-low <= 1 {
		return
	}

	p := partition(arr, low, high)
	quickSort(arr, low, p-1)
	quickSort(arr, p+1, high)
}

func partition(arr []int, low, high int) (pivotNum int) {
	pivot := arr[high]
	pivotIdx := low

	for i := low; i < high; i++ {
		if arr[i] < pivot {
			arr[i], arr[pivotIdx] = arr[pivotIdx], arr[i]
			pivotIdx++
		}
	}

	arr[high], arr[pivotIdx] = arr[pivotIdx], arr[high]

	return pivotIdx
}
