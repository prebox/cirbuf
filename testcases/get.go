package testcases

import "github.com/prebox/cirbuf/errors"

// Test cases for Get.
var Get = []struct {
	Name   string
	Index  int
	Expect int
	Err    error
}{
	{Name: "get first", Index: 0, Expect: 1, Err: nil},
	{Name: "get second", Index: 1, Expect: 2, Err: nil},
	{Name: "get third", Index: 2, Expect: 3, Err: nil},
	{Name: "get out of bounds", Index: 3, Expect: 0, Err: errors.ErrorOutOfBounds},
}
