package testdata

import "github.com/prebox/cirbuf/errors"

var TestCasesPeekFront = []struct {
	Name   string
	Data   []int
	Head   int
	Tail   int
	Count  int
	Expect int
	Err    error
}{
	{Name: "Peek empty buffer", Data: []int{0, 0, 0, 0, 0}, Head: 0, Tail: 0, Count: 0, Expect: 0, Err: errors.ErrorIsEmpty},
	{Name: "Peek buffer (head at index 2)", Data: []int{1, 2, 3, 0, 0}, Head: 2, Tail: 3, Count: 3, Expect: 3, Err: nil},
	{Name: "Peek non-empty buffer", Data: []int{1, 2, 3, 0, 0}, Head: 0, Tail: 3, Count: 3, Expect: 1, Err: nil},
	{Name: "Peek full buffer", Data: []int{1, 2, 3, 4, 5}, Head: 0, Tail: 0, Count: 5, Expect: 1, Err: nil},
}
