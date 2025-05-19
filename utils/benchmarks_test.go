package utils

import "testing"

// STRUCTS

// 100k - 18105091 ns/op	     3590028 B/op	  247022 allocs/op
// 10k  - 2157853 ns/op          382260 B/op      27166 allocs/op
func BenchmarkUseAtomicForStruct(b *testing.B) {
	for b.Loop() {
		UseAtomicPointerForStruct(100_000)
	}
}

// 100k -  20035294 ns/op	      3223628 B/op	   100129 allocs/op
// 10k  -  2127428 ns/op          325108 B/op      10028 allocs/op
func BenchmarkUseMutexForStruct(b *testing.B) {
	for b.Loop() {
		UseMutexForStruct(100_000)
	}
}

// INTEGERS

// 100k - 17927631 ns/op         2408764 B/op     100022 allocs/op
// 10k  - 1835937 ns/op	  		 242631 B/op	  10008 allocs/op
func BenchmarkUseAtomicForInt(b *testing.B) {
	for b.Loop() {
		UseAtomicPointerForInt(10_000)
	}
}

// 100k - 19957969 ns/op         3226063 B/op     100146 allocs/op
// 10k  - 1946682 ns/op	  		 322231 B/op	  10015 allocs/op
func BenchmarkUseMutexForInt(b *testing.B) {
	for b.Loop() {
		UseMutexForInt(10_000)
	}
}
