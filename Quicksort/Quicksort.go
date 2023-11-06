package quicksort

func QuickSort(array []int, low int, high int) {
	if low >= high {
		return
	}

	pivotIndex := partition(array, low, high)

	QuickSort(array, low, pivotIndex-1)
	QuickSort(array, pivotIndex+1, high)
}

func partition(array []int, low int, high int) int {
	pivot := array[high]
	index := low - 1

	for i := low; i < high; i++ {
		if array[i] <= pivot {
			index++
			temp := array[i]
			array[i] = array[index]
			array[index] = temp
		}
	}

	index++
	array[high] = array[index]
	array[index] = pivot
	return index
}
