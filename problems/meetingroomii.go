package main

import (
	"container/heap"
	"sort"
)

type PQ []int

func (q PQ) Len() int            { return len(q) }
func (q PQ) Less(i, j int) bool  { return q[i] < q[j] }
func (q PQ) Swap(i, j int)       { q[i], q[j] = q[j], q[i] }
func (q *PQ) Push(i interface{}) { *q = append(*q, i.(int)) }
func (q *PQ) Pop() interface{} {
	res := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return res
}

func minMeetingRooms(intervals [][]int) int {
	if intervals == nil {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	pq := make(PQ, 0)
	heap.Init(&pq)
	res := -1
	for _, v := range intervals {
		heap.Push(&pq, v[1])

		for v[0] >= pq[0] {
			heap.Pop(&pq)
		}

		res = max(res, pq.Len())
	}

	return res
}
