package main

import "log"

type Stack[T any] struct {
	store    []T
	stackPtr int
}

func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		store:    make([]T, 0, size),
		stackPtr: 0,
	}
}

func (s *Stack[T]) Size() int {
	return s.stackPtr
}

func (s *Stack[T]) Print() {
	log.Println("stack value: ", s.store)
	log.Println("stack size: ", s.Size())
}

func (s *Stack[T]) Push(items ...T) {
	if len(items)+len(s.store) > cap(s.store) {
		doubledCap := (len(items) + len(s.store)) * 2
		newStore := make([]T, len(s.store), doubledCap)
		copy(newStore, s.store)
		s.store = newStore
		println("realloc called: new cap is ", doubledCap)
	}

	for _, item := range items {
		s.store = append(s.store, item)
		s.stackPtr++
	}
}

func (s *Stack[T]) Pop(count ...int) []T {
	var popCount int
	if len(count) == 0 {
		popCount = 1
	} else {
		popCount = count[0]
	}

	prevLen := len(s.store)
	if popCount > prevLen {
		popCount = prevLen
	}
	popped := s.store[s.stackPtr-popCount:]
	s.store = s.store[:s.stackPtr-popCount]
	s.stackPtr -= popCount

	return popped
}
