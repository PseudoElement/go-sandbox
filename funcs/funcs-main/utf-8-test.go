package main

import "log"

func Utf8Test() {
	buf := []byte("👻")
	buf2 := []byte("01234")
	buf3 := []byte("🏴󠁧󠁢󠁷󠁬󠁳󠁿")

	// buf3[i] is a CodeUnit
	for i := 0; i < len(buf3); i++ {
		log.Println("idx_", i, ": ", buf3[i])
	}
	// it's a CodePoint
	for i, utf8Rune := range "🏴󠁧󠁢󠁷󠁬󠁳󠁿" {
		log.Println("FlAG_idx_", i, ": ", utf8Rune)
	}

	log.Println("BUF ==>", buf)
	log.Println("BUF2 ==>", buf2)
	log.Println("BUF3 ==>", buf3)
}
