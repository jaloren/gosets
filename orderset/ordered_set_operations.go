package orderset

func Union[S ~[]E, E comparable](first, second *OrderedSet[S, E]) *OrderedSet[S, E] {
	if (first.Empty() && second.Empty()) || second.Empty() {
		return first.Clone()
	} else if first.Empty() {
		return second.Clone()
	}
	setUnion := first.Clone()
	setUnion.Add(second.items...)
	return setUnion
}

func Intersect[S ~[]E, E comparable](first, second *OrderedSet[S, E]) *OrderedSet[S, E] {
	setIntersect := NewWithComparator[S, E](first.cmp)
	if first.Empty() || second.Empty() {
		return setIntersect
	}
	allItems := append(first.items, second.items...)
	var intersectItems []E
	for _, item := range allItems {
		if first.Contains(item) && second.Contains(item) {
			intersectItems = append(intersectItems, item)
		}
	}
	setIntersect.Add(intersectItems...)
	return setIntersect
}
