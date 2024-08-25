package main

/*
[] [] [] [] [] []
[] [] [] [] [] []
[] []    [] [] []
[]       [] [] []


[] = pair: snapshot id, value

*/

type Pair struct {
	ID    int
	Value int
}

type SnapshotArray struct {
	data       [][]Pair
	snapshotID int
}

func NewSnapshotArray(length int) SnapshotArray {
	data := make([][]Pair, length)
	for i := range data {
		data[i] = append(data[i], Pair{ID: 0, Value: 0})
	}

	return SnapshotArray{
		data:       data,
		snapshotID: 0,
	}
}

func (r *SnapshotArray) Set(index int, value int) {
	lastIndex := len(r.data[index]) - 1
	if r.data[index][lastIndex].ID == r.snapshotID {
		r.data[index][lastIndex].Value = value
	} else {
		r.data[index] = append(r.data[index], Pair{ID: r.snapshotID, Value: value})
	}
}

func (r *SnapshotArray) Snap() int {
	r.snapshotID++
	return r.snapshotID - 1
}

func (r *SnapshotArray) Get(index int, snapId int) int {
	data := r.data[index]

	low, high := 0, len(data)
	for low <= high {
		mid := low + (high-low)/2
		if data[mid].ID == snapId {
			return data[mid].Value
		}

		if data[mid].ID < snapId {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if high >= 0 {
		return data[high].Value
	}

	return 0
}
