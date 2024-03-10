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
