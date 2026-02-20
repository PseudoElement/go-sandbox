package main

type Queue[T comparable] struct {
	queue []T
}

func NewQueue[T comparable](length int8) *Queue[T] {
	return &Queue[T]{
		queue: make([]T, length),
	}
}

func (q *Queue[T]) Enqueue(el T) {
	q.queue = append(q.queue, el)
}

func (q *Queue[T]) Dequeue() (first T, ok bool) {
	if q.Size() == 0 {
		return *new(T), false
	}
	first = q.queue[0]
	q.queue = q.queue[1:]
	return first, true
}

func (q *Queue[T]) Size() int {
	return len(q.queue)
}
