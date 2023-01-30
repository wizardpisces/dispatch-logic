package tree

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	key       int
	value     string
	leftNode  *TreeNode
	rightNode *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
	// lock     sync.RWMutex
}

func newTreeNode(key int, value string) *TreeNode {
	return &TreeNode{key: key, value: value}
}

func (bst *BinarySearchTree) InsertNode(key int, value string) *BinarySearchTree {
	node := newTreeNode(key, value)
	if bst.root == nil {
		bst.root = node
	} else {
		insert(bst.root, node)
	}
	return bst
}

func insert(parentNode *TreeNode, newNode *TreeNode) {
	if newNode.key < parentNode.key {
		if parentNode.leftNode == nil {
			parentNode.leftNode = newNode
		} else {
			insert(parentNode.leftNode, newNode)
		}
	} else {
		if parentNode.rightNode == nil {
			parentNode.rightNode = newNode
		} else {
			insert(parentNode.rightNode, newNode)
		}
	}
}

func (bst *BinarySearchTree) InOrderTraverseTree() []int {
	var val_list []int
	var inOrderTraverse func(node *TreeNode)

	inOrderTraverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrderTraverse(node.leftNode)
		val_list = append(val_list, node.key)
		inOrderTraverse(node.rightNode)
	}

	inOrderTraverse(bst.root)

	return val_list
}

func (bst *BinarySearchTree) MinNode() *TreeNode {
	if bst.root != nil {
		return minNode(bst.root)
	}
	return nil
}

func minNode(node *TreeNode) *TreeNode {
	if node.leftNode != nil {
		return minNode(node.leftNode)
	}
	return node
}

func (bst *BinarySearchTree) MaxNode() *TreeNode {
	if bst.root != nil {
		return maxNode(bst.root)
	}
	return nil
}

func maxNode(node *TreeNode) *TreeNode {
	if node.rightNode != nil {
		return maxNode(node.rightNode)
	}
	return node
}

func (bst *BinarySearchTree) SearchNode(key int) bool {
	return searchNode(bst.root, key)
}

func searchNode(node *TreeNode, key int) bool {
	if node == nil {
		return false
	}
	if key == node.key {
		return true
	}
	if key < node.key {
		return searchNode(node.leftNode, key)
	}
	if key > node.key {

		return searchNode(node.rightNode, key)
	}
	return false
}

func (bst *BinarySearchTree) Display() string {
	var val_list []string
	for _, key := range bst.InOrderTraverseTree() {
		val_list = append(val_list, strconv.Itoa(key))
	}
	result := strings.Join(val_list, "->")

	fmt.Println(result)

	return result
}
