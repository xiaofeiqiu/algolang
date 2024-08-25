package main

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
	}
}

func (s *SnapshotSet) Remove(v int) {
	if e, exists := s.data[v]; exists {
		e.removedAt = s.t
	}
}

func (s *SnapshotSet) Contains(v int) bool {
	if metadata, exists := s.data[v]; exists {
		return metadata.createdAt > metadata.removedAt
	}
	return false
}

func (s *SnapshotSet) Interator() Iterator {
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
		if v.createdAt <= t && t < v.removedAt {
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

	for _, v := range it.data[it.index+1:] {
		if v.createdAt <= it.snapshotTime && it.snapshotTime < v.removedAt {
			it.index++
			return v.val
		}
	}
	return -1
}
