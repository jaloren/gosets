package orderset

import (
	"cmp"
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddRemove(t *testing.T) {
	claim := require.New(t)
	set, input := testSet()
	claim.True(len(input) == set.Len())
	for _, item := range input {
		claim.Truef(set.Contains(item), "set does not contain %d when it should have this element", item)
	}
	isSorted := slices.IsSortedFunc(set.items, set.cmp)
	claim.Truef(isSorted, "set is not ordered: %v", set.items)
	// prove i can remove elements even if one of the elements isn't already in the set
	removeElements := append([]int{10000}, input...)
	set.Remove(removeElements...)
	claim.True(set.Len() == 0)
}

func TestAddRemoveWithDuplicates(t *testing.T) {
	claim := require.New(t)
	set, input := testSet()
	input = append(input, 1000, 2000, 3000)
	set.Add(input...)
	for _, item := range input {
		claim.Truef(set.Contains(item), "set does not contain %d when it should have this element", item)
	}
	isSorted := slices.IsSortedFunc(set.items, set.cmp)
	claim.Truef(isSorted, "set is not ordered: %v", set.items)
	slices.Sort(input)
	set.Remove(input...)
	claim.True(set.Len() == 0)

}

func TestClone(t *testing.T) {
	claim := require.New(t)
	setOne, _ := testSet()
	setTwo := setOne.Clone()
	claim.True(setOne.Equal(setTwo))
	claim.True(maps.Equal(setOne.exists, setTwo.exists))
	setOne.Remove(1)
	setTwo.Add(1)
	claim.False(setOne.Equal(setTwo))
	claim.False(maps.Equal(setOne.exists, setTwo.exists))
	setOne.Clear()
	claim.True(setOne.Empty())
	claim.False(setTwo.Empty())
}

func TestContains(t *testing.T) {
	claim := require.New(t)
	setOne, _ := testSet()
	setTwo, _ := testSet()
	setOne.Add(9922, 4590, 560)
	claim.True(setOne.ContainsAll(setTwo))
	claim.False(setTwo.ContainsAll(setOne))
	setOne.Clear()
	claim.False(setOne.ContainsAll(setTwo))
}

func TestClear(t *testing.T) {
	claim := require.New(t)
	set, input := testSet()
	claim.True(len(input) == set.Len())
	set.Clear()
	claim.Len(set.exists, 0)
	claim.Len(set.items, 0)
}

func TestEqual(t *testing.T) {
	claim := require.New(t)
	setOne, _ := testSet()
	setTwo, _ := testSet()
	claim.True(setOne.Equal(setTwo))
	setOne.Add(67)
	claim.False(setOne.Equal(setTwo))
}

func TestIteration(t *testing.T) {
	claim := require.New(t)
	set, input := testSet()
	slices.Sort(input)
	set.Add(input...)
	var actual []int
	for v := range set.All() {
		actual = append(actual, v)
	}
	claim.Equal(actual, input)

}

type Person struct {
	Name string
	Age  int
}

func TestCustomComparator(t *testing.T) {
	claim := require.New(t)
	cmpFunc := func(a Person, b Person) int {
		return cmp.Compare[int](a.Age, b.Age)
	}
	set := NewWithComparator[[]Person, Person](cmpFunc)
	persons := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 20},
	}
	set.Add(persons...)
	slices.SortFunc(persons, cmpFunc)
	claim.Equal(persons, set.items)
}

func testSet() (*OrderedSet[[]int, int], []int) {
	set := New[[]int, int]()
	input := []int{5, 100, 1, 90, 35, 45}
	set.Add(input...)
	return set, input
}
