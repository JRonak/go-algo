package bst

import (
	"errors"
	"github.com/JRonak/go-algo/stack"
)

var (
	BSTNOTFOUND = errors.New("BST: Element not found")
)

type bstNode struct {
	data        interface{}
	left, right *bstNode
}

type BinarySeachTree struct {
	root        *bstNode                   // counts the nodes
	compareFunc func(a, b interface{}) int // returns true if the value of a is greater than or equal to b
}

func NewBinarySeachTree(compare func(a, b interface{}) int) BinarySeachTree {
	tree := BinarySeachTree{}
	tree.compareFunc = compare
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
	if this.root == nil {
		this.root = newBstNode(data)
	} else {
		root := this.root
		var previous *bstNode
		previous = nil
		for root != nil {
			previous = root
			if this.compareFunc(root.data, data) == 1 {
				root = root.left
			} else {
				root = root.right
			}
		}
		if this.compareFunc(previous.data, data) == 1 {
			previous.left = newBstNode(data)
		} else {
			previous.right = newBstNode(data)
		}
	}
}

func (this *BinarySeachTree) Delete(data interface{}) bool {
	parentleft := false //if the node to be deleted is the left child of the parent
	node := this.root
	var parent *bstNode
	for node != nil {
		if this.compareFunc(data, node.data) == 0 {
			break
		}
		parent = node
		if this.compareFunc(node.data, data) == 1 {
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
	if node == this.root {
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
		pred, parent := this.predecessor(node)
		if pred != node.left {
			parent.right = pred.left
			pred.left = node.left
			pred.right = node.right
		} else {
			pred.right = node.right
		}
		finalNode = pred
	}
	if isRoot {
		this.root = finalNode
	} else if parentleft {
		parent.left = finalNode
	} else {
		parent.right = finalNode
	}
	return true
}

//Finds the right most node and updates the parent node of the right most node(predecessor)
func (this *BinarySeachTree) predecessor(root *bstNode) (node, parent *bstNode) {
	if root.left == nil {
		node = nil
		return
	}
	parent = root
	node = root.left
	if node.right == nil {
		return
	}
	for node.right != nil && node != root {
		parent = node
		node = node.right
	}
	/*
		parent.right = node.left
		node.left = nil*/
	return
}

//Morris inorder traversal
//non-recursice and non-stack based inorder traversal of the tree
func (this *BinarySeachTree) Inorder() []interface{} {
	data := stack.NewStack()
	node := this.root
	for node != nil {
		if node.left == nil {
			data.Push(node.data)
			node = node.right
		} else {
			pred, parent := this.predecessor(node)
			if pred != node {
				pred.right = node
				node = node.left
			} else {
				parent.right = nil
				data.Push(node.data)
				node = node.right
			}
		}
	}
	return data.Data()
}

//non-recursive pre-order traversal
func (this *BinarySeachTree) Preorder() []interface{} {
	if this.root == nil {
		return nil
	}
	data := stack.NewStack()
	stack := stack.NewStack()
	stack.Push(this.root)
	for !stack.Empty() {
		node := stack.Pop().(*bstNode)
		data.Push(node.data)
		if node.right != nil {
			stack.Push(node.right)
		}
		if node.left != nil {
			stack.Push(node.left)
		}
	}
	return data.Data()
}

//non-recursive post-order traversal using stack
func (this *BinarySeachTree) Postorder() []interface{} {
	if this.root == nil {
		return nil
	}
	data := stack.NewStack()
	stack := stack.NewStack()
	node := this.root
	for node != nil || !stack.Empty() {
		if node != nil {
			if node.right != nil {
				stack.Push(node.right)
			}
			stack.Push(node)
			node = node.left
		} else {
			temp := stack.Pop().(*bstNode)
			if temp.right == nil {
				data.Push(temp.data)
			} else {
				if !stack.Empty() && stack.Peek() == temp.right {
					node = stack.Pop().(*bstNode)
					stack.Push(temp)
				} else {
					data.Push(temp.data)
				}
			}
		}
	}
	return data.Data()
}

func (this *BinarySeachTree) Search(data interface{}) (interface{}, error) {
	root := this.root
	for root != nil {
		if this.compareFunc(data, root.data) == 0 {
			return root.data, nil
		} else if this.compareFunc(root.data, data) == 1 {
			root = root.left
		} else {
			root = root.right
		}
	}
	return nil, BSTNOTFOUND
}
