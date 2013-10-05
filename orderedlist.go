// Package orderedlist is a basic wrapper around container/list.
package orderedlist

import (
	"container/list"
	"fmt"
)

type Comparable interface {
	Compare(c Comparable) int
}

// OrderedList is an ordered linked list.
type OrderedList struct {
	linkedlist *list.List
}

// New returns an initialized OrderedList.
func New() *OrderedList {
	return &OrderedList{
		linkedlist: list.New(),
	}
}

// Insert inserts a key string into the ordered list.
func (l *OrderedList) Insert(c Comparable) {
	// Empty list or greatest key
	if l.linkedlist.Len() == 0 || l.linkedlist.Back().Value.(Comparable).Compare(c) < 0 {
		l.linkedlist.PushBack(c)
		return
	}

	// Insert in O(n) time
	for e := l.linkedlist.Front(); e != nil; e = e.Next() {
		if e.Value.(Comparable).Compare(c) > 0 {
			l.linkedlist.InsertBefore(c, e)
			return
		}
	}

	fmt.Println(c, 0)
}

// Remove removes a key from the ordered list.
func (l *OrderedList) Remove(c Comparable) {
	for e := l.linkedlist.Front(); e != nil; e = e.Next() {
		if e.Value.(Comparable).Compare(c) == 0 {
			l.linkedlist.Remove(e)
			return
		}
	}
}

// firstGreaterThanOrEqual returns the first Element
// greater-than or equal-to the given key.
func (l *OrderedList) firstGreaterThanOrEqual(c Comparable) *list.Element {
	elem := l.linkedlist.Front()
	for e := elem; e != nil; e = e.Next() {
		if e.Value.(Comparable).Compare(c) >= 0 {
			return e
		}
	}

	return elem
}

func (l *OrderedList) GetRange(start Comparable, end Comparable) (keys []Comparable) {
	keys = make([]Comparable, 0)
	startElem := l.firstGreaterThanOrEqual(start)
	for e := startElem; e != nil; e = e.Next() {
		if e.Value.(Comparable).Compare(end) < 0 {
			keys = append(keys, e.Value.(Comparable))
		}
	}
	return
}

// Print prints the values stored in the list.
func (l *OrderedList) Print() {
	for e := l.linkedlist.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
