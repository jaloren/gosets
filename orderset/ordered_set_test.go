package gosets

import (
	"cmp"
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
	set.Remove(input...)
	claim.True(set.Len() == 0)
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
