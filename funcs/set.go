package funcs

import "log"

type Set struct {
	set []string
}

func NewSet() *Set {
	return &Set{}
}

func (this *Set) Has(val string) bool {
	for _, el := range this.set {
		if string(el) == val {
			return true
		}
	}
	return false
}

func (this *Set) Add(val string) {
	if !this.Has(val) {
		this.set = append(this.set, val)
	}
}

func (this *Set) Remove(val string) {
	idx := -1
	for i, el := range this.set {
		if string(el) == val {
			idx = i
			break
		}
	}

	if idx == -1 {
		return
	}

	this.set = append(this.set[0:idx], this.set[idx+1:]...)
}

func (this *Set) Log() {
	log.Println(this.set)
}

func (this *Set) Size() int {
	size := 0
	for range this.set {
		size++
	}
	return size
}

func (this *Set) ForEach(callback func(idx int, el string)) {
	for idx, el := range this.set {
		callback(idx, el)
	}
}

func (this *Set) Pop() string {
	poppedElement := this.set[len(this.set)-1]
	this.set = this.set[0 : len(this.set)-1]
	return poppedElement
}
