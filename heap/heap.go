package heap

import (
	"errors"
)

type Heap[T any] struct {
	data     []T
	lessFunc func(a, b T) bool
}

// create, heapify, push, pop, peek, size

/*
NewHeap creates a new empty heap with the provided less function.
Can create a min-heap or max-heap depending on the lessFunc provided.
*/
func NewHeap[T any](lessFunc func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		data:     make([]T, 0),
		lessFunc: lessFunc,
	}
}

/*
Heapify builds a heap from an existing slice of data.
*/
func (h *Heap[T]) Heapify(data []T) {
	h.data = data

	// Build the heap property
	for i := len(data)/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
}

/*
Push adds a new element to the heap.
*/
func (h *Heap[T]) Push(value T) {
	h.data = append(h.data, value)
	// Sift up (min-heap): swap while child < parent
	curIndex := len(h.data) - 1
	for curIndex > 0 {
		p := parent(curIndex)
		if h.lessFunc(h.data[curIndex], h.data[p]) {
			h.data[curIndex], h.data[p] = h.data[p], h.data[curIndex]
			curIndex = p
		} else {
			break
		}
	}
}

/*
Pop removes and returns the top element from the heap.
Returns an error if the heap is empty.
*/
func (h *Heap[T]) Pop() (T, error) {
	if len(h.data) == 0 {
		var zero T
		return zero, errors.New("Heap is empty")
	}
	top := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.siftDown(0)
	return top, nil
}

/*
Peek returns the top element without removing it.
*/
func (h *Heap[T]) Peek() (T, error) {
	if len(h.data) == 0 {
		var zero T
		return zero, errors.New("Heap is empty")
	}
	return h.data[0], nil
}

/*
Size returns the number of elements in the heap.
*/
func (h *Heap[T]) Size() int {
	return len(h.data)
}

// Helpers

/*
Sift down the element at index i to restore the heap property (recursive).
*/
func (h *Heap[T]) siftDown(i int) {
	n := len(h.data)
	l := left(i)
	r := right(i)
	smallest := i

	if l < n && h.lessFunc(h.data[l], h.data[smallest]) {
		smallest = l
	}
	if r < n && h.lessFunc(h.data[r], h.data[smallest]) {
		smallest = r
	}

	if smallest != i {
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		h.siftDown(smallest)
	}
}

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}

func parent(index int) int {
	return (index - 1) / 2
}
