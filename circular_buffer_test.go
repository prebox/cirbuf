package cirbuf

import (
	"testing"

	"github.com/prebox/cirbuf/testdata"
)

func TestIsFull(t *testing.T) {
	for _, test := range testdata.TestCasesIsFull {
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
	for _, test := range testdata.TestCasesIsEmpty {
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
	for _, test := range testdata.TestCasesEnqueue {
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
	for _, test := range testdata.TestCasesDequeue {
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
	for _, test := range testdata.TestCasesGet {
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
	for _, test := range testdata.TestCasesPeekFront {
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
	for _, test := range testdata.TestCasesPeekBack {
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
	for _, test := range testdata.TestCasesReset {
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
