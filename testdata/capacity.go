package testdata

// Test cases for Capacity.
var TestCasesCapacity = []struct {
	Name   string
	Data   []int
	Head   int
	Tail   int
	Count  int
	Expect int
}{
	{Name: "Empty buffer with capacity of 5", Data: []int{0, 0, 0, 0, 0}, Head: 0, Tail: -1, Count: 0, Expect: 5},
	{Name: "Non-empty buffer with capacity of 5", Data: []int{1, 2, 3, 0, 0}, Head: 0, Tail: 2, Count: 3, Expect: 5},
	{Name: "Full buffer with capacity of 5", Data: []int{1, 2, 3, 4, 5}, Head: 0, Tail: 4, Count: 5, Expect: 5},
	{Name: "Empty buffer with capacity of 3", Data: []int{0, 0, 0}, Head: 0, Tail: -1, Count: 0, Expect: 3},
	{Name: "Non-empty buffer with capacity of 3", Data: []int{1, 2, 0}, Head: 0, Tail: 1, Count: 2, Expect: 3},
	{Name: "Full buffer with capacity of 3", Data: []int{1, 2, 3}, Head: 0, Tail: 2, Count: 3, Expect: 3},
}
