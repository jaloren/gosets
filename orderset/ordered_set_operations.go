package orderset

func Union[S ~[]E, E comparable](first, second *OrderedSet[S, E]) *OrderedSet[S, E] {
	setUnion := NewWithComparator[S, E](first.cmp)
	if first.Empty() && second.Empty() {
		return setUnion
	} else if second.Empty() {

	}

	return nil
}
