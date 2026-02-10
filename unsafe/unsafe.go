package main

import (
	"unsafe"

	"github.com/pseudoelement/go-sandbox/unsafe/dir1"
	"github.com/pseudoelement/go-sandbox/unsafe/dir2"
)

// var

func main() {
	dir1.Init()
	dir2.Init()

	dir1.CrossCall()
	dir2.CrossCall()

	fn := makeCall
	fnPtr := unsafe.Pointer(&fn)

	println("Ptr ==>", fnPtr)

	makeCallFn := (*func(value string) *HexString)(fnPtr)

	hexStr1 := (*makeCallFn)("0x2210F8De2F3406Aa38e7388C176B5e5C9b8352E8")
	hexStr2 := (*makeCallFn)("2210F8De2F3406Aa38e7388C176B5e5C9b8352E8")

	println("hexStr1", hexStr1.Value(), hexStr1.Valid())
	println("hexStr2", hexStr2.Value(), hexStr2.Valid())

	if hexStr1.Valid() {
		decimal, err := hexStr1.ToDecimal()
		println("hexStr1_decimal", decimal.String(), decimal.Int64(), decimal.IsInt64(), err)
	}
	if hexStr2.Valid() {
		decimal, err := hexStr2.ToDecimal()
		println("hexStr2_decimal", decimal, err)
	}
}
