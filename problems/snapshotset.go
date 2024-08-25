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

type Element struct {
	val     int
	isValid map[int]bool
}

type SnapshotSet struct {
	data      map[int]*Element
	arr       []*Element
	versionID int
}

func NewSnapshotSet() SnapshotSet {
	data := make(map[int]*Element)
	return SnapshotSet{
		data:      data,
		versionID: 0,
		arr:       make([]*Element, 0),
	}
}

func (s *SnapshotSet) Add(v int) {
	if _, exist := s.data[v]; !exist {
		s.data[v] = &Element{
			val: v,
			isValid: map[int]bool{
				s.versionID: true,
			},
		}
		s.arr = append(s.arr, s.data[v])
	} else {
		s.data[v].isValid[s.versionID] = true
	}
}

func (s *SnapshotSet) Remove(v int) {
	if e, exists := s.data[v]; exists {
		e.isValid[s.versionID] = false
	}
}

func (s *SnapshotSet) Contains(v int) bool {
	if element, exists := s.data[v]; exists {
		return element.isValid[s.versionID]
	}
	return false
}

func (s *SnapshotSet) Iterator() Iterator {
	for _, v := range s.data {
		if v.isValid[s.versionID] {
			v.isValid[s.versionID+1] = true
		}
	}
	s.versionID++
	return NewIterator(s.arr, s.versionID-1)
}

type Iterator struct {
	data      []*Element
	index     int
	totalLen  int
	versionID int
}

func NewIterator(data []*Element, versionID int) Iterator {

	totalLen := 0
	for _, v := range data {
		if v.isValid[versionID] {
			totalLen++
		}
	}

	return Iterator{
		data:      data,
		index:     0,
		versionID: versionID,
		totalLen:  totalLen,
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
		if v.isValid[it.versionID] {
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
	set.Add(1)
	itr2 := set.Iterator()

	// Test 3: Add and remove more elements
	set.Add(4)
	set.Remove(2)
	itr3 := set.Iterator()

	for itr1.HasNext() {
		fmt.Printf("it1: %d\n", itr1.Next())
	}

	for itr2.HasNext() {
		fmt.Printf("it2: %d\n", itr2.Next())
	}

	for itr3.HasNext() {
		fmt.Printf("it3: %d\n", itr3.Next())
	}

	fmt.Println(set.versionID)
	for _, v := range set.data {
		fmt.Printf("%v\n", v)
	}

	//Test 4: Contains method
	fmt.Println("Contains(1):", set.Contains(1)) // Should return true (removed)
	fmt.Println("Contains(2):", set.Contains(2)) // Should return false (removed)
	fmt.Println("Contains(3):", set.Contains(3)) // Should return false (removed)
	fmt.Println("Contains(4):", set.Contains(4)) // Should return true (exists)
}
