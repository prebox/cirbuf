package testdata

import "github.com/prebox/cirbuf/errors"

// Test cases for Get.
var TestCasesGet = []struct {
	Name   string
	Data   []int
	Head   int
	Tail   int
	Count  int
	Index  int
	Expect int
	Err    error
}{
	{Name: "get empty", Data: []int{0, 0, 0}, Head: 0, Tail: -1, Count: 0, Index: 1, Expect: 0, Err: errors.ErrorIsEmpty},
	{Name: "get first", Data: []int{1, 2, 3}, Head: 0, Tail: 2, Count: 3, Index: 0, Expect: 1, Err: nil},
	{Name: "get second", Data: []int{1, 2, 3}, Head: 0, Tail: 2, Count: 3, Index: 1, Expect: 2, Err: nil},
	{Name: "get third", Data: []int{1, 2, 3}, Head: 0, Tail: 2, Count: 3, Index: 2, Expect: 3, Err: nil},
	{Name: "get out of bounds", Data: []int{1, 2, 3}, Head: 0, Tail: 2, Count: 3, Index: 3, Expect: 0, Err: errors.ErrorOutOfBounds},
}
