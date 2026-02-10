package dir1

import (
	"unsafe"

	"github.com/pseudoelement/go-sandbox/unsafe/store"
)

func Init() {
	fn := dir1Call
	store.GLOBAL["dir1Call"] = unsafe.Pointer(&fn)
}

func CrossCall() {
	ptr := store.GLOBAL["dir2Call"]
	dir2Call := (*func() string)(ptr)
	value := (*dir2Call)()
	println("dir1_CrossCall_value", value)
}

func dir1Call() string {
	println("dir1Call func called!")
	return "dir1_Call"
}
