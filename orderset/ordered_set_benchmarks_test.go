package orderset

import (
	"math/rand/v2"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	items := randomInts()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set := New[[]int, int]()
		set.Add(items...)
	}

}

func randomInts() []int {
	var randomInts []int
	for range 1000 {
		randomInts = append(randomInts, rand.Int())
	}
	return randomInts
}
