package structures

type TreeNode[T any] struct {
	Val   *T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func TransformToPointer[T interface{}](arr []interface{}) (res []*T) {
	for _, v := range arr {
		if v == nil {
			res = append(res, nil)
		} else {
			if v, ok := v.(T); ok {
				res = append(res, &v)
			}
		}
	}
	return res
}

func TransformToVal[T interface{}](arr []*T) (res []any) {
	for _, v := range arr {
		if v == nil {
			res = append(res, nil)
		} else {
			res = append(res, *v)
		}
	}
	return res
}

// 从广度优先遍历序列化的数组还原树  construct tree from a tree's bfs array data
func ConstructTree[T any](arr []*T) (root *TreeNode[T]) {
	if len(arr) == 0 {
		return &TreeNode[T]{}
	}
	i := 0
	stack := []*TreeNode[T]{{Val: arr[i]}}
	root = stack[0]

	for i+2 < len(arr) && len(stack) > 0 {
		curNode := stack[0]
		stack = stack[1:]
		if curNode.Val != nil {
			curNode.Left = &TreeNode[T]{Val: arr[i+1]}
			curNode.Right = &TreeNode[T]{Val: arr[i+2]}
			i = i + 2
			stack = append(stack, curNode.Left, curNode.Right)
		}
	}

	return root
}

func SerializeTree[T any](root *TreeNode[T]) (arr []*T) { // bfs serialize tree
	stack := []*TreeNode[T]{root}

	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]

		if node == nil {
			arr = append(arr, nil)
			continue
		}
		arr = append(arr, node.Val)

		if node.Left == nil && node.Right == nil {
			continue
		}

		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}

	return arr
}

func InOrderTraverseTree[T any](node *TreeNode[T]) []*T {
	var val_list []*T
	var inOrderTraverse func(node *TreeNode[T])

	inOrderTraverse = func(node *TreeNode[T]) {
		if node == nil {
			return
		}
		inOrderTraverse(node.Left)
		val_list = append(val_list, node.Val)
		inOrderTraverse(node.Right)
	}

	inOrderTraverse(node)

	return val_list
}
