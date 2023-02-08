# Circular Buffer Package for Go
The cirbuf package is a generic implementation of a circular buffer in Go.
## Features
* Supports any single data type through the use of generics.
* Efficient error handling for a variety of different exceptions.
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
To create a new circular buffer, you can use the `New` function, followed by the type of data you want to store in square brackets. The function argument is the circular buffer's capacity.
```go
buffer := cirbuf.New[int](10)
```
## Available Methods
Here is a list of methods available for the circular buffer:
* `Count() int` - Returns the number of items in the queue.
* `Enqueue(value T) error` - Enqueues an item to the back of the queue.
* `Dequeue() (*T, error)` - Dequeues an item from the front of the queue.
* `Get(index int) (*T, error)` - Returns the item at a specified index in the queue.
* `PeekFront() (*T, error)` - Returns the item at the front of the queue.
* `PeekBack() (*T, error)` - Returns the item at the back of the queue.
* `IsEmpty() bool` - Returns true if the queue is empty.
* `IsFull() bool` - Returns true if the queue is full.
* `Reset()` - Resets the circular buffer.
## Code Example
Here is an example of how you can use the cirbuf package to create a circular buffer and perform various operations on it:  
`NOTE: Methods return pointers to queued items.`
```go
package main

import "github.com/prebox/cirbuf"

func main() {
	buffer := cirbuf.New[int](10)

	buffer.Enqueue(1)
	buffer.Enqueue(2)
	buffer.Enqueue(3)

	fmt.Println(buffer.PeekFront())
	// Output: (*1, nil)

	fmt.Println(buffer.PeekBack())
	// Output: (*3, nil)

	fmt.Println(buffer.Dequeue())
	// Output: (*1, nil)

	fmt.Println(buffer.Get(0))
	// Output: (*2, nil)

	fmt.Println(buffer.IsEmpty())
	// Output: false

	fmt.Println(buffer.IsFull())
	// Output: false

	fmt.Println(buffer.Count())
	// Output: 2

	buffer.Reset()
	
	fmt.Println(buffer.Count())
	// Output: 0
}
```
