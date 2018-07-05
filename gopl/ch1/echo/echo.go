package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	// echo 2
	// s, sep := "", ""

	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }

	// fmt.Println(s)

	// echo 3
	// fmt.Println(strings.Join(os.Args[1:], " "))

	// 1.1
	// fmt.Println(strings.Join(os.Args[0:], " "))

	// 1.2
	// for i, arg := range os.Args[1:] {
	// 	fmt.Printf("%d %s\n", i+1, arg)
	// }

	// 1.3
	// start := time.Now()

	// s, sep := "", ""
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)

	// fmt.Println(time.Since(start))

	start := time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println(time.Since(start))
}
