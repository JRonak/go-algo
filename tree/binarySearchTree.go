package tree

import (
	"errors"
	"github.com/JRonak/containers"
)

var (
	BSTFATAL    = errors.New("Binary search tree operation failed")
	BSTNOTFOUND = errors.New("Element not found in the tree")
)

type bstNode struct {
	data        interface{}
	left, right *bstNode
}

type BinarySeachTree struct {
	Root        *bstNode                    // counts the nodes
	CompareFunc func(a, b interface{}) bool // returns true if the value of a is greater than or equal to b
	EqualFunc   func(a, b interface{}) bool //true if value matches
}

func GetBinarySeachTree(compare, equal func(a, b interface{}) bool) BinarySeachTree {
	tree := BinarySeachTree{}
	tree.CompareFunc = compare
	tree.EqualFunc = equal
	return tree
}

func newBstNode(data interface{}) *bstNode {
	node := new(bstNode)
	node.left = nil
	node.right = nil
	node.data = data
	return node
}

func (this *BinarySeachTree) Insert(data interface{}) {
	if this.Root == nil {
		this.Root = newBstNode(data)
	} else {
		root := this.Root
		var previous *bstNode
		previous = nil
		for root != nil {
			previous = root
			if this.CompareFunc(root.data, data) {
				root = root.left
			} else {
				root = root.right
			}
		}
		if this.CompareFunc(previous.data, data) {
			previous.left = newBstNode(data)
		} else {
			previous.right = newBstNode(data)
		}
	}
}

func (this *BinarySeachTree) Delete(data interface{}) bool {
	parentleft := false //if the node to be deleted is the left child of the parent
	node := this.Root
	var parent *bstNode
	for node != nil {
		if this.EqualFunc(data, node.data) {
			break
		}
		parent = node
		if this.CompareFunc(node.data, data) {
			node = node.left
			parentleft = true
		} else {
			node = node.right
			parentleft = false
		}
	}
	// node not found in the tree
	if node == nil {
		return false
	}
	isRoot := false
	if node == this.Root {
		isRoot = true
	}
	var finalNode *bstNode //final node to be updated in the parent
	if node.left == nil && node.right == nil {
		finalNode = nil
	} else if node.left == nil {
		finalNode = node.right
	} else if node.right == nil {
		finalNode = node.left
	} else {
		pred := this.modifyPredecessor(node)
		if pred != node.left {
			pred.left = node.left
			pred.right = node.right
		} else {
			pred.right = node.right
		}
		finalNode = pred
	}
	if isRoot {
		this.Root = finalNode
	} else if parentleft {
		parent.left = finalNode
	} else {
		parent.right = finalNode
	}
	return true
}

//Finds the right most node and updates the parent node of the right most node(predecessor)
func (this *BinarySeachTree) modifyPredecessor(root *bstNode) *bstNode {
	if root.left == nil {
		return nil
	}
	node := root.left
	if node.right == nil {
		return node
	}
	var parent *bstNode
	for node.right != nil {
		parent = node
		node = node.right
	}
	parent.right = node.left
	node.left = nil
	return node
}

//returns predecessor where the root is not equal to the  root
func (this *BinarySeachTree) predecessor(root *bstNode) *bstNode {
	if root.left == nil {
		return nil
	}
	node := root.left
	for node != root && node.right != nil {
		node = node.right
	}
	return node
}

//returns predecessor where the node right is not equal to the root
func (this *BinarySeachTree) predecessorNonRight(root *bstNode) *bstNode {
	if root.left == nil {
		return nil
	}
	node := root.left
	for node.right != root && node.right != nil {
		node = node.right
	}
	return node
}

//Morris inorder traversal
//non-recursice and non-stack based inorder traversal of the tree
func (this *BinarySeachTree) Inorder() []interface{} {
	data := containers.NewStack()
	node := this.Root
	for node != nil {
		if node.left == nil {
			data.Push(node.data)
			node = node.right
		} else {
			pred := this.predecessor(node)
			if pred != node {
				pred.right = node
				node = node.left
			} else {
				this.predecessorNonRight(node).right = nil
				data.Push(node.data)
				node = node.right
			}
		}
	}
	return data.Data()
}

//non-recursive pre-order traversal
func (this *BinarySeachTree) Preorder() []interface{} {
	if this.Root == nil {
		return nil
	}
	var (
		list  []interface{}
		stack []*bstNode
		node  *bstNode
	)
	index := 0
	stack = append(stack, this.Root)
	index++
	for index > 0 {
		node = stack[index-1]
		stack = stack[:index-1]
		index--
		list = append(list, node.data)
		if node.right != nil {
			stack = append(stack, node.right)
			index++
		}
		if node.left != nil {
			stack = append(stack, node.left)
			index++
		}
	}
	return list
}

//non-recursive post-order traversal using stack
func (this *BinarySeachTree) Postorder() []interface{} {
	var (
		list  []interface{}
		stack []*bstNode
		node  *bstNode
	)
	if this.Root == nil {
		return nil
	}
	index := 0
	node = this.Root
	for node != nil || index > 0 {
		if node != nil {
			if node.right != nil {
				stack = append(stack, node.right)
				index++
			}
			stack = append(stack, node)
			index++
			node = node.left
		} else {
			temp := stack[index-1]
			stack = stack[:index-1]
			index--
			if temp.right == nil {
				list = append(list, temp.data)
			} else {
				if index > 0 && stack[index-1] == temp.right {
					node = stack[index-1]
					stack[index-1] = temp
				} else {
					list = append(list, temp.data)
				}
			}
		}
	}
	return list
}

func (this *BinarySeachTree) Search(data interface{}) (interface{}, error) {
	root := this.Root
	for root != nil {
		if this.EqualFunc(data, root.data) {
			return root.data, nil
		} else if this.CompareFunc(root.data, data) {
			root = root.left
		} else {
			root = root.right
		}
	}
	return nil, BSTNOTFOUND
}
