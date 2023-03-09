package prune_tree

// type Digit *int
type Digit any

// const (
// 	Zero Digit = iota
// 	One
// )

// type Digit interface {
// 	Digit | any
// }

type Node struct {
	Val   Digit
	left  *Node
	right *Node
}

// 从广度优先遍历序列化的数组还原树  construct tree from a tree's bfs array data
func ConstructTree(arr []Digit) (root *Node) {
	if len(arr) == 0 {
		return &Node{}
	}
	i := 0
	stack := []*Node{{Val: arr[i]}}
	root = stack[0]

	for i+2 < len(arr) && len(stack) > 0 {
		curNode := stack[0]
		stack = stack[1:]
		if curNode.Val != nil {
			curNode.left = &Node{Val: arr[i+1]}
			curNode.right = &Node{Val: arr[i+2]}
			i = i + 2
			stack = append(stack, curNode.left, curNode.right)
		}
	}

	return root
}

// func ConstructTree(arr []Digit) *Node {
// 	if len(arr) == 0 {
// 		return &Node{}
// 	}

// 	arr = append([]Digit{nil}, arr...)
// 	i := 1
// 	root := &Node{Val: arr[i]}
// 	i = 2 * i
// 	for i+1 <= len(arr) {
// 		root.left = &Node{Val: arr[i]}
// 		root.right = &Node{Val: arr[i+1]}
// 		i = 2 * i
// 	}
// 	return root
// }

func PruneTree(root *Node) *Node {
	if root == nil {
		return nil
	}
	root.left = PruneTree(root.left)
	root.right = PruneTree(root.right)
	if root.left == nil && root.right == nil && root.Val == 0 {
		return nil
	}
	return root
}

func SerializeTree(root *Node) (arr []Digit) { // bfs serialize tree
	stack := []*Node{root}

	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]

		if node == nil {
			arr = append(arr, nil)
			continue
		}
		arr = append(arr, node.Val)

		if node.left == nil && node.right == nil {
			continue
		}

		stack = append(stack, node.left)
		stack = append(stack, node.right)
	}

	return arr
}

// todo 抽象成遍历树的模板公共方法
func InOrderTraverseTree(node *Node) []Digit {
	var val_list []Digit
	var inOrderTraverse func(node *Node)

	inOrderTraverse = func(node *Node) {
		if node == nil {
			return
		}
		inOrderTraverse(node.left)
		val_list = append(val_list, node.Val)
		inOrderTraverse(node.right)
	}

	inOrderTraverse(node)

	return val_list
}
