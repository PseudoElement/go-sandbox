package main

import (
	"log"

	"github.com/pseudoelement/go-sandbox/funcs"
)

func main() {
	res := funcs.GetLastDecisions(funcs.Decisions_1, funcs.Decisions_2, 7)
	log.Printf("res ==> %+v\n", res)
}
