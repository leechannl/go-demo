package main

import (
	"golang.org/x/tour/tree"
)
import "fmt"

// _walk walks the tree t sending all values from the tree to the channel ch
func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

func _walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		_walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		_walk(t.Right, ch)
	}
}

// Same determine whether the trees t1 and t2 contain the same values
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v1 := range ch1 {
		if v1 != <-ch2 {
			return false
		}
	}
	if _, ok := <-ch2; ok {
		return false
	}
	return true
}

func insert(t *tree.Tree, v int) *tree.Tree {
	if t == nil {
		return &tree.Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for elem := range ch {
		fmt.Println(elem)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	tree2 := tree.New(1)
	insert(tree2, 11)
	fmt.Println(Same(tree.New(1), tree2))
	fmt.Println(tree.New(1))
	fmt.Println(tree2)
}
