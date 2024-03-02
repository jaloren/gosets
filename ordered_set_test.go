package gosets

import (
	"cmp"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrderedSetAddRemove(t *testing.T) {
	claim := require.New(t)
	set := NewOrderedSet[[]int, int]()
	input := []int{5, 100, 1, 90, 35, 45}
	set.Add(input...)
	claim.True(len(input) == set.Size())
	for _, item := range input {
		claim.Truef(set.Contains(item), "set does not contain %d when it should have this element", item)
	}
	isSorted := slices.IsSortedFunc(set.items, set.cmp)
	claim.Truef(isSorted, "set is not ordered: %v", set.items)
	set.Remove(input...)
	claim.True(set.Size() == 0)
}
func TestOrderedSetIteration(t *testing.T) {
	claim := require.New(t)
	set := NewOrderedSet[[]int, int]()
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

func TestOrderedSetCustomComparator(t *testing.T) {
	claim := require.New(t)
	cmpFunc := func(a Person, b Person) int {
		return cmp.Compare[int](a.Age, b.Age)
	}
	set := NewOrderedSetWithComparator[[]Person, Person](cmpFunc)

	person1 := Person{Name: "Alice", Age: 25}
	person2 := Person{Name: "Bob", Age: 30}
	person3 := Person{Name: "Charlie", Age: 20}

	set.Add(person1, person2, person3)

	expected := []Person{person3, person1, person2}

	actual := make([]Person, 0)
	for v := range set.All() {
		actual = append(actual, v)
	}

	claim.Equal(expected, actual)
}
