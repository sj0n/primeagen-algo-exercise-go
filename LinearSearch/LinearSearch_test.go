package linearsearch

import "testing"

func TestLinearSearch(t *testing.T) {
	var numArray = []int{3, 17, 21, 1, 45, 13}
	t.Parallel()
	t.Run("Found number", func(t *testing.T) {
		var index = linearSearch(numArray, 21)

		if index != 2 {
			t.Errorf("Expected %d", 2)
			t.Errorf("Received %d", index)
		}
	})

	t.Run("Number not found", func(t *testing.T) {
		var index = linearSearch(numArray, 55)

		if index != -1 {
			t.Errorf("Expected %d", -1)
			t.Errorf("Received %d", index)
		}
	})
}
