package main

import (
	"fmt"

	"github.com/pseudoelement/go-sandbox/funcs"
)

func main() {
	fmt.Println(funcs.Merge(
		[]int{1, 2, 3, 0, 0, 0},
		3,
		[]int{2, 5, 6},
		3,
	))
}
