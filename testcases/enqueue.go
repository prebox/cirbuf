package testcases

import "github.com/prebox/cirbuf/errors"

// Test cases for Enqueue.
var Enqueue = []struct {
	Name   string
	Item   int
	Count  int
	Head   int
	Tail   int
	Expect []int
	Err    error
}{
	{Name: "enqueue first", Item: 1, Count: 1, Head: 0, Tail: 1, Expect: []int{1, 0, 0}, Err: nil},
	{Name: "enqueue second", Item: 2, Count: 2, Head: 0, Tail: 2, Expect: []int{1, 2, 0}, Err: nil},
	{Name: "enqueue third", Item: 3, Count: 3, Head: 0, Tail: 0, Expect: []int{1, 2, 3}, Err: nil},
	{Name: "enqueue full", Item: 4, Count: 3, Head: 0, Tail: 0, Expect: []int{1, 2, 3}, Err: errors.ErrorIsFull},
}
