// Package orderedlist is a basic wrapper around container/list.
package orderedlist

import (
	"container/list"
	"fmt"
)

// OrderedList is an ordered linked list.
type OrderedList struct {
	linkedlist *list.List
	compare    func(l, r interface{}) int
}

type RangeIterator struct {
	e          *list.Element
	compare    func(l, r interface{}) int
	start, end interface{}
}

func (r *RangeIterator) withinRange(start, end interface{}) bool {
	if r.e == nil {
		return false
	}

	return r.compare(r.e.Value, start) > 0 &&
		r.compare(r.e.Value, end) < 0
}

// Next returns a pointer to the next RangeIterator. Nil is returned
// if the next value is out of the original range requested.
func (r *RangeIterator) Next() *RangeIterator {
	next := &RangeIterator{
		e:       r.e.Next(),
		start:   r.start,
		end:     r.end,
		compare: r.compare,
	}

	if next.withinRange(r.start, r.end) {
		return next
	}

	return nil
}

// Prev returns a pointer to the next RangeIterator. Nil is returned
// if the previous value is out of the original range requested.
func (r *RangeIterator) Prev() *RangeIterator {
	prev := &RangeIterator{
		e:       r.e.Prev(),
		start:   r.start,
		end:     r.end,
		compare: r.compare,
	}

	if prev.withinRange(r.start, r.end) {
		return prev
	}

	return nil
}

// Value returns the value at the current element.
func (r *RangeIterator) Value() interface{} {
	return r.e.Value
}

// New returns an initialized OrderedList.
func New(compare func(a, b interface{}) int) *OrderedList {
	return &OrderedList{
		linkedlist: list.New(),
		compare:    compare,
	}
}

// Insert inserts a key string into the ordered list.
func (l *OrderedList) Insert(c interface{}) {
	// Empty list or greatest key
	if l.linkedlist.Len() == 0 || l.compare(l.linkedlist.Back().Value, c) < 0 {
		l.linkedlist.PushBack(c)
		return
	}

	// Insert in O(n) time
	for e := l.linkedlist.Front(); e != nil; e = e.Next() {
		if l.compare(e.Value, c) > 0 {
			l.linkedlist.InsertBefore(c, e)
			return
		}
	}
}

// Remove removes a key from the ordered list.
func (l *OrderedList) Remove(c interface{}) {
	for e := l.linkedlist.Front(); e != nil; e = e.Next() {
		if l.compare(e.Value, c) == 0 {
			l.linkedlist.Remove(e)
			return
		}
	}
}

// firstGreaterThanOrEqual returns the first Element
// greater-than or equal-to the given key.
func (l *OrderedList) firstGreaterThanOrEqual(c interface{}) *list.Element {
	elem := l.linkedlist.Front()
	for e := elem; e != nil; e = e.Next() {
		if l.compare(e.Value, c) >= 0 {
			return e
		}
	}

	return elem
}

// GetRange returns a slice of interface{} types in the range [start, end).
func (l *OrderedList) GetRange(start interface{}, end interface{}) (keys []interface{}) {
	keys = make([]interface{}, 0, l.linkedlist.Len())
	startElem := l.firstGreaterThanOrEqual(start)
	for e := startElem; e != nil; e = e.Next() {
		if l.compare(e.Value, end) < 0 {
			keys = append(keys, e.Value)
		}
	}
	return
}

// GetRangeIterator returns a pointer to a RangeIterator.
func (l *OrderedList) GetRangeIterator(start, end interface{}) *RangeIterator {
	elem := l.firstGreaterThanOrEqual(start)
	if elem == nil {
		return nil
	}

	return &RangeIterator{
		e:       elem,
		start:   start,
		end:     end,
		compare: l.compare,
	}
}

// String returns a string representation of the values stored in the list.
func (l *OrderedList) String() string {
	ret := "["
	for e := l.linkedlist.Front(); e != nil; e = e.Next() {
		ret += fmt.Sprintf(" %v", e.Value)
	}
	return ret + " ]"
}
