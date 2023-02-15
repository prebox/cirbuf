package testdata

import "github.com/prebox/cirbuf/errors"

var TestCasesPeekBack = []struct {
	Name   string
	Data   []int
	Head   int
	Tail   int
	Count  int
	Expect int
	Err    error
}{
	{Name: "PeekBack empty buffer", Data: []int{0, 0, 0, 0, 0}, Head: 0, Tail: -1, Count: 0, Expect: 0, Err: errors.ErrorIsEmpty},
	{Name: "PeekBack buffer (tail at index 3)", Data: []int{1, 2, 3, 4, 5}, Head: 0, Tail: 2, Count: 3, Expect: 3, Err: nil},
	{Name: "PeekBack non-full buffer", Data: []int{1, 2, 3, 0, 0}, Head: 0, Tail: 2, Count: 3, Expect: 3, Err: nil},
	{Name: "PeekBack full buffer", Data: []int{1, 2, 3, 4, 5}, Head: 0, Tail: 4, Count: 5, Expect: 5, Err: nil},
}
