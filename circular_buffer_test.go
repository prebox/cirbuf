package cirbuf

import "testing"

func TestIsFull(t *testing.T) {
	// Test cases for IsFull.
	var TestCasesIsFull = []struct {
		Name   string
		Data   []int
		Head   int
		Tail   int
		Count  int
		Expect bool
	}{
		{Name: "Empty buffer", Data: []int{0, 0, 0}, Head: 0, Tail: -1, Count: 0, Expect: false},
		{Name: "Partially full buffer", Data: []int{1, 0, 0}, Head: 0, Tail: 0, Count: 1, Expect: false},
		{Name: "Full buffer", Data: []int{1, 2, 3}, Head: 0, Tail: 2, Count: 3, Expect: true},
	}

	for _, test := range TestCasesIsFull {
		q := &CircularBuffer[int]{
			data:  test.Data,
			head:  test.Head,
			tail:  test.Tail,
			count: test.Count,
		}

		// Run test case.
		t.Run(test.Name, func(t *testing.T) {
			if got := q.IsFull(); got != test.Expect {
				t.Errorf("Expected %v, but got: %v", test.Expect, got)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	// Test cases for IsEmpty.
	var TestCasesIsEmpty = []struct {
		Name   string
		Data   []int
		Head   int
		Tail   int
		Count  int
		Expect bool
	}{
		{Name: "Empty buffer", Data: []int{0, 0, 0}, Head: 0, Tail: -1, Count: 0, Expect: true},
		{Name: "Partially full buffer", Data: []int{1, 0, 0}, Head: 0, Tail: 0, Count: 1, Expect: false},
		{Name: "Full buffer", Data: []int{1, 2, 3}, Head: 0, Tail: 2, Count: 3, Expect: false},
	}

	for _, test := range TestCasesIsEmpty {
		q := &CircularBuffer[int]{
			data:  test.Data,
			head:  test.Head,
			tail:  test.Tail,
			count: test.Count,
		}

		// Run test case.
		t.Run(test.Name, func(t *testing.T) {
			if got := q.IsEmpty(); got != test.Expect {
				t.Errorf("Expected %v, but got: %v", test.Expect, got)
			}
		})
	}
}

func TestEnqueue(t *testing.T) {
	// Test cases for Enqueue.
	var TestCasesEnqueue = []struct {
		Name       string
		Data       []int
		Head       int
		Tail       int
		Count      int
		Item       int
		ExpectData []int
		Err        error
	}{
		{Name: "Enqueue to empty buffer", Data: []int{0, 0, 0, 0, 0}, Head: 0, Tail: -1, Count: 0, Item: 42, ExpectData: []int{42, 0, 0, 0, 0}, Err: nil},
		{Name: "Enqueue to buffer (tail at index 1)", Data: []int{1, 2, 3, 0, 0}, Head: 0, Tail: 1, Count: 2, Item: 4, ExpectData: []int{1, 2, 4, 0, 0}, Err: nil},
		{Name: "Enqueue to non-full buffer", Data: []int{1, 2, 3, 0, 0}, Head: 0, Tail: 2, Count: 3, Item: 4, ExpectData: []int{1, 2, 3, 4, 0}, Err: nil},
		{Name: "Enqueue to buffer (wrapping)", Data: []int{0, 0, 1, 2, 3}, Head: 2, Tail: 4, Count: 3, Item: 4, ExpectData: []int{4, 0, 1, 2, 3}, Err: nil},
		{Name: "Enqueue to full buffer", Data: []int{1, 2, 3, 4, 5}, Head: 0, Tail: 4, Count: 5, Item: 6, ExpectData: []int{1, 2, 3, 4, 5}, Err: ErrorIsFull},
	}

	for _, test := range TestCasesEnqueue {
		q := &CircularBuffer[int]{
			data:  test.Data,
			head:  test.Head,
			tail:  test.Tail,
			count: test.Count,
		}

		// Run test case.
		t.Run(test.Name, func(t *testing.T) {
			err := q.Enqueue(test.Item)
			if err != test.Err {
				t.Errorf("Expected %v, but got: %v", test.Err, err)
			}
			for i := range q.data {
				if q.data[i] != test.ExpectData[i] {
					t.Errorf("Expected %v, but got: %v", test.ExpectData, q.data)
				}
			}
		})
	}
}

func TestDequeue(t *testing.T) {
	// Test cases for Dequeue.
	var TestCasesDequeue = []struct {
		Name         string
		Data         []int
		Head         int
		Tail         int
		Count        int
		ExpectOutput int
		Err          error
	}{
		{Name: "Dequeue from empty buffer", Data: []int{0, 0, 0, 0, 0}, Head: 0, Tail: -1, Count: 0, ExpectOutput: 0, Err: ErrorIsEmpty},
		{Name: "Dequeue from buffer (head at index 2)", Data: []int{1, 2, 3, 4, 0}, Head: 2, Tail: 3, Count: 4, ExpectOutput: 3, Err: nil},
		{Name: "Dequeue from non-empty buffer", Data: []int{1, 2, 3, 4, 0}, Head: 0, Tail: 3, Count: 4, ExpectOutput: 1, Err: nil},
		{Name: "Dequeue from buffer (wrapping)", Data: []int{2, 3, 0, 0, 1}, Head: 4, Tail: 1, Count: 3, ExpectOutput: 1, Err: nil},
		{Name: "Dequeue from full buffer", Data: []int{1, 2, 3, 4, 5}, Head: 0, Tail: 4, Count: 5, ExpectOutput: 1, Err: nil},
	}

	for _, test := range TestCasesDequeue {
		q := &CircularBuffer[int]{
			data:  test.Data,
			head:  test.Head,
			tail:  test.Tail,
			count: test.Count,
		}

		// Run test case.
		t.Run(test.Name, func(t *testing.T) {
			item, err := q.Dequeue()
			if err != test.Err {
				t.Errorf("Expected %v, but got: %v", test.Err, err)
			}
			if item != test.ExpectOutput {
				t.Errorf("Expected %v, but got: %v", test.ExpectOutput, item)
			}
		})
	}
}

func TestGet(t *testing.T) {
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
		{Name: "Get empty", Data: []int{0, 0, 0}, Head: 0, Tail: -1, Count: 0, Index: 1, Expect: 0, Err: ErrorIsEmpty},
		{Name: "Get first", Data: []int{1, 2, 3}, Head: 0, Tail: 2, Count: 3, Index: 0, Expect: 1, Err: nil},
		{Name: "Get second", Data: []int{1, 2, 3}, Head: 0, Tail: 2, Count: 3, Index: 1, Expect: 2, Err: nil},
		{Name: "Get third", Data: []int{1, 2, 3}, Head: 0, Tail: 2, Count: 3, Index: 2, Expect: 3, Err: nil},
		{Name: "Get first (wrapping)", Data: []int{2, 3, 1}, Head: 2, Tail: 1, Count: 3, Index: 0, Expect: 1, Err: nil},
		{Name: "Get second (wrapping)", Data: []int{2, 3, 1}, Head: 2, Tail: 1, Count: 3, Index: 1, Expect: 2, Err: nil},
		{Name: "Get third (wrapping)", Data: []int{2, 3, 1}, Head: 2, Tail: 1, Count: 3, Index: 2, Expect: 3, Err: nil},
		{Name: "Get out of bounds", Data: []int{1, 2, 3}, Head: 0, Tail: 2, Count: 3, Index: 3, Expect: 0, Err: ErrorOutOfBounds},
	}

	for _, test := range TestCasesGet {
		q := &CircularBuffer[int]{
			data:  test.Data,
			head:  test.Head,
			tail:  test.Tail,
			count: test.Count,
		}

		// Run test case.
		t.Run(test.Name, func(t *testing.T) {
			item, err := q.Get(test.Index)
			if err != test.Err {
				t.Errorf("Expected %v, but got: %v", test.Err, err)
			}
			if item != test.Expect {
				t.Errorf("Expected %v, but got: %v", test.Expect, item)
			}
		})
	}
}

func TestPeekFront(t *testing.T) {
	// Test cases for PeekFront.
	var TestCasesPeekFront = []struct {
		Name   string
		Data   []int
		Head   int
		Tail   int
		Count  int
		Expect int
		Err    error
	}{
		{Name: "PeekFront empty buffer", Data: []int{0, 0, 0, 0, 0}, Head: 0, Tail: -1, Count: 0, Expect: 0, Err: ErrorIsEmpty},
		{Name: "PeekFront buffer (head at index 2)", Data: []int{1, 2, 3, 4, 5}, Head: 2, Tail: 2, Count: 3, Expect: 3, Err: nil},
		{Name: "PeekFront non-empty buffer", Data: []int{1, 2, 3, 0, 0}, Head: 0, Tail: 2, Count: 3, Expect: 1, Err: nil},
		{Name: "PeekFront full buffer", Data: []int{1, 2, 3, 4, 5}, Head: 0, Tail: 4, Count: 5, Expect: 1, Err: nil},
	}

	for _, test := range TestCasesPeekFront {
		q := &CircularBuffer[int]{
			data:  test.Data,
			head:  test.Head,
			tail:  test.Tail,
			count: test.Count,
		}

		// Run test case.
		t.Run(test.Name, func(t *testing.T) {
			item, err := q.PeekFront()
			if err != test.Err {
				t.Errorf("Expected %v, but got: %v", test.Err, err)
			}
			if item != test.Expect {
				t.Errorf("Expected %v, but got: %v", test.Expect, item)
			}
		})
	}
}

func TestPeekBack(t *testing.T) {
	// Test cases for PeekBack.
	var TestCasesPeekBack = []struct {
		Name   string
		Data   []int
		Head   int
		Tail   int
		Count  int
		Expect int
		Err    error
	}{
		{Name: "PeekBack empty buffer", Data: []int{0, 0, 0, 0, 0}, Head: 0, Tail: -1, Count: 0, Expect: 0, Err: ErrorIsEmpty},
		{Name: "PeekBack buffer (tail at index 3)", Data: []int{1, 2, 3, 4, 5}, Head: 0, Tail: 2, Count: 3, Expect: 3, Err: nil},
		{Name: "PeekBack non-full buffer", Data: []int{1, 2, 3, 0, 0}, Head: 0, Tail: 2, Count: 3, Expect: 3, Err: nil},
		{Name: "PeekBack full buffer", Data: []int{1, 2, 3, 4, 5}, Head: 0, Tail: 4, Count: 5, Expect: 5, Err: nil},
	}

	for _, test := range TestCasesPeekBack {
		q := &CircularBuffer[int]{
			data:  test.Data,
			head:  test.Head,
			tail:  test.Tail,
			count: test.Count,
		}

		// Run test case.
		t.Run(test.Name, func(t *testing.T) {
			item, err := q.PeekBack()
			if err != test.Err {
				t.Errorf("Expected %v, but got: %v", test.Err, err)
			}
			if item != test.Expect {
				t.Errorf("Expected %v, but got: %v", test.Expect, item)
			}
		})
	}
}

func TestReset(t *testing.T) {
	// Test cases for Reset.
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

	for _, test := range TestCasesReset {
		q := &CircularBuffer[int]{
			data:  test.Data,
			head:  test.Head,
			tail:  test.Tail,
			count: test.Count,
		}

		// Run test case.
		t.Run(test.Name, func(t *testing.T) {
			q.Reset()
			if q.head != test.ExpectHead {
				t.Errorf("Expected %v, but got: %v", test.ExpectHead, q.head)
			}
			if q.tail != test.ExpectTail {
				t.Errorf("Expected %v, but got: %v", test.ExpectTail, q.tail)
			}
			if q.count != test.ExpectCount {
				t.Errorf("Expected %v, but got: %v", test.ExpectCount, q.count)
			}
		})
	}
}

func TestCapacity(t *testing.T) {
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

	for _, test := range TestCasesCapacity {
		q := &CircularBuffer[int]{
			data:  test.Data,
			head:  test.Head,
			tail:  test.Tail,
			count: test.Count,
		}

		// Run test case.
		t.Run(test.Name, func(t *testing.T) {
			if got := q.Capacity(); got != test.Expect {
				t.Errorf("Expected %v, but got: %v", test.Expect, got)
			}
		})
	}
}

func TestCount(t *testing.T) {
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

	for _, test := range TestCasesCount {
		q := &CircularBuffer[int]{
			data:  test.Data,
			head:  test.Head,
			tail:  test.Tail,
			count: test.Count,
		}

		// Run test case.
		t.Run(test.Name, func(t *testing.T) {
			if got := q.Count(); got != test.Expect {
				t.Errorf("Expected %v, but got: %v", test.Expect, got)
			}
		})
	}
}

func TestNew(t *testing.T) {
	// Test cases for New.
	var TestCasesNew = []struct {
		Name     string
		Capacity int
		Expect   int
		Err      error
	}{
		{Name: "New circular buffer with capacity 1", Capacity: 1, Expect: 1, Err: nil},
		{Name: "New circular buffer with capacity 5", Capacity: 5, Expect: 5, Err: nil},
		{Name: "New circular buffer with capacity 0", Capacity: 0, Expect: 1, Err: ErrorInvalidCapacity},
		{Name: "New circular buffer with negative capacity", Capacity: -1, Expect: 1, Err: ErrorInvalidCapacity},
	}

	for _, test := range TestCasesNew {
		q, err := New[int](test.Capacity)

		// Run test case.
		t.Run(test.Name, func(t *testing.T) {
			if err != test.Err {
				t.Errorf("Expected %v, but got: %v", test.Err, err)
			}
			if err == nil {
				if got := len(q.data); got != test.Expect {
					t.Errorf("Expected %v, but got: %v", test.Expect, got)
				}
			}
		})
	}
}
