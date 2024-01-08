package main

import "container/heap"

type IntHeap []int

func (r IntHeap) Len() int           { return len(r) }
func (r IntHeap) Less(i, j int) bool { return r[i] < r[j] }
func (r IntHeap) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r *IntHeap) Push(v any)        { *r = append(*r, v.(int)) }
func (r *IntHeap) Pop() any          { n := len(*r); res := (*r)[n-1]; *r = (*r)[:n-1]; return res }

type MedianFinder struct {
	minHeap IntHeap
	maxHeap IntHeap
}

func Constructor2() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.maxHeap, -num)
	heap.Push(&this.minHeap, -heap.Pop(&this.maxHeap).(int))
	if this.minHeap.Len() > this.maxHeap.Len() {
		heap.Push(&this.maxHeap, -heap.Pop(&this.minHeap).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() > this.minHeap.Len() {
		return float64(-this.maxHeap[0])
	}

	return float64((-this.maxHeap[0] + this.minHeap[0])) / 2
}
