package gosets

import (
	"cmp"
	"fmt"

	"iter"
	"slices"
)

type OrderedSet[S ~[]E, E comparable] struct {
	items  S
	exists map[E]struct{}
	cmp    func(E, E) int
}

func (o *OrderedSet[S, E]) String() string {
	return fmt.Sprintf("%+v", o.items)
}

func (o *OrderedSet[S, E]) Len() int {
	return len(o.items)
}

func New[S ~[]E, E cmp.Ordered]() *OrderedSet[S, E] {
	cmpFunc := func(e E, t E) int {
		return cmp.Compare(e, t)
	}
	return NewWithComparator[S, E](cmpFunc)
}

func NewWithComparator[S ~[]E, E comparable](cmp func(E, E) int) *OrderedSet[S, E] {
	if cmp == nil {
		panic("comparison func cannot be nil")
	}
	return &OrderedSet[S, E]{
		exists: make(map[E]struct{}),
		cmp:    cmp,
	}
}

func (o *OrderedSet[S, E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, item := range o.items {
			if !yield(item) {
				return
			}
		}
	}
}

func (o *OrderedSet[S, E]) Add(items ...E) {
	for _, item := range items {
		if _, ok := o.exists[item]; ok {
			return
		}
		o.exists[item] = struct{}{}
		index, _ := slices.BinarySearchFunc(o.items, item, o.cmp)
		o.items = append(o.items, item)
		if index < len(o.items)-1 {
			copy(o.items[index+1:], o.items[index:])
			o.items[index] = item
		}
	}
}

func (o *OrderedSet[S, E]) Contains(item E) bool {
	_, ok := o.exists[item]
	return ok
}

func (o *OrderedSet[S, E]) Remove(items ...E) {
	for _, item := range items {
		if _, ok := o.exists[item]; !ok {
			return
		}
		delete(o.exists, item)
		index, _ := slices.BinarySearchFunc(o.items, item, o.cmp)
		if index < len(o.items) && o.items[index] == item {
			o.items = append(o.items[:index], o.items[index+1:]...)
		}
	}
}
