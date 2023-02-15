package testcases

// Test cases for IsEmpty.
var IsEmpty = []struct {
	Name   string
	Data   []int
	Head   int
	Tail   int
	Count  int
	Expect bool
}{
	{Name: "empty buffer", Data: []int{0, 0, 0}, Head: 0, Tail: 0, Count: 0, Expect: true},
	{Name: "partially full buffer", Data: []int{1, 0, 0}, Head: 0, Tail: 1, Count: 1, Expect: false},
	{Name: "full buffer", Data: []int{1, 2, 3}, Head: 0, Tail: 0, Count: 3, Expect: false},
}
