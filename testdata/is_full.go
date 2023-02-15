package testdata

// Test cases for IsFull.
var TestCasesIsFull = []struct {
	Name   string
	Data   []int
	Head   int
	Tail   int
	Count  int
	Expect bool
}{
	{Name: "empty buffer", Data: []int{0, 0, 0}, Head: 0, Tail: 0, Count: 0, Expect: false},
	{Name: "partially full buffer", Data: []int{1, 0, 0}, Head: 0, Tail: 1, Count: 1, Expect: false},
	{Name: "full buffer", Data: []int{1, 2, 3}, Head: 0, Tail: 0, Count: 3, Expect: true},
}
