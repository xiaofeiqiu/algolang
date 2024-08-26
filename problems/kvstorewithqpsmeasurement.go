package main

import (
	"sync"
	"time"
)

/*
put(key, val)
get(key)
measure_put_load() # return the put qps in last 5 mins measure_get_load() # return the get qps in last 5 mins
# follow up:
# 1. how to test your code
# 2. what if we want to measure qps for last 1 hr, or 1 day?
*/

const windowSize = 5 * 60

type KVNode struct {
	key  int
	val  int
	next *KVNode
}

type HashMap struct {
	buckets      []*KVNode
	size         int
	putCount     []int
	getlock      sync.Mutex
	putlock      sync.Mutex
	putCountLock sync.Mutex
	lastPutTime  int
}

func NewHashMap(size int) HashMap {
	return HashMap{
		buckets:  make([]*KVNode, size),
		size:     size,
		putCount: make([]int, windowSize),
	}
}

func (r *HashMap) hash(key int) int {
	return key % r.size
}

func (r *HashMap) Put(key int, val int) {
	r.putlock.Lock()
	defer r.putlock.Unlock()

	r.doPut(key, val)

	r.putCountLock.Lock()
	defer r.putCountLock.Unlock()

	r.updatePutCount()
	r.putCount[r.lastPutTime%windowSize]++
}

func (r *HashMap) doPut(key int, val int) {
	index := r.hash(key)
	if r.buckets[index] == nil {
		r.buckets[index] = &KVNode{key: key, val: val}
	} else {

		curr := r.buckets[index]
		for {
			if curr.key == key {
				curr.val = val
				return
			}

			if curr.next == nil {
				break
			}

			curr = curr.next
		}

		curr.next = &KVNode{key: key, val: val}
	}
}

func (r *HashMap) Get(key int) int {
	r.getlock.Lock()
	defer r.getlock.Unlock()

	res := r.doGet(key)
	return res
}

func (r *HashMap) doGet(key int) int {
	index := r.hash(key)

	if r.buckets[index] != nil {
		curr := r.buckets[index]
		for curr != nil {
			if curr.key == key {
				return curr.val
			}
			curr = curr.next
		}
	}
	return -1
}

func (r *HashMap) updatePutCount() {
	currentMin := int(time.Now().Unix() / 60)
	diff := currentMin - r.lastPutTime
	for i := 1; i <= diff && i < windowSize; i++ {
		r.putCount[(r.lastPutTime+i)%windowSize] = 0
	}
	r.lastPutTime = int(time.Now().Unix())
}

func (r *HashMap) MeasurePutLoad() int {
	r.putCountLock.Lock()
	defer r.putCountLock.Unlock()

	r.updatePutCount()
	currentMin := int(time.Now().Unix())
	sum := 0
	for i := 0; i < windowSize; i++ {
		index := r.putCount[(currentMin-i+windowSize)%windowSize]
		sum += r.putCount[index]
	}

	return sum / windowSize * 60
}
