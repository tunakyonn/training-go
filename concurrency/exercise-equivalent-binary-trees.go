package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t.Left != nil {
			walker(t.Left)
		}
		ch <- t.Value
		if t.Right != nil {
			walker(t.Right)
		}
	}
	walker(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
//func Same(t1, t2 *tree.Tree) bool

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
}
