package avl

import (
	"testing"
)

func TestInt(t *testing.T) {
	tree := New(intCompare)

	values := []int{5, 3, 2, 4, 7, 6, 8}
	for i := range values {
		tree.Insert(values[i])
	}
	inorder := []int{2, 3, 4, 5, 6, 7, 8}
	postorder := []int{2, 4, 3, 6, 8, 7, 5}
	preorder := []int{5, 3, 2, 4, 7, 6, 8}

	compare := func(list []int, list1 []interface{}) bool {
		for i := range list {
			if list[i] != list1[i].(int) {
				return false
			}
		}
		return true
	}

	if !compare(inorder, tree.Inorder()) || !compare(postorder, tree.Postorder()) || !compare(preorder, tree.Preorder()) {
		t.Error("Tree traversal failed")
	}
	tree.Delete(2)

	if _, err := tree.Search(2); err == nil {
		t.Error("Tree delete failed")
	}

	tree.Insert(12)

	value, err := tree.Search(1)
	if err == nil {
		t.Error("Tree search failed")
	}

	value, err = tree.Search(3)
	if err != nil {
		t.Error("Tree search failed")
	}

	if value.(int) != 3 {
		t.Error("Tree search returned value incorrect")
	}

}

func intCompare(a, b interface{}) int {
	if a.(int) < b.(int) {
		return -1
	} else if a.(int) == b.(int) {
		return 0
	} else {
		return 1
	}
}
