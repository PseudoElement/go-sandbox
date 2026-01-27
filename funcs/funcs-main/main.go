package main

import (
	"log"

	"github.com/pseudoelement/go-sandbox/funcs"
)

func main() {
	res2 := funcs.Exist(
		[][]byte{
			{'A', 'A', 'A', 'A'},
			{'A', 'A', 'A', 'A'},
			{'A', 'A', 'A', 'A'},
		},
		"AAAAAAAAAAAAA",
	)
	log.Println("RES =>>", res2)
}
