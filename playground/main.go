package main

import "fmt"

//Design a snapshot set functionality.
//
//Once the snapshot is taken, the iterator of the class should only return values that were present in the function.
//
//The class should provide add, remove, and contains functionality. The iterator always returns elements that were present in the snapshot even though the element might be removed from set after the snapshot.
//
//The snapshot of the set is taken when the iterator function is called.
//
//interface SnapshotSet {
//  void add(int num);
//  void remove(num);
//  boolean contains(num);
//  Iterator<Integer> iterator(); // the first call to this function should trigger a snapshot of the set
//}

type SnapshotElement struct {
	val       int
	createdAt int
	removedAt int
}

type SnapshotSet struct {
	data map[int]*SnapshotElement
	arr  []*SnapshotElement
	t    int
}

func NewSnapshotSet() SnapshotSet {
	data := make(map[int]*SnapshotElement)
	return SnapshotSet{
		data: data,
		t:    1,
		arr:  make([]*SnapshotElement, 0),
	}
}

func (s *SnapshotSet) Add(v int) {
	if _, exist := s.data[v]; !exist {
		s.data[v] = &SnapshotElement{val: v, createdAt: s.t}
		s.arr = append(s.arr, s.data[v])
	} else {
		if s.data[v].removedAt > 0 {
			s.data[v].createdAt = s.t
		}
	}
	s.t++
}

func (s *SnapshotSet) Remove(v int) {
	if e, exists := s.data[v]; exists {
		e.removedAt = s.t
	}
	s.t++
}

func (s *SnapshotSet) Contains(v int) bool {
	if metadata, exists := s.data[v]; exists {
		return metadata.createdAt > metadata.removedAt
	}
	return false
}

func (s *SnapshotSet) Iterator() Iterator {
	s.t++
	return NewIterator(s.arr, s.t-1)
}

type Iterator struct {
	data         []*SnapshotElement
	index        int
	totalLen     int
	snapshotTime int
}

func NewIterator(data []*SnapshotElement, t int) Iterator {

	totalLen := 0
	for _, v := range data {
		if v.createdAt <= t && (t < v.removedAt || v.removedAt == 0) {
			totalLen++
		}
	}

	return Iterator{
		data:         data,
		index:        0,
		snapshotTime: t,
		totalLen:     totalLen,
	}
}

func (it *Iterator) HasNext() bool {
	return it.index < it.totalLen
}

func (it *Iterator) Next() int {
	if !it.HasNext() {
		return -1
	}

	for _, v := range it.data[it.index:] {
		if v.createdAt <= it.snapshotTime && (it.snapshotTime < v.removedAt || v.removedAt == 0) {
			it.index++
			return v.val
		}
	}
	return -1
}

func main() {
	// Test cases
	set := NewSnapshotSet()

	// Test 1: Add elements and check snapshot
	set.Add(1)
	set.Add(2)
	set.Add(3)
	itr1 := set.Iterator()

	// Test 2: Remove elements and create a new iterator
	set.Remove(1)
	set.Remove(3)
	itr2 := set.Iterator()

	// Test 3: Add and remove more elements
	set.Add(4)
	set.Remove(2)
	itr3 := set.Iterator()

	// Test Iteration of Iterator 1 (Should return 1, 2, 3)
	fmt.Println("Iterator 1:")
	for itr1.HasNext() {
		fmt.Println(itr1.Next())
	}

	// Test Iteration of Iterator 2 (Should return 2)
	fmt.Println("Iterator 2:")
	for itr2.HasNext() {
		fmt.Println(itr2.Next())
	}

	// Test Iteration of Iterator 3 (Should return 4)
	fmt.Println("Iterator 3:")
	for itr3.HasNext() {
		fmt.Println(itr3.Next())
	}

	// Test 4: Contains method
	fmt.Println("Contains(1):", set.Contains(1)) // Should return false (removed)
	fmt.Println("Contains(2):", set.Contains(2)) // Should return false (removed)
	fmt.Println("Contains(3):", set.Contains(3)) // Should return false (removed)
	fmt.Println("Contains(4):", set.Contains(4)) // Should return true (exists)
}
