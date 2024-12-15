package h

import (
	"cmp"
	"container/heap"
	"fmt"
)

// PC ist eine generische Priority Queue für Typen T.
type PC[T cmp.Ordered | Point] struct {
	queue PriorityQueue[T]
	index int
}

// NewPC erstellt eine neue Priority Queue.
func NewPC[T cmp.Ordered | Point]() *PC[T] {
	return &PC[T]{
		queue: make(PriorityQueue[T], 0),
		index: 0,
	}
}

func (pc *PC[T]) Push(content T, priority int) {
	item := &Item[T]{
		Content:  content,
		Priority: priority,
		Index:    pc.index,
	}
	pc.index++
	heap.Push(&pc.queue, item)
}

func (pc *PC[T]) Pop() T {
	if pc.queue.Len() == 0 {
		var zero T
		return zero
	}
	return heap.Pop(&pc.queue).(*Item[T]).Content
}

func (pc *PC[T]) UpdatePriority(content T, newPriority int) bool {
	for i, item := range pc.queue {
		if item.Content == content {
			item.Priority = newPriority
			heap.Fix(&pc.queue, i)
			return true
		}
	}
	return false
}

func (pc *PC[T]) Len() int {
	return len(pc.queue)
}

func (pc *PC[T]) First() *Item[T] {
	if len(pc.queue) == 0 {
		return nil
	}
	return pc.queue[0]
}

func (pc *PC[T]) Contains(content T) bool {
	for _, item := range pc.queue {
		if item.Content == content {
			return true
		}
	}
	return false
}

func (pc *PC[T]) Print() {
	fmt.Println(pc.queue)
}
