package dfs

type BinaryNode struct {
	Value int
	Left  *BinaryNode
	Right *BinaryNode
}

func search(node *BinaryNode, needle int) bool {
	if node == nil {
		return false
	}

	if node.Value == needle {
		return true
	}

	if node.Value < needle {
		return search(node.Right, needle)
	}

	return search(node.Left, needle)
}

func dfs(root *BinaryNode, needle int) bool {
	return search(root, needle)
}
