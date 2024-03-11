package comparebt

type BinaryNode struct {
	Value int
	Left  *BinaryNode
	Right *BinaryNode
}

func compare(a *BinaryNode, b *BinaryNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Value != b.Value {
		return false
	}

	return compare(a.Left, b.Left) && compare(a.Right, b.Right)
}
