package cirbuf

import "testing"

func TestIsFull(t *testing.T) {
	q := New[int](3)

	// Check if IsFull returns false for an empty buffer.
	if q.IsFull() != false {
		t.Errorf("Expected false, but got: %v", q.IsFull())
	}

	// Manually set the state of the buffer to partially full.
	// Check if IsFull returns false for a partially full buffer.
	q.data = []int{1, 2, 0}
	q.head = 0
	q.tail = 2
	q.count = 2
	if q.IsFull() != false {
		t.Errorf("Expected false, but got: %v", q.IsFull())
	}

	// Manually set the state of the buffer to full.
	// Check if IsFull returns true for a full buffer.
	q.data = []int{1, 2, 3}
	q.head = 0
	q.tail = 0
	q.count = 3
	if q.IsFull() != true {
		t.Errorf("Expected true, but got: %v", q.IsFull())
	}
}

func TestEnqueue(t *testing.T) {
	q := New[int](3)

	// Setup test cases.
	testCases := []struct {
		name     string
		item     int
		expected []int
		count    int
		head     int
		tail     int
		err      error
	}{
		{"enqueue first", 1, []int{1, 0, 0}, 1, 0, 1, nil},
		{"enqueue second", 2, []int{1, 2, 0}, 2, 0, 2, nil},
		{"enqueue third", 3, []int{1, 2, 3}, 3, 0, 0, nil},
		{"enqueue full", 4, []int{1, 2, 3}, 3, 0, 0, ErrorIsFull},
	}

	// Run test cases.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := q.Enqueue(tc.item)
			for j := range tc.expected {
				if q.data[j] != tc.expected[j] {
					t.Errorf("Expected %v, but got: %v", tc.expected, q.data)
				}
			}
			if q.count != tc.count {
				t.Errorf("Expected %d, but got: %d", tc.count, q.count)
			}
			if q.head != tc.head {
				t.Errorf("Expected %d, but got: %d", tc.head, q.head)
			}
			if q.tail != tc.tail {
				t.Errorf("Expected %d, but got: %d", tc.tail, q.tail)
			}
			if err != tc.err {
				t.Errorf("Expected %v, but got: %v", tc.err, err)
			}
		})
	}
}

func TestDequeue(t *testing.T) {
	q := New[int](3)

	// Manually set the state of the buffer.
	q.data = []int{1, 2, 3}
	q.tail = 0
	q.count = 3

	// Setup test cases.
	testCases := []struct {
		name     string
		expected int
		count    int
		head     int
		tail     int
		err      error
	}{
		{"dequeue first", 1, 2, 1, 0, nil},
		{"dequeue second", 2, 1, 2, 0, nil},
		{"dequeue third", 3, 0, 0, 0, nil},
		{"dequeue empty", 0, 0, 0, 0, ErrorIsEmpty},
	}

	// Run test cases.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			item, err := q.Dequeue()
			if item != tc.expected {
				t.Errorf("Expected %v, but got: %v", tc.expected, item)
			}
			if q.count != tc.count {
				t.Errorf("Expected %v, but got: %d", tc.count, q.count)
			}
			if q.head != tc.head {
				t.Errorf("Expected %v, but got: %d", tc.head, q.head)
			}
			if q.tail != tc.tail {
				t.Errorf("Expected %v, but got: %d", tc.tail, q.tail)
			}
			if err != tc.err {
				t.Errorf("Expected %v, but got: %v", tc.err, err)
			}
		})
	}
}

func TestGet(t *testing.T) {
	q := New[int](3)

	// Run test on empty queue.
	t.Run("get empty", func(t *testing.T) {
		if _, err := q.Get(0); err != ErrorIsEmpty {
			t.Errorf("Expected %v, but got: %v", ErrorIsEmpty, err)
		}
	})

	// Manually set the state of the buffer.
	q.data = []int{1, 2, 3}
	q.tail = 0
	q.count = 3

	// Setup test cases.
	testCases := []struct {
		name     string
		index    int
		expected int
		err      error
	}{
		{"get first", 0, 1, nil},
		{"get second", 1, 2, nil},
		{"get third", 2, 3, nil},
		{"get out of bounds", 3, 0, ErrorOutOfBounds},
	}

	// Run test cases.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			item, err := q.Get(tc.index)
			if item != tc.expected {
				t.Errorf("Expected %v, but got: %v", tc.expected, item)
			}
			if err != tc.err {
				t.Errorf("Expected %v, but got: %v", tc.err, err)
			}
		})
	}
}
