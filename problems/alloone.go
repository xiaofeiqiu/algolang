package main

import "container/list"

type Bucket struct {
	count  int
	keySet map[string]bool
}

type AllOne struct {
	list      *list.List               // Doubly linked list of Buckets
	bucketMap map[string]*list.Element // Map from key to its Bucket in the list
}

func Constructorx() AllOne {
	return AllOne{
		list:      list.New(),
		bucketMap: make(map[string]*list.Element),
	}
}

func (ao *AllOne) Inc(key string) {
	if elem, exists := ao.bucketMap[key]; exists {
		// Key exists, increment its count
		ao.moveKeyUp(key, elem)
	} else {
		// Key does not exist, insert with count 1
		if ao.list.Len() == 0 || ao.list.Front().Value.(*Bucket).count > 1 {
			// Insert new bucket at front with count 1
			newBucket := &Bucket{
				count:  1,
				keySet: make(map[string]bool),
			}
			ao.list.PushFront(newBucket)
		}
		frontElem := ao.list.Front()
		// Add key to count 1 bucket
		ao.addKey(key, frontElem)
	}
}

func (ao *AllOne) Dec(key string) {
	elem, exists := ao.bucketMap[key]
	if !exists {
		return // Key does not exist, nothing to do
	}
	currBucket := elem.Value.(*Bucket)

	if currBucket.count == 1 {
		// Remove key entirely
		ao.removeKey(key, elem)
		delete(ao.bucketMap, key)

	} else {
		// Key exists and count > 1, decrement its count
		ao.moveKeyDown(key, elem)
	}
}

func (ao *AllOne) moveKeyUp(key string, currElem *list.Element) {
	currBucket := currElem.Value.(*Bucket)
	nextCount := currBucket.count + 1

	var nextElem *list.Element
	if currElem.Next() != nil && currElem.Next().Value.(*Bucket).count == nextCount {
		// Next bucket has the required count
		nextElem = currElem.Next()
	} else {
		// Insert new bucket after current
		newBucket := &Bucket{
			count:  nextCount,
			keySet: make(map[string]bool),
		}
		nextElem = ao.list.InsertAfter(newBucket, currElem)
	}

	// Add key to next bucket
	ao.addKey(key, nextElem)

	// Remove key from current bucket
	ao.removeKey(key, currElem)
}

func (ao *AllOne) moveKeyDown(key string, currElem *list.Element) {
	currBucket := currElem.Value.(*Bucket)
	prevCount := currBucket.count - 1

	var prevElem *list.Element
	if currElem.Prev() != nil && currElem.Prev().Value.(*Bucket).count == prevCount {
		// Previous bucket has the required count
		prevElem = currElem.Prev()
	} else {
		// Insert new bucket before current
		newBucket := &Bucket{
			count:  prevCount,
			keySet: make(map[string]bool),
		}
		prevElem = ao.list.InsertBefore(newBucket, currElem)
	}

	// Add key to previous bucket
	ao.addKey(key, prevElem)

	// Remove key from current bucket
	ao.removeKey(key, currElem)
}

func (ao *AllOne) addKey(key string, elem *list.Element) {
	bucket := elem.Value.(*Bucket)
	bucket.keySet[key] = true
	ao.bucketMap[key] = elem
}

func (ao *AllOne) removeKey(key string, elem *list.Element) {
	bucket := elem.Value.(*Bucket)
	delete(bucket.keySet, key)
	if len(bucket.keySet) == 0 {
		ao.list.Remove(elem)
	}
}

func (ao *AllOne) GetMaxKey() string {
	if ao.list.Len() == 0 {
		return ""
	}
	// Get any key from the tail bucket (max count)
	maxBucket := ao.list.Back().Value.(*Bucket)
	for key := range maxBucket.keySet {
		return key
	}
	return ""
}

func (ao *AllOne) GetMinKey() string {
	if ao.list.Len() == 0 {
		return ""
	}
	// Get any key from the head bucket (min count)
	minBucket := ao.list.Front().Value.(*Bucket)
	for key := range minBucket.keySet {
		return key
	}
	return ""
}
