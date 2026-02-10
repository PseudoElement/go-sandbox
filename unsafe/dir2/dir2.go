package dir2

import (
	"unsafe"

	"github.com/pseudoelement/go-sandbox/unsafe/store"
)

func Init() {
	fn := dir2Call
	store.GLOBAL["dir2Call"] = unsafe.Pointer(&fn)
}

func CrossCall() {
	ptr := store.GLOBAL["dir1Call"]
	dir1Call := (*func() string)(ptr)
	value := (*dir1Call)()
	println("dir2_CrossCall_value", value)
}

func dir2Call() string {
	println("dir2Call func called!")
	return "dir2_Call"
}
