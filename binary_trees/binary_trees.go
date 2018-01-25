package main

import (
	"fmt"
	"sort"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	value, right, left := t.Value, t.Right, t.Left

	if right.String() != "()" {
		Walk(right, ch)
	}
	if left.String() != "()" {
		Walk(left, ch)
	}

	ch <- value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int, 10), make(chan int, 10)
	ar1, ar2 := make([]int, 10), make([]int, 10)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		ar1 = append(ar1, <-ch1)
		ar2 = append(ar2, <-ch2)
	}

	sort.Ints(ar1)
	sort.Ints(ar2)

	for i, v := range ar1 {
		if v != ar2[i] {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < cap(ch); i++ {
			fmt.Println(<-ch)
		}
	}()
	Walk(tree.New(1), ch)

	fmt.Printf("tree.New(1) == tree.New(1) : %v\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("tree.New(1) == tree.New(2) : %v\n", Same(tree.New(1), tree.New(2)))
}
