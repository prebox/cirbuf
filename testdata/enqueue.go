package testdata

import "github.com/prebox/cirbuf/errors"

// Test cases for Enqueue.
var TestCasesEnqueue = []struct {
	Name       string
	Data       []int
	Head       int
	Tail       int
	Count      int
	Item       int
	ExpectData []int
	Err        error
}{
	{Name: "Enqueue to empty buffer", Data: []int{0, 0, 0, 0, 0}, Head: 0, Tail: -1, Count: 0, Item: 42, ExpectData: []int{42, 0, 0, 0, 0}, Err: nil},
	{Name: "Enqueue to buffer (tail at index 1)", Data: []int{1, 2, 3, 0, 0}, Head: 0, Tail: 1, Count: 2, Item: 4, ExpectData: []int{1, 2, 4, 0, 0}, Err: nil},
	{Name: "Enqueue to non-full buffer", Data: []int{1, 2, 3, 0, 0}, Head: 0, Tail: 2, Count: 3, Item: 4, ExpectData: []int{1, 2, 3, 4, 0}, Err: nil},
	{Name: "Enqueue to buffer (wrapping)", Data: []int{0, 0, 1, 2, 3}, Head: 2, Tail: 4, Count: 3, Item: 4, ExpectData: []int{4, 0, 1, 2, 3}, Err: nil},
	{Name: "Enqueue to full buffer", Data: []int{1, 2, 3, 4, 5}, Head: 0, Tail: 4, Count: 5, Item: 6, ExpectData: []int{1, 2, 3, 4, 5}, Err: errors.ErrorIsFull},
}
