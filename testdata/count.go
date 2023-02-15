package testdata

// Test cases for Count.
var TestCasesCount = []struct {
	Name   string
	Data   []int
	Head   int
	Tail   int
	Count  int
	Expect int
}{
	{Name: "Empty buffer", Data: []int{0, 0, 0, 0, 0}, Head: 0, Tail: -1, Count: 0, Expect: 0},
	{Name: "Three items in buffer", Data: []int{1, 2, 3, 0, 0}, Head: 0, Tail: 2, Count: 3, Expect: 3},
	{Name: "Five items in buffer (full)", Data: []int{1, 2, 3, 4, 5}, Head: 0, Tail: 4, Count: 5, Expect: 5},
	{Name: "Three items in buffer (wrapped)", Data: []int{3, 0, 0, 1, 2}, Head: 3, Tail: 0, Count: 3, Expect: 3},
}
