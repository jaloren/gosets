package gosets

import (
	"cmp"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddRemove(t *testing.T) {
	claim := require.New(t)
	set := New[[]int, int]()
	input := []int{5, 100, 1, 90, 35, 45}
	set.Add(input...)
	claim.True(len(input) == set.Len())
	for _, item := range input {
		claim.Truef(set.Contains(item), "set does not contain %d when it should have this element", item)
	}
	isSorted := slices.IsSortedFunc(set.items, set.cmp)
	claim.Truef(isSorted, "set is not ordered: %v", set.items)
	set.Remove(input...)
	claim.True(set.Len() == 0)
}

func TestIteration(t *testing.T) {
	claim := require.New(t)
	set := New[[]int, int]()
	input := []int{5, 100, 1, 90, 35, 45}
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
	actual := make([]Person, 0)
	for v := range set.All() {
		actual = append(actual, v)
	}
	slices.SortFunc(persons, cmpFunc)

	claim.Equal(persons, actual)
}
