package main

import (
	"fmt"

	"github.com/pseudoelement/go-sandbox/funcs"
)

func main() {
	fmt.Println("PascalsTriangle(1) ==>", funcs.PascalsTriangle(1))
	fmt.Println("PascalsTriangle(5) ==>", funcs.PascalsTriangle(5))

	// t, err := parseTimeString("02.01.06 15:04:05 MST", "10.12.98 22:11:59 MST")
	// t, err := parseTimeString("2006-01-02 15:04:05 MST", "2008-02-03 17:18:19 MST")
	// t, err := parseTimeString("Jun 2, 2006 15::04::05 MST", "Dec 23, 2004 21::15::15 MST")
	//
	// log.Println("err: ", err)
	// log.Println("year: ", t.Year())
	// log.Println("month: ", t.Month())
	// log.Println("day: ", t.Day())
	// log.Println("hour: ", t.Hour())
	// log.Println("min: ", t.Minute())
	// log.Println("sec: ", t.Second())
	// log.Println("location: ", t.Location().String())
}
