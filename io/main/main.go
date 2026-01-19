package main

import "log"

func main() {
	// start address 0x001
	dst := make([]int, 0, 6)
	dst = append(dst, 0, 1, 2, 3)

	// 0x001, 0x002, 0x003
	src := dst[0:2]
	src = append(src, 4)
	src[0] = 15

	// copy(dst, src)

	log.Println("src", src)
	log.Println("dst", dst)
}
