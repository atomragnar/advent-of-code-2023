package util

import (
	"container/heap"
)

type QueueItem struct {
	Index    int
	Value    interface{}
	Priority int
}

type PriorityQueue struct {
	Q            []QueueItem
	PriorityFunc func(this, other QueueItem) bool
}

func NewQueue(fn func(this, other QueueItem) bool) *PriorityQueue {
	q := PriorityQueue{
		Q:            make([]QueueItem, 0),
		PriorityFunc: fn,
	}
	heap.Init(&q)
	return &q
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.Q[i], pq.Q[j] = pq.Q[j], pq.Q[i]
	pq.Q[i].Index = i
	pq.Q[j].Index = j
}

func (pq *PriorityQueue) Len() int {
	return len(pq.Q)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.PriorityFunc(pq.Q[i], pq.Q[j])
}

func (pq *PriorityQueue) Push(x any) {
	n := len(pq.Q)
	item := x.(QueueItem)
	item.Index = n
	pq.Q = append(pq.Q, item)
}

func (pq *PriorityQueue) Pop() any {
	old := pq.Q
	n := len(old)
	item := old[n-1]
	item.Index = -1 // For safety
	pq.Q = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(itemIndex int, value any, priority int) {
	if itemIndex < 0 || itemIndex >= len(pq.Q) {
		return
	}
	pq.Q[itemIndex].Value = itemIndex
	pq.Q[itemIndex].Value = value
	pq.Q[itemIndex].Priority = priority
	heap.Fix(pq, itemIndex)
}
