package main

type MaxHeap []int

func NewMaxHeap(values []int) *MaxHeap {
	maxHeap := make(MaxHeap, len(values))
	copy(maxHeap, values)
	for i := len(maxHeap)/2 - 1; i >= 0; i-- {
		maxHeap.heapify(i)
	}

	return &maxHeap
}

func (h *MaxHeap) heapify(i int) {
	if (*h)[i*2+1] > (*h)[i] {
		temp := (*h)[i]
		(*h)[i] = (*h)[i*2+1]
		(*h)[i*2+1] = temp
	}
	if len(*h) > i*2+2 && (*h)[i*2+2] > (*h)[i] {
		temp := (*h)[i]
		(*h)[i] = (*h)[i*2+2]
		(*h)[i*2+2] = temp
	}
}
