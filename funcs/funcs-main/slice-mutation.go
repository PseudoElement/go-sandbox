package main

import "log"

func SliceMutation() {
	data1 := make([]int, 3, 6)
	data2 := data1[1:2] // 1:6
	data3 := data1[1:3] // 2:6

	// 1st solution to not overwrite data3
	// var data3 []int = make([]int, 2)
	// copy(data3, data1[1:3])

	// 2nd solution to not overwrite data3
	// data3 := data1[1:3:3]

	println("d1 -", len(data1), cap(data1))
	println("d2 -", len(data2), cap(data2))
	println("d3 -", len(data3), cap(data3))
	// data2[0] = 12 // 3:4
	// data1[1] = 5  // 5:7

	data3 = append(data3, 333)      // -> doesn't mutate data1, data2 (cause len(data1) = 3 and it appends in next memory slot)
	data2 = append(data2, 228, 225) // -> mutates data1, data3 (cause appends in data1[2] element)

	log.Printf("data1 ==> %+v\n", data1) // [0 0 228]
	log.Printf("data2 ==> %+v\n", data2) // [0 228 225]
	log.Printf("data3 ==> %+v\n", data3) // [0 228 225]
}
