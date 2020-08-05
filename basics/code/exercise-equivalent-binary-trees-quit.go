package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, quit chan int) {
	WalkHelper(t, ch, quit)
	close(ch)
}

// Implements recursion logic for in-order tree traversal
func WalkHelper(t *tree.Tree, ch chan int, quit chan int) {
	if t == nil {
		return
	}
	WalkHelper(t.Left, ch, quit)
	select {
	case ch <- t.Value:
		// value sent to channel successfully
	case <-quit:
		return
	}
	WalkHelper(t.Right, ch, quit)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2, quit := make(chan int), make(chan int), make(chan int)
	defer close(quit)

	go Walk(t1, c1, quit)
	go Walk(t2, c2, quit)

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2

		if !ok1 || !ok2 {
			return ok1 == ok2 // true if both end together
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	// TestWalk()
	fmt.Print("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

	fmt.Print("tree.New(1) != tree.New(2): ")
	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}
}
