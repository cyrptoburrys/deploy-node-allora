package testcommon

import "math/rand"

type RandomKeyMap[K comparable] struct {
	rand *rand.Rand
	m    map[K]int
	s    []K
}

// RandomKeyMap is a map that is O(1) for insertion, deletion, and random key selection
func NewRandomKeyMap[K comparable](r *rand.Rand) *RandomKeyMap[K] {
	return &RandomKeyMap[K]{
		rand: r,
		m:    make(map[K]int),
		s:    []K{},
	}
}

// Get returns an element from the map
func (rkm *RandomKeyMap[K]) Get(k K) (*K, bool) {
	index, ok := rkm.m[k]
	if !ok {
		return nil, false
	}
	return &rkm.s[index], true
}

// Insert element into the map
func (rkm *RandomKeyMap[K]) Insert(k K) {
	if _, ok := rkm.m[k]; ok {
		return
	}
	rkm.m[k] = len(rkm.s)
	rkm.s = append(rkm.s, k)
}

// Remove element from the map by swapping in the last element in its place.
func (rkm *RandomKeyMap[K]) Delete(k K) {
	indexOfK, ok := rkm.m[k]
	if !ok {
		return
	}
	lastElement := rkm.s[len(rkm.s)-1]
	// set the slice position of the deleted element to the last element
	rkm.s[indexOfK] = lastElement
	// chop off the last element of the slice
	rkm.s = rkm.s[:len(rkm.s)-1]
	// update the index of the last element in the map to its new slice position
	rkm.m[lastElement] = indexOfK
	// delete the element from the map
	delete(rkm.m, k)
}

// Get a random key from the map
func (rkm *RandomKeyMap[K]) RandomKey() K {
	if len(rkm.s) == 0 {
		panic("RandomKey called on empty M")
	}
	return rkm.s[rkm.rand.Intn(len(rkm.s))]
}

// length of the map & slice
func (rkm *RandomKeyMap[K]) Len() int {
	return len(rkm.s)
}
