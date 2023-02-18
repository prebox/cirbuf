package errors

import "errors"

var (
	ErrorIsFull          = errors.New("queue is full")
	ErrorIsEmpty         = errors.New("queue is empty")
	ErrorOutOfBounds     = errors.New("index out of bounds")
	ErrorInvalidCapacity = errors.New("invalid capacity")
)
