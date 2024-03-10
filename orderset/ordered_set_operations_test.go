package orderset

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnion(t *testing.T) {
	claim := require.New(t)
	setOne := New[[]int, int]()
	setOne.Add(1, 2, 3)
	setTwo := New[[]int, int]()
	setTwo.Add(1, 2, 3, 4, 5, 6, 7)
	setUnion := Union(setOne, setTwo)
	claim.True(setUnion.ContainsAll(setOne))
	claim.True(setUnion.ContainsAll(setTwo))
	claim.True(setUnion.Equal(setTwo))
}
