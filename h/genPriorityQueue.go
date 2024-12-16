package h

import (
	"cmp"
	"container/heap"
	"fmt"
)

// PQ ist eine generische Priority Queue für Typen T.
type PQ[T cmp.Ordered | Point | Node] struct {
	queue PriorityQueue[T]
	index int
}

// NewPC erstellt eine neue Priority Queue.
func NewPC[T cmp.Ordered | Point]() *PQ[T] {
	return &PQ[T]{
		queue: make(PriorityQueue[T], 0),
		index: 0,
	}
}

func (pc *PQ[T]) Push(content T, priority int) {
	item := &Item[T]{
		Content:  content,
		Priority: priority,
		Index:    pc.index,
	}
	pc.index++
	heap.Push(&pc.queue, item)
}

func (pc *PQ[T]) Pop() T {
	if pc.queue.Len() == 0 {
		var zero T
		return zero
	}
	return heap.Pop(&pc.queue).(*Item[T]).Content
}

func (pc *PQ[T]) UpdatePriority(content T, newPriority int) bool {
	for i, item := range pc.queue {
		if item.Content == content {
			item.Priority = newPriority
			heap.Fix(&pc.queue, i)
			return true
		}
	}
	return false
}

func (pc *PQ[T]) Len() int {
	return len(pc.queue)
}

func (pc *PQ[T]) First() *Item[T] {
	if len(pc.queue) == 0 {
		return nil
	}
	return pc.queue[0]
}

func (pc *PQ[T]) Contains(content T) bool {
	for _, item := range pc.queue {
		if item.Content == content {
			return true
		}
	}
	return false
}

func (pc *PQ[T]) Print() {
	fmt.Println(pc.queue)
}
