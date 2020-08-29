package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkHelper(t, ch)
	close(ch)
}

// Implements recursion logic for in-order tree traversal
func WalkHelper(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	WalkHelper(t.Left, ch)
	ch <- t.Value
	WalkHelper(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
// NOTE: The implementation leaks goroutines when trees are different.
// See exercise-equivalent-binary-trees-quit.go for a better solution.
// The leak happens coz Same() terminates early when trees are different
// so one of the Walk() goroutines is blocked forever waiting for receiver
// of the remaining values of the bigger tree to be passed into its channel
// https://www.ardanlabs.com/blog/2018/11/goroutine-leaks-the-forgotten-sender.html
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

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

func TestWalk() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
}
