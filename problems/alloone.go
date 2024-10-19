package main

import "container/list"

type Bucket struct {
	count int
	keys  map[string]bool
}

type AllOne struct {
	list      *list.List
	bucketMap map[string]*list.Element
}

func Constructor() AllOne {
	return AllOne{
		list:      list.New(),
		bucketMap: make(map[string]*list.Element),
	}
}

// add key to element
func (r *AllOne) addKey(key string, elem *list.Element) {
	bucket := elem.Value.(*Bucket)
	bucket.keys[key] = true
	r.bucketMap[key] = elem
}

// remove key from element
func (r *AllOne) removeKey(key string, elem *list.Element) {
	bucket := elem.Value.(*Bucket)
	delete(bucket.keys, key)
	if len(bucket.keys) == 0 {
		r.list.Remove(elem)
	}
}

func (r *AllOne) Inc(key string) {
	// if bucket found, move key up
	// else check bucket with count == 1 exist or not, if not create bucket, else add key to front bucket
	if elem, found := r.bucketMap[key]; found {
		r.moveUp(key, elem)
	} else {
		if r.list.Len() == 0 || r.list.Front().Value.(*Bucket).count > 1 {
			bucket := &Bucket{
				count: 1,
				keys:  make(map[string]bool),
			}
			r.list.PushFront(bucket)
		}
		elem := r.list.Front()
		r.addKey(key, elem)
	}
}

func (r *AllOne) moveUp(key string, elem *list.Element) {
	currBucket := elem.Value.(*Bucket)
	newCount := currBucket.count + 1

	nextElem := elem.Next()
	// if next bucket exist, add key to next bucket, remove key from current buket
	// else create new bucket, add key to new buket, remove key from current bucket
	if nextElem != nil && nextElem.Value.(*Bucket).count == newCount {
		r.addKey(key, nextElem)
		r.removeKey(key, elem)
	} else {
		newBucket := &Bucket{
			count: newCount,
			keys:  make(map[string]bool),
		}
		newElem := r.list.InsertAfter(newBucket, elem)
		r.addKey(key, newElem)
		r.removeKey(key, elem)
	}
}

// if key does not exist, return
// if key count is 1, remove key from bucket, and remove key from map
// else move key down
func (r *AllOne) Dec(key string) {
	if _, found := r.bucketMap[key]; !found {
		return
	}

	currElem := r.bucketMap[key]
	currBucket := currElem.Value.(*Bucket)

	if currBucket.count == 1 {
		r.removeKey(key, currElem)
		delete(r.bucketMap, key)
	} else {
		r.moveDown(key, currElem)
	}
}

func (r *AllOne) moveDown(key string, currElem *list.Element) {
	currBucket := currElem.Value.(*Bucket)
	newCount := currBucket.count - 1

	prevElem := currElem.Prev()
	// if prev bucket exist, move key
	// else create bucket, move key
	if prevElem != nil && prevElem.Value.(*Bucket).count == newCount {
		r.addKey(key, prevElem)
		r.removeKey(key, currElem)
	} else {
		bucket := &Bucket{
			count: newCount,
			keys:  make(map[string]bool),
		}

		newElem := r.list.InsertBefore(bucket, currElem)
		r.addKey(key, newElem)
		r.removeKey(key, currElem)
	}

}

func (r *AllOne) GetMaxKey() string {
	if r.list.Len() == 0 {
		return ""
	}

	keys := r.list.Back().Value.(*Bucket).keys
	for key := range keys {
		return key
	}

	return ""
}

func (r *AllOne) GetMinKey() string {
	if r.list.Len() == 0 {
		return ""
	}

	keys := r.list.Front().Value.(*Bucket).keys
	for key := range keys {
		return key
	}

	return ""
}
