package gosets_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/jlorenzini/gosets"
)

func TestOrderedSetAll(t *testing.T) {
	claim := require.New(t)
	set := gosets.NewOrderedSet[int, int](gosets.DefaultCmp[int]())
	input := []int{5, 100, 1, 90, 35, 45}
	set.Add(input...)

	allItems := set.All()
	claim.Len(allItems, len(input))

	for _, item := range input {
		claim.Contains(allItems, item)
	}
}
