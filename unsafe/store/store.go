package store

import "unsafe"

type MethodName string

var GLOBAL = map[MethodName]unsafe.Pointer{}
