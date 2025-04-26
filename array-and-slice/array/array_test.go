package array

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)

		expect := 15

		assertEqual(t, got, expect, numbers)
	})
}

func BenchmarkSum(b *testing.B) {

	numbers := [5]int{1, 2, 3, 4, 5}

	for b.Loop() {
		// BenchmarkSum-16    	256567549	         5.163 ns/op	       0 B/op	       0 allocs/op
		Sum(numbers)
	}
}

func assertEqual[T comparable](t testing.TB, got, want T, origin any) {
	t.Helper()

	if got != want {
		t.Errorf("Expect \"%v\", got \"%v\", input: %v", want, got, origin)
	}
}
