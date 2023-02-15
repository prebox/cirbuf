package testdata

import "github.com/prebox/cirbuf/errors"

// Test cases for Dequeue.
var TestCasesDequeue = []struct {
	Name         string
	Data         []int
	Head         int
	Tail         int
	Count        int
	ExpectOutput int
	Err          error
}{
	{Name: "Dequeue from empty buffer", Data: []int{0, 0, 0, 0, 0}, Head: 0, Tail: 0, Count: 0, ExpectOutput: 0, Err: errors.ErrorIsEmpty},
	{Name: "Dequeue from buffer (head at index 2)", Data: []int{1, 2, 3, 4, 0}, Head: 2, Tail: 0, Count: 4, ExpectOutput: 3, Err: nil},
	{Name: "Dequeue from non-empty buffer", Data: []int{1, 2, 3, 4, 0}, Head: 0, Tail: 0, Count: 4, ExpectOutput: 1, Err: nil},
	{Name: "Dequeue from full buffer", Data: []int{1, 2, 3, 4, 5}, Head: 0, Tail: 0, Count: 5, ExpectOutput: 1, Err: nil},
}
