package avl

import (
	"errors"
	"github.com/JRonak/go-algo/stack"
	"log"
)

var (
	AERR = errors.New("AVL tree operation failed")
	ADUP = errors.New("Avl tree already has a element")
)

type avlNode struct {
	data        interface{}
	height      int
	left, right *avlNode
}

type AvlTree struct {
	root    *avlNode
	compare func(a, b interface{}) int
}

func New(compare func(a, b interface{}) int) *AvlTree {
	tree := new(AvlTree)
	tree.compare = compare
	return tree
}

func newAvlNode(data interface{}) *avlNode {
	n := new(avlNode)
	n.height = 1
	n.data = data
	return n
}

func (this *AvlTree) Insert(data interface{}) {
	this.root = this.recursiveInsert(this.root, &data)
}

func (this *AvlTree) Delete(data interface{}) {
	this.root = this.recursiveDelete(this.root, &data)
}

func (this *AvlTree) Search(data interface{}) (interface{}, error) {
	root := this.searchNode(data)
	if root != nil {
		return root.data, nil
	} else {
		return nil, AERR
	}
}

func (this *AvlTree) Update(data interface{}) error {
	node := this.searchNode(data)
	if node != nil {
		node.data = data
		return nil
	} else {
		return AERR
	}
}

//recursive insert
func (this *AvlTree) recursiveInsert(root *avlNode, data *interface{}) *avlNode {
	if root != nil {
		if this.compare(root.data, *data) == 1 {
			root.left = this.recursiveInsert(root.left, data)
		} else {
			root.right = this.recursiveInsert(root.right, data)
		}
		return balance(root)
	} else {
		return newAvlNode(*data)
	}
}

func (this *AvlTree) recursiveDelete(root *avlNode, data *interface{}) *avlNode {
	if root == nil {
		return nil
	} else {
		status := this.compare(root.data, *data)
		if status == 1 {
			root.left = this.recursiveDelete(root.left, data)
		} else if status == 0 {
			root = this.deleteNode(root)
		} else {
			root.right = this.recursiveDelete(root.right, data)
		}
		root = balance(root)
	}
	return root
}

func (this *AvlTree) deleteNode(root *avlNode) *avlNode {
	if root.left == nil && root.right == nil {
		return nil
	} else if root.left != nil && root.right == nil {
		return root.left
	} else if root.right != nil && root.left == nil {
		return root.right
	} else {
		newRoot, _ := predecessor(root)
		root.data = newRoot.data
		root.left = this.recursiveDelete(root.left, &newRoot.data)
		return root
	}
}

func (this *AvlTree) searchNode(data interface{}) *avlNode {
	root := this.root
	if root == nil {
		return nil
	} else {
		for root != nil {
			status := this.compare(root.data, data)
			if status == 1 {
				root = root.left
			} else if status == 0 {
				return root
			} else {
				root = root.right
			}
		}
	}
	return nil
}

func predecessor(root *avlNode) (child *avlNode, parent *avlNode) {
	if root.left == nil {
		log.Println(AERR)
		return
	}
	parent = nil
	child = root.left
	for child.right != nil && child != root {
		parent = child
		child = child.right
	}
	return
}

func balance(root *avlNode) *avlNode {
	if root == nil {
		return nil
	}
	diff := getHeight(root.left) - getHeight(root.right)
	if diff > 1 && getHeight(root.left.left) > getHeight(root.left.right) {
		root = rightRotate(root)
	} else if diff > 1 {
		root.left = leftRotate(root.left)
		root = rightRotate(root)
	} else if diff < -1 && getHeight(root.right.left) > getHeight(root.right.right) {
		root.right = rightRotate(root.right)
		root = leftRotate(root)
	} else if diff < -1 {
		root = leftRotate(root)
	}
	root.height = 1 + maxHeight(root.left, root.right)
	return root
}

//Morris inorder traversal
//non-recursice and non-stack based inorder traversal of the tree
func (this *AvlTree) Inorder() []interface{} {
	data := stack.NewStack()
	node := this.root
	for node != nil {
		if node.left == nil {
			data.Push(node.data)
			node = node.right
		} else {
			pred, parent := predecessor(node)
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
func (this *AvlTree) Preorder() []interface{} {
	if this.root == nil {
		return nil
	}
	data := stack.NewStack()
	stack := stack.NewStack()
	stack.Push(this.root)
	for !stack.Empty() {
		node := stack.Pop().(*avlNode)
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
func (this *AvlTree) Postorder() []interface{} {
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
			temp := stack.Pop().(*avlNode)
			if temp.right == nil {
				data.Push(temp.data)
			} else {
				if !stack.Empty() && stack.Peek() == temp.right {
					node = stack.Pop().(*avlNode)
					stack.Push(temp)
				} else {
					data.Push(temp.data)
				}
			}
		}
	}
	return data.Data()
}

//TO-do rotate error
func leftRotate(root *avlNode) *avlNode {
	if root.right == nil {
		panic(AERR)
	}
	newRoot := root.right
	root.right = newRoot.left
	newRoot.left = root
	root.height = 1 + maxHeight(root.left, root.right)
	newRoot.height = 1 + maxHeight(newRoot.left, newRoot.right)
	return newRoot
}

func rightRotate(root *avlNode) *avlNode {
	if root.left == nil {
		panic(AERR)
	}
	newRoot := root.left
	root.left = newRoot.right
	newRoot.right = root
	root.height = 1 + maxHeight(root.left, root.right)
	newRoot.height = 1 + maxHeight(newRoot.left, newRoot.right)
	return newRoot
}

func getHeight(a *avlNode) int {
	if a != nil {
		return a.height
	} else {
		return 0
	}
}

func maxHeight(a *avlNode, b *avlNode) int {
	asize := getHeight(a)
	bsize := getHeight(b)
	if asize > bsize {
		return asize
	}
	return bsize
}
