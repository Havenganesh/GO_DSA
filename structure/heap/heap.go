package main

import (
	"fmt"
)

// MaxHeap struct has a slice that holds the heap
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

// Extract removes the largest element from the heap
func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		fmt.Println("Heap is empty")
		return -1
	}

	extracted := h.array[0]
	lastIndex := len(h.array) - 1

	// Move the last item to the root
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]

	h.heapifyDown(0)
	return extracted
}

// heapifyUp maintains the max heap property by moving the element up
func (h *MaxHeap) heapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// heapifyDown maintains the max heap property by moving the element down
func (h *MaxHeap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		// Swap if the child is greater than the current node
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// parent, left, right, and swap functions to support heap operations
func parent(i int) int { return (i - 1) / 2 }
func left(i int) int   { return 2*i + 1 }
func right(i int) int  { return 2*i + 2 }

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println("Building the heap...")
	m.Insert(10)
	m.Insert(20)
	m.Insert(5)
	m.Insert(30)
	m.Insert(15)

	fmt.Println("Heap array:", m.array) // Output: [30 20 5 10 15]

	fmt.Println("Extracted:", m.Extract())               // Output: 30
	fmt.Println("Heap array after extraction:", m.array) // Output: [20 15 5 10]
}
