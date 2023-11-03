package main

func BinarySearch(array []int, number int) int {
	start := 0
	end := len(array)
	
	for start < end {
		midIndex := (start + end) / 2

		if array[midIndex] == number {
			return midIndex
		} else if array[midIndex] > number {
			end = midIndex
		} else {
			start = midIndex + 1
		}
	}

	return -1
}
