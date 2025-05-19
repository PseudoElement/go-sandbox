package utils

import "testing"

// 100k - 64	        18105091 ns/op	       3590028 B/op	    247022 allocs/op
// 10k  - 514           2157853 ns/op          382260 B/op      27166 allocs/op
func BenchmarkUseAtomicPointerForStruct(b *testing.B) {
	for b.Loop() {
		UseAtomicPointerForStruct(100_000)
	}
}

// 100k -  58	         20035294 ns/op	        3223628 B/op	 100129 allocs/op
// 10k  -  565           2127428 ns/op          325108 B/op      10028 allocs/op
func BenchmarkUseMutexForStruct(b *testing.B) {
	for b.Loop() {
		UseMutexForStruct(100_000)
	}
}
