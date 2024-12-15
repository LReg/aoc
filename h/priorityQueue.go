package h

// Item repräsentiert ein Element in der Priority Queue.
type Item[T any] struct {
	Content  T
	Priority int
	Index    int
}

// PriorityQueue ist ein generischer Typ, der die heap.Interface implementiert.
type PriorityQueue[T any] []*Item[T]

// Len implementiert heap.Interface.
func (pq PriorityQueue[T]) Len() int { return len(pq) }

// Less implementiert heap.Interface. Kleinere Prioritäten haben Vorrang.
func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swap implementiert heap.Interface.
func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push fügt ein Element hinzu.
func (pq *PriorityQueue[T]) Push(x any) {
	item := x.(*Item[T])
	item.Index = len(*pq)
	*pq = append(*pq, item)
}

// Pop entfernt das Element mit der höchsten Priorität.
func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // Speicher freigeben
	*pq = old[0 : n-1]
	return item
}
