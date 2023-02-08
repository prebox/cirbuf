// Package cirbuf provides a generic implementation of a circular buffer.
package cirbuf

import "errors"

// Represents an error message.
var (
	ErrorIsFull      = errors.New("queue is full")
	ErrorIsEmpty     = errors.New("queue is empty")
	ErrorOutOfBounds = errors.New("index out of bounds")
)

// Represents a circular buffer.
type CircularBuffer[T any] struct {
	data              []T
	head, tail, count int
}

// Creates a new queue with the specified length.
func New[T any](length int) *CircularBuffer[T] {
	return &CircularBuffer[T]{data: make([]T, length)}
}

// Returns the number of items in the queue.
func (q *CircularBuffer[T]) Count() int {
	return q.count
}

// Enqueues an item to the back of the queue.
func (q *CircularBuffer[T]) Enqueue(item T) error {
	if q.IsFull() {
		return ErrorIsFull
	}
	q.data[q.tail] = item
	q.tail = (q.tail + 1) % len(q.data)
	q.count++
	return nil
}

// Dequeues an item from the front of the queue.
func (q *CircularBuffer[T]) Dequeue() (*T, error) {
	if q.IsEmpty() {
		return nil, ErrorIsEmpty
	}
	item := &q.data[q.head]
	q.head = (q.head + 1) % len(q.data)
	q.count--
	return item, nil
}

// Returns the item at a specified index in the queue.
func (q *CircularBuffer[T]) Get(index int) (*T, error) {
	if q.IsEmpty() {
		return nil, ErrorIsEmpty
	} else if index < 0 || index >= q.count {
		return nil, ErrorOutOfBounds
	}
	internalIndex := (q.head + index) % len(q.data)
	return &q.data[internalIndex], nil
}

// Returns the item at the front of the queue.
func (q *CircularBuffer[T]) PeekFront() (*T, error) {
	if q.IsEmpty() {
		return nil, ErrorIsEmpty
	}
	return &q.data[q.head], nil
}

// Returns the item at the back of the queue.
func (q *CircularBuffer[T]) PeekBack() (*T, error) {
	if q.IsEmpty() {
		return nil, ErrorIsEmpty
	}
	return &q.data[(q.tail-1)%len(q.data)], nil
}

// Returns true if the queue is empty.
func (q *CircularBuffer[T]) IsEmpty() bool {
	return q.count == 0
}

// Returns true if the queue is full.
func (q *CircularBuffer[T]) IsFull() bool {
	return q.count == len(q.data)
}

// Resets the circular buffer.
func (q *CircularBuffer[T]) Reset() {
	q.head, q.tail, q.count = 0, 0, 0
}
