package main

import (
	"log"
)

func main() {
	// tokens, err := tokenizeLayout("2006-01-02 15:04:05 MST")
	// tokens, err := tokenizeLayout("02-01-06 15:04")
	t, err := parseTimeString("2006-01-02 15:04:05 MST", "2008-02-03 17:18:19 MST")
	log.Println("err: ", err)
	log.Println("year: ", t.Year())
	log.Println("month: ", t.Month())
	log.Println("day: ", t.Day())
	log.Println("hour: ", t.Hour())
	log.Println("min: ", t.Minute())
	log.Println("sec: ", t.Second())
	log.Println("location: ", t.Location().String())
	// for idx, char := range " .-:,;" {
	// 	log.Println(idx, " utf8 code - ", char)
	// 	log.Println(idx, " utf8 string - ", string(char))
	// 	log.Println("========")
	// }
}
