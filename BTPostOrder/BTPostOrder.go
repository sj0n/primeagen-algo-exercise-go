package btpostorder

type BinaryNode[T any] struct {
	left  *BinaryNode[T]
	right *BinaryNode[T]
	value T
}

func PostOrderSearch(head *BinaryNode[int]) []int {
	return walk(head, &[]int{})
}

func walk(current *BinaryNode[int], path *[]int) []int {
	if current == nil {
		return *path
	}

	
	walk(current.left, path)
	walk(current.right, path)
	*path = append(*path, current.value)

	return *path
}