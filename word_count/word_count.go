package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

/*
WordCount :
  The function counts words in slice "s",
  then returns the result in a map
*/
func WordCount(s string) map[string]int {
	arr := strings.Split(s, " ")
	fmt.Printf("%q\n", arr)

	var dict = make(map[string]int)

	for _, word := range arr {
		dict[word]++
	}
	fmt.Printf("%v\n", dict)

	return dict
}

func main() {
	wc.Test(WordCount)
}
