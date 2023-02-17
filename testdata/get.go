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
	{Name: "Get empty", Data: []int{0, 0, 0}, Head: 0, Tail: -1, Count: 0, Index: 1, Expect: 0, Err: errors.ErrorIsEmpty},
	{Name: "Get first", Data: []int{1, 2, 3}, Head: 0, Tail: 2, Count: 3, Index: 0, Expect: 1, Err: nil},
	{Name: "Get second", Data: []int{1, 2, 3}, Head: 0, Tail: 2, Count: 3, Index: 1, Expect: 2, Err: nil},
	{Name: "Get third", Data: []int{1, 2, 3}, Head: 0, Tail: 2, Count: 3, Index: 2, Expect: 3, Err: nil},
	{Name: "Get first (wrapping)", Data: []int{2, 3, 1}, Head: 2, Tail: 1, Count: 3, Index: 0, Expect: 1, Err: nil},
	{Name: "Get second (wrapping)", Data: []int{2, 3, 1}, Head: 2, Tail: 1, Count: 3, Index: 1, Expect: 2, Err: nil},
	{Name: "Get third (wrapping)", Data: []int{2, 3, 1}, Head: 2, Tail: 1, Count: 3, Index: 2, Expect: 3, Err: nil},
	{Name: "Get out of bounds", Data: []int{1, 2, 3}, Head: 0, Tail: 2, Count: 3, Index: 3, Expect: 0, Err: errors.ErrorOutOfBounds},
}
