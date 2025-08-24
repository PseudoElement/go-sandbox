package main

import (
	"log"
	"reflect"
)

type TestArr struct {
	value *Obj
}

func (this *TestArr) SetName() {
	this.value.Name = "Peter"
}

type Obj struct {
	Name string
	Age  int
}

func change(arr *TestArr) {
	arr.value.Name = "Tomas"
}

func ChangeSlices() {
	arr := &TestArr{value: &Obj{}}
	arr.value = &Obj{Name: "Bill", Age: 11}
	arr2 := arr

	m := reflect.ValueOf(arr).MethodByName("SetName")
	if m.IsValid() {
		println("valid!")
		m.Call([]reflect.Value{})
	}

	change(arr)

	log.Printf("arr %+v", arr.value)
	log.Printf("arr2 %+v", arr2.value)

	a := []int{1, 2}
	a2 := a
	a3 := make([]int, 2)
	copy(a3, a)

	a2[0] = 99

	log.Println(a)
	log.Println(a2)
	log.Println(a3)
}
