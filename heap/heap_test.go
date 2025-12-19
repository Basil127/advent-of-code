package heap_test

import (
	"testing"

	"github.com/basil127/advent-of-code/heap"
)

func TestHeapPushPop(t *testing.T) {
	h := heap.NewHeap(intLess)
	h.Push(3)
	h.Push(1)
	h.Push(2)

	if h.Size() != 3 {
		t.Errorf("Expected size 3, got %d", h.Size())
	}

	if val, err := h.Peek(); err != nil || val != 1 {
		t.Errorf("Expected peek 1, got %d", val)
	}

	if val, err := h.Pop(); err != nil || val != 1 {
		t.Errorf("Expected pop 1, got %d", val)
	}

	if val, err := h.Pop(); err != nil || val != 2 {
		t.Errorf("Expected pop 2, got %d", val)
	}

	if val, err := h.Pop(); err != nil || val != 3 {
		t.Errorf("Expected pop 3, got %d", val)
	}

	if h.Size() != 0 {
		t.Errorf("Expected heap to be empty")
	}
}

func TestHeapHeapify(t *testing.T) {
	h := heap.NewHeap(intLess)
	h.Heapify([]int{5, 3, 8, 1, 2})
	if h.Size() != 5 {
		t.Errorf("Expected size 5, got %d", h.Size())
	}
	expectedOrder := []int{1, 2, 3, 5, 8}
	for _, expected := range expectedOrder {
		if val, err := h.Pop(); err != nil || val != expected {
			t.Errorf("Expected pop %d", expected)
		}
	}
	if h.Size() != 0 {
		t.Errorf("Expected heap to be empty")
	}
}

func TestHeapEmpty(t *testing.T) {
	h := heap.NewHeap(intLess)
	if _, err := h.Pop(); err == nil {
		t.Errorf("Expected error when popping from empty heap")
	}
	if _, err := h.Peek(); err == nil {
		t.Errorf("Expected error when peeking into empty heap")
	}
}

func TestHeapSize(t *testing.T) {
	h := heap.NewHeap(intLess)
	if h.Size() != 0 {
		t.Errorf("Expected size 0, got %d", h.Size())
	}
	h.Push(10)
	if h.Size() != 1 {
		t.Errorf("Expected size 1, got %d", h.Size())
	}
	h.Push(20)
	if h.Size() != 2 {
		t.Errorf("Expected size 2, got %d", h.Size())
	}
	h.Pop()
	if h.Size() != 1 {
		t.Errorf("Expected size 1, got %d", h.Size())
	}
	h.Pop()
	if h.Size() != 0 {
		t.Errorf("Expected size 0, got %d", h.Size())
	}
}

func TestLargeHeap(t *testing.T) {
	h := heap.NewHeap(intLess)
	numElements := 1000
	for i := numElements; i >= 1; i-- {
		h.Push(i)
	}
	if h.Size() != numElements {
		t.Errorf("Expected size %d, got %d", numElements, h.Size())
	}
	for i := 1; i <= numElements; i++ {
		if val, err := h.Pop(); err != nil || val != i {
			t.Errorf("Expected pop %d, got %d", i, val)
		}
	}
	if h.Size() != 0 {
		t.Errorf("Expected heap to be empty")
	}
}

// helper less function for integers
func intLess(a, b int) bool {
	return a < b
}
