// Package cirbuf provides a robust generic static circular buffer.
package cirbuf

import "errors"

var (
	ErrorIsFull          = errors.New("queue is full")
	ErrorIsEmpty         = errors.New("queue is empty")
	ErrorOutOfBounds     = errors.New("index out of bounds")
	ErrorInvalidCapacity = errors.New("invalid capacity")
)

// Represents a circular buffer.
type CircularBuffer[T any] struct {
	data              []T
	head, tail, count int
}

// Creates a new static size circular buffer.
func New[T any](size int) (*CircularBuffer[T], error) {
	if size < 1 {
		return nil, ErrorInvalidCapacity
	}
	return &CircularBuffer[T]{
		data:  make([]T, size),
		head:  0,
		tail:  -1,
		count: 0,
	}, nil
}

// Returns the number of items in the circular buffer.
func (q *CircularBuffer[T]) Count() int {
	return q.count
}

// Enqueues an item to the back of the circular buffer.
func (q *CircularBuffer[T]) Enqueue(item T) error {
	if q.IsFull() {
		return ErrorIsFull
	} else if q.tail += 1; q.tail == len(q.data) {
		q.tail = 0
	}
	q.data[q.tail] = item
	q.count++
	return nil
}

// Dequeues an item from the front of the circular buffer.
func (q *CircularBuffer[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		return *new(T), ErrorIsEmpty
	}
	item := q.data[q.head]
	if q.head += 1; q.head == len(q.data) {
		q.head = 0
	}
	q.count--
	return item, nil
}

// Returns the item at the index in the circular buffer.
func (q *CircularBuffer[T]) Get(index int) (T, error) {
	if q.IsEmpty() {
		return *new(T), ErrorIsEmpty
	} else if index < 0 || index >= q.count {
		return *new(T), ErrorOutOfBounds
	} else if index += q.head; index >= len(q.data) {
		index -= len(q.data)
	}
	return q.data[index], nil
}

// Returns the item at the front of the circular buffer.
func (q *CircularBuffer[T]) PeekFront() (T, error) {
	if q.IsEmpty() {
		return *new(T), ErrorIsEmpty
	}
	return q.data[q.head], nil
}

// Returns the item at the back of the circular buffer.
func (q *CircularBuffer[T]) PeekBack() (T, error) {
	if q.IsEmpty() {
		return *new(T), ErrorIsEmpty
	}
	return q.data[q.tail], nil
}

// Returns true if the circular buffer is empty.
func (q *CircularBuffer[T]) IsEmpty() bool {
	return q.count == 0
}

// Returns true if the circular buffer is full.
func (q *CircularBuffer[T]) IsFull() bool {
	return q.count == len(q.data)
}

// Resets the circular buffer.
func (q *CircularBuffer[T]) Reset() {
	q.head = 0
	q.tail = -1
	q.count = 0
}

// Returns the capacity of the circular buffer.
func (q *CircularBuffer[T]) Capacity() int {
	return len(q.data)
}
