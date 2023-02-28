# Circular Buffer
![Go Test Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/prebox/cirbuf)](https://goreportcard.com/report/github.com/prebox/cirbuf)
[![Go Reference](https://pkg.go.dev/badge/github.com/prebox/cirbuf.svg)](https://pkg.go.dev/github.com/prebox/cirbuf)

The cirbuf package is a robust generic static circular buffer.
## Features
* Extensive tests providing 100% test coverage.
* Supports any single data type through the use of generics.
* Provides O(1) time complexity for all of the circular buffer methods.
## Getting Started
To start using cirbuf, you need to install it by running the following command:
```go
go get github.com/prebox/cirbuf
```
Once you have the package installed, you can start using it by importing it in your code:
```go
import "github.com/prebox/cirbuf"
```
## Initializing a Circular Buffer
To create a new circular buffer, you can use the `New` function, followed by the type of data you want to store in square brackets. The function argument is the circular buffer's capacity. The `New` function will return an error if the capacity is invalid.
```go
buffer, err := cirbuf.New[int](10)
```
## Available Methods
Here is a list of methods available for the circular buffer:
* `Count() int` - Returns the number of items in the queue.
* `Enqueue(value T) error` - Enqueues an item to the back of the queue.
* `Dequeue() (T, error)` - Dequeues an item from the front of the queue.
* `Get(index int) (T, error)` - Returns the item at a specified index in the queue.
* `PeekFront() (T, error)` - Returns the item at the front of the queue.
* `PeekBack() (T, error)` - Returns the item at the back of the queue.
* `IsEmpty() bool` - Returns true if the queue is empty.
* `IsFull() bool` - Returns true if the queue is full.
* `Reset()` - Resets the circular buffer.
* `Capacity() int` - Returns the capacity of the circular buffer.
## Code Example
Here is an example of how you can use the cirbuf package:  
```go
package main

import "github.com/prebox/cirbuf"

func main() {
	buffer, _ := cirbuf.New[int](3)

	buffer.Enqueue(1)
	buffer.Enqueue(2)
	buffer.Enqueue(3)
	buffer.Enqueue(4) 	// Output: "queue is full"
	buffer.Count()    	// Output: 3

	buffer.PeekFront()	// Output: (1, nil)
	buffer.PeekBack() 	// Output: (3, nil)

	buffer.Dequeue()  	// Output: (1, nil)
	buffer.Get(0)     	// Output: (2, nil)
	buffer.Count()    	// Output: 2

	buffer.IsEmpty()  	// Output: false
	buffer.IsFull()   	// Output: false

	buffer.Reset()
	buffer.Count()    	// Output: 0
	buffer.Capacity()  	// Output: 3
}
```
