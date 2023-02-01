package tree

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type TreeNode struct {
	key       int
	Value     string
	leftNode  *TreeNode
	rightNode *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
	lock sync.RWMutex
}

func newTreeNode(key int, value string) *TreeNode {
	return &TreeNode{key: key, Value: value}
}

func (bst *BinarySearchTree) InsertNode(key int, value string) *BinarySearchTree {
	bst.lock.Lock()
	defer bst.lock.Unlock()

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
	bst.lock.RLock()
	defer bst.lock.RUnlock()
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
	bst.lock.RLock()
	defer bst.lock.RUnlock()
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
	bst.lock.RLock()
	defer bst.lock.RUnlock()
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
	bst.lock.RLock()
	defer bst.lock.RUnlock()
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
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	var val_list []string
	for _, key := range bst.InOrderTraverseTree() {
		val_list = append(val_list, strconv.Itoa(key))
	}
	result := strings.Join(val_list, "->")

	fmt.Println(result)

	return result
}

func (bst *BinarySearchTree) DeleteNode(key int) *TreeNode {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	bst.root = deleteNode(bst.root, key)
	return bst.root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.key == key {
		if root.leftNode == nil {
			return root.rightNode
		}

		if root.rightNode == nil {
			return root.leftNode
		}

		cur := root.rightNode
		for cur.leftNode != nil {
			cur = cur.rightNode
		}

		cur.leftNode = root.leftNode

		root = root.rightNode
		return root
	}

	if root.key < key {
		root.rightNode = deleteNode(root.rightNode, key)
	}

	if root.key > key {
		root.leftNode = deleteNode(root.leftNode, key)
	}

	return root
}
