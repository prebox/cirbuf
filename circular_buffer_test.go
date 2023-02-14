package cirbuf

import "testing"

func TestEnqueue(t *testing.T) {
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

	q := New[int](3)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := q.Enqueue(tc.item)
			if q.count != tc.count {
				t.Errorf("Expected %d, but got: %d", tc.count, q.count)
			}
			if q.head != tc.head {
				t.Errorf("Expected %d, but got: %d", tc.head, q.head)
			}
			if q.tail != tc.tail {
				t.Errorf("Expected %d, but got: %d", tc.tail, q.tail)
			}
			for j := range tc.expected {
				if q.data[j] != tc.expected[j] {
					t.Errorf("Expected %v, but got: %v", tc.expected, q.data)
				}
			}
			if err != tc.err {
				t.Errorf("Expected %v, but got: %v", tc.err, err)
			}
		})
	}
}

func TestDequeue(t *testing.T) {
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

	q := New[int](3)

	// Manually set the state of the buffer.
	q.data = []int{1, 2, 3}
	q.tail = 0
	q.count = 3

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
