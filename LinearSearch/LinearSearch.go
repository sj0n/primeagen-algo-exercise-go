package linearsearch

func linearSearch(array []int, num int) int {
	for i := 0; i < len(array); i++ {
		if array[i] == num {
			return i
		}
	}
	return -1
}