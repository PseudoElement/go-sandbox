package main

import "github.com/pseudoelement/go-sandbox/funcs"

func main() {
	println("Start zed app...")

	// funcs.RemoveElement([]int{1, 2, 3, 4, 5, 1, 1}, 1)
	// funcs.RemoveElement([]int{3, 2, 2, 3}, 3)
	// funcs.RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	funcs.RemoveElement([]int{3, 3}, 3)

	println("End zed app...")
}
