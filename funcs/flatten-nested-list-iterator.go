package funcs

type NestedIterator struct {
	flattened []int
	index     int
}

func NewNestedIterator(nestedList []*NestedInteger) *NestedIterator {
	iterator := &NestedIterator{}
	iterator.flatten(nestedList)
	return iterator
}

func (ni *NestedIterator) flatten(list []*NestedInteger) {
	for _, item := range list {
		if item.IsInteger() {
			ni.flattened = append(ni.flattened, item.GetInteger())
		} else {
			ni.flatten(item.GetList())
		}
	}
}

func (ni *NestedIterator) Next() int {
	if ni.HasNext() {
		val := ni.flattened[ni.index]
		ni.index++
		return val
	}
	return 0
}

func (ni *NestedIterator) HasNext() bool {
	return ni.index < len(ni.flattened)
}

type NestedInteger struct {
	Integer int
	Nested  []*NestedInteger
}

func NewNestedInteger(val int) *NestedInteger {
	return &NestedInteger{Integer: val, Nested: make([]*NestedInteger, 0)}
}

func (ni *NestedInteger) IsInteger() bool { return len(ni.Nested) == 1 }

func (ni *NestedInteger) GetInteger() int { return ni.Integer }

func (ni *NestedInteger) SetInteger(value int) { ni.Integer = value }

func (ni *NestedInteger) Add(elem *NestedInteger) {
	ni.Nested = append(ni.Nested, elem)
}

func (ni NestedInteger) GetList() []*NestedInteger { return ni.Nested }
