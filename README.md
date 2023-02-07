# cirbuf - Circular Buffer Package for Go

The cirbuf package is an implementation of a circular buffer in Go, making use of generics introduced in Go 1.18.  
The package is meant to be used as a FIFO queue data structure that can store any single data type.

## Features
* Supports any data type through the use of generics
* Efficient enqueuing and dequeuing operations
* Ability to check if the buffer is full or empty
* Peek at the front and back elements without modifying the buffer

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
To create a new circular buffer, you can use the New function, followed by the type of data you want to store in square brackets. The function argument is the size of the buffer.
```go
buffer := cirbuf.New[int](10)
```
## Available Methods
Here is a list of methods available in the circular buffer:
* `Enqueue(value T) error` - Adds an element to the end of the buffer.
* `Dequeue() (*T, error)` - Removes and returns the first element in the buffer.
* `Get(index int) (*T, error)` - Returns the element at the specified index in the buffer.
* `IsEmpty() bool` - Returns true if the buffer is empty, false otherwise.
* `IsFull() bool` - Returns true if the buffer is full, false otherwise.
* `PeekBack() (*T, error)` - Returns the last element in the buffer without modifying it.
* `PeekFront() (*T, error)` - Returns the first element in the buffer without modifying it.
## Code Example
Here is an example of how you can use the cirbuf package to create a circular buffer and perform various operations on it:  
`NOTE: The methods return pointers.`
```go
package main

import "github.com/prebox/cirbuf"

func main() {
	buffer := cirbuf.New[int](10)

	buffer.Enqueue(1)
	buffer.Enqueue(2)
	buffer.Enqueue(3)

	fmt.Println(buffer.PeekFront())
	// Output: 1

	fmt.Println(buffer.PeekBack())
	// Output: 3

	fmt.Println(buffer.Dequeue())
	// Output: 1

	fmt.Println(buffer.Get(0))
	// Output: 2

	fmt.Println(buffer.IsEmpty())
	// Output: false

	fmt.Println(buffer.IsFull())
	// Output: false
}
```
In this example, we initialize a circular buffer with type int and a capacity of 10. Then we enqueue the int `1, 2, 3` and perform the operations `PeekFront` and `PeekBack` to verify that the order is correct. We then use `Dequeue` to dequeue the value `1` and verify that it has been successfully removed using the `Get` method on index `0`. Additionally `IsEmpty` and `IsFull` both return false since the buffer is neither full nor empty.
