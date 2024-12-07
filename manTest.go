package main

import (
	"container/heap"
	"fmt"
)

// An Item is something we manage in a priority queue.
type Item struct {
	Value    string // The value of the item.
	Priority int    // The priority of the item in the queue.
	Index    int    // The index of the item in the heap (managed by the heap.Interface).
}

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest priority item, so we use less-than.
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push adds an item to the priority queue.
func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the lowest priority item from the queue.
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // For safety
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority of an item in the queue.
func (pq *PriorityQueue) Update(item *Item, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

func main() {
	// Create a priority queue and add some items.
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Add some items to the queue.
	heap.Push(&pq, &Item{Value: "task1", Priority: 3})
	heap.Push(&pq, &Item{Value: "task2", Priority: 1})
	heap.Push(&pq, &Item{Value: "task3", Priority: 2})

	// Update an item's priority.
	item := &Item{Value: "task4", Priority: 5}
	heap.Push(&pq, item)
	pq.Update(item, item.Value, 0)

	// Pop all items from the queue.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("Popped: %s with priority %d\n", item.Value, item.Priority)
	}
}
