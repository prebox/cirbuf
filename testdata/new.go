package testdata

import "github.com/prebox/cirbuf/errors"

// Test cases for New.
var TestCasesNew = []struct {
	Name     string
	Capacity int
	Expect   int
	Err      error
}{
	{Name: "New circular buffer with capacity 1", Capacity: 1, Expect: 1, Err: nil},
	{Name: "New circular buffer with capacity 5", Capacity: 5, Expect: 5, Err: nil},
	{Name: "New circular buffer with capacity 0", Capacity: 0, Expect: 1, Err: errors.ErrorInvalidCapacity},
	{Name: "New circular buffer with negative capacity", Capacity: -1, Expect: 1, Err: errors.ErrorInvalidCapacity},
}
