// Package orderedlist is a basic wrapper around container/list.
package orderedlist

import (
	"container/list"
	"fmt"
)

type Comparable interface {
	Compare(c interface{}) int
}

// OrderedList is an ordered linked list.
type OrderedList struct {
	linkedlist *list.List
}

type RangeIterator struct {
	e          *list.Element
	start, end Comparable
}

func (r *RangeIterator) withinRange(start, end Comparable) bool {
	if r.e == nil {
		return false
	}

	return r.e.Value.(Comparable).Compare(start) > 0 &&
		r.e.Value.(Comparable).Compare(end) < 0
}

// Next returns a pointer to the next RangeIterator. Nil is returned
// if the next value is out of the original range requested.
func (r *RangeIterator) Next() *RangeIterator {
	next := &RangeIterator{
		e:     r.e.Next(),
		start: r.start,
		end:   r.end,
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
		e:     r.e.Prev(),
		start: r.start,
		end:   r.end,
	}

	if prev.withinRange(r.start, r.end) {
		return prev
	}

	return nil
}

// Value returns the comparable value at the current element.
func (r *RangeIterator) Value() Comparable {
	if r.e == nil {
		return nil
	}

	return r.e.Value.(Comparable)
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

// GetRange returns a slice of Comparables in the range [start, end).
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

// GetRangeIterator returns a pointer to a RangeIterator.
func (l *OrderedList) GetRangeIterator(start, end Comparable) *RangeIterator {
	elem := l.firstGreaterThanOrEqual(start)
	if elem == nil {
		return nil
	}

	return &RangeIterator{
		e:     elem,
		start: start,
		end:   end,
	}
}

// Print prints the values stored in the list.
func (l *OrderedList) Print() {
	for e := l.linkedlist.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
