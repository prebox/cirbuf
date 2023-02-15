package testcases

import "github.com/prebox/cirbuf/errors"

// Test cases for Dequeue.
var Dequeue = []struct {
	Name   string
	Count  int
	Head   int
	Tail   int
	Expect int
	Err    error
}{
	{Name: "dequeue first", Count: 2, Head: 1, Tail: 0, Expect: 1, Err: nil},
	{Name: "dequeue second", Count: 1, Head: 2, Tail: 0, Expect: 2, Err: nil},
	{Name: "dequeue third", Count: 0, Head: 0, Tail: 0, Expect: 3, Err: nil},
	{Name: "dequeue empty", Count: 0, Head: 0, Tail: 0, Expect: 0, Err: errors.ErrorIsEmpty},
}
