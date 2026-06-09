package main

import (
	"log"
	"math"
)

func Mul2(x int) int {
	return x * 2
}

func Pow2(x int) int {
	return int(math.Pow(float64(x), float64(2)))
}

func pipePromises() {
	res := NewPromise(func(arg int) int {
		return 6
	}, 0).
		Then(Mul2).
		Then(Pow2).
		Unwrap()

	log.Println("RES ==> ", res)
}

type Thenable[T any] interface {
	Then(fn func(args T) T) Thenable[T]
	Unwrap() T
}

type Promise[T any] struct {
	value T
	err   error
}

func NewPromise[T any](callback func(args T) T, arg T) *Promise[T] {
	val := callback(arg)
	return &Promise[T]{value: val, err: nil}
}

func (p *Promise[T]) Then(fn func(args T) T) Thenable[T] {
	p.err = nil
	return NewPromise(fn, p.value)
}

func (p *Promise[T]) Unwrap() T {
	return p.value
}
