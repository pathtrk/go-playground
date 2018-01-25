package main

import "fmt"

// fibonacci returns a function that calcurates the series
func fibonacci() func() int {
	n0, n1 := 0, 1
	return func() int {
		r := n0
		n2 := n0 + n1
		n0, n1 = n1, n2
		return r
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
