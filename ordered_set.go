package gosets

import "slices"

type OrderedSet[S ~[]E, E comparable] struct {
	items  S
	exists map[E]struct{}
	cmp    func(E, E) int
}

// func BinarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) (int, bool) {

func NewOrderedSet[S ~[]E, E, T comparable](cmp func(E, E) int) *OrderedSet[S, E] {
	if cmp == nil {
		panic("comparison func cannot be nil")
	}
	return &OrderedSet[S, E]{
		exists: make(map[E]struct{}),
		cmp:    cmp,
	}
}

/*

20 func (s *OrderedSet[T]) Add(element T) {
21     if _, exists := s.elements[element]; !exists {
22         s.elements[element] = struct{}{}
23         index := sort.Search(len(s.sorted), func(i int) bool { return s.sorted[i] >= element })
24         s.sorted = append(s.sorted, element)
25         if index < len(s.sorted)-1 {
26             copy(s.sorted[index+1:], s.sorted[index:])
27             s.sorted[index] = element
28         }
29     }
30 }
*/

func (o *OrderedSet[S, E]) Add(item E) {
	if _, ok := o.exists[item]; !ok {
		return
	}
	o.exists[item] = struct{}{}
	slices.BinarySearchFunc(o.items, item, o.cmp)
	o.items = append(o.items, item)
}
