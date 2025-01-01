package helpers

import (
	"sync"
)

type Queue[T any] struct {
	mutex sync.Mutex
	cond  *sync.Cond
	items []*T
}

func NewQueue[T any]() *Queue[T] {
	q := &Queue[T]{}
	q.cond = sync.NewCond(&q.mutex)
	return q
}

func (q *Queue[T]) Enqueue(value *T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.items = append(q.items, value)

	q.cond.Signal()
}

func (q *Queue[T]) Dequeue() *T {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	for q.unsafeIsEmpty() {
		q.cond.Wait()
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue[T]) Size() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return len(q.items)
}

func (q *Queue[T]) unsafeIsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) IsEmpty() bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.unsafeIsEmpty()
}
