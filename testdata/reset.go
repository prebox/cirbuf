package testdata

var TestCasesReset = []struct {
	Name        string
	Data        []int
	Head        int
	Tail        int
	Count       int
	ExpectHead  int
	ExpectTail  int
	ExpectCount int
}{
	{Name: "Reset full buffer", Data: []int{1, 2, 3, 4, 5}, Head: 0, Tail: 4, Count: 5, ExpectHead: 0, ExpectTail: -1, ExpectCount: 0},
	{Name: "Reset buffer (head at index 2)", Data: []int{1, 2, 3, 4, 5}, Head: 2, Tail: 4, Count: 5, ExpectHead: 0, ExpectTail: -1, ExpectCount: 0},
	{Name: "Reset non-full buffer", Data: []int{1, 2, 3, 0, 0}, Head: 0, Tail: 2, Count: 3, ExpectHead: 0, ExpectTail: -1, ExpectCount: 0},
	{Name: "Reset empty buffer", Data: []int{0, 0, 0, 0, 0}, Head: 0, Tail: -1, Count: 0, ExpectHead: 0, ExpectTail: -1, ExpectCount: 0},
}
