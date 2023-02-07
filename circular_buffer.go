package cirbuf

import "errors"

var (
	ErrorIsFull      = errors.New("queue is full")
	ErrorIsEmpty     = errors.New("queue is empty")
	ErrorOutOfBounds = errors.New("index out of bounds")
)

// Represents a circular buffer/queue.
type CircularBuffer[T any] struct {
	data              []T
	head, tail, count int
}

// Creates a new queue with the specified length.
func New[T any](length int) *CircularBuffer[T] {
	return &CircularBuffer[T]{data: make([]T, length)}
}

// Enqueues an item to the back of the queue.
func (q *CircularBuffer[T]) Enqueue(item T) error {
	if q.IsFull() {
		return ErrorIsFull
	} else {
		q.count++
		q.data[q.tail] = item
		q.tail = (q.tail + 1) % len(q.data)
		return nil
	}
}

// Dequeues an item from the front of the queue.
func (q *CircularBuffer[T]) Dequeue() (*T, error) {
	if q.IsEmpty() {
		return nil, ErrorIsEmpty
	} else {
		q.count--
		item := q.data[q.head]
		q.data[q.head] = *new(T)
		q.head = (q.head + 1) % len(q.data)
		return &item, nil
	}
}

// Returns the item at a specified index in the queue.
func (q *CircularBuffer[T]) Get(index int) (*T, error) {
	if q.IsEmpty() {
		return nil, ErrorIsEmpty
	} else if index < 0 || index >= q.count {
		return nil, ErrorOutOfBounds
	} else {
		return &q.data[(q.head+index)%len(q.data)], nil
	}
}

// Returns the item at the front of the queue.
func (q *CircularBuffer[T]) PeekFront() (*T, error) {
	if q.IsEmpty() {
		return nil, ErrorIsEmpty
	} else {
		return &q.data[q.head], nil
	}
}

// Returns the item at the back of the queue.
func (q *CircularBuffer[T]) PeekBack() (*T, error) {
	if q.IsEmpty() {
		return nil, ErrorIsEmpty
	} else {
		var index int
		if q.tail != 0 {
			index = q.tail - 1
		} else if q.tail == 0 {
			index = len(q.data) - 1
		}
		return &q.data[index], nil
	}
}

// Returns true if the queue is empty.
func (q *CircularBuffer[T]) IsEmpty() bool {
	return q.count == 0
}

// Returns true if the queue is full.
func (q *CircularBuffer[T]) IsFull() bool {
	return q.count == len(q.data)
}
