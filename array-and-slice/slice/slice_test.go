package slice

import (
	"slices"
	"testing"
)

// Use `go test -cover` to generate coverage
// Use `go test ./... -cover` to run all test code
func TestSum(t *testing.T) {
	// this is redundant
	// t.Run("collection of 5 numbers", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5}

	// 	got := Sum(numbers)

	// 	expect := 15

	// 	assertEqual(t, got, expect, numbers)
	// })

	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)

		expect := 6

		assertEqual(t, got, expect, numbers)
	})
}

func BenchmarkSum(b *testing.B) {

	numbers := []int{1, 2, 3, 4, 5}

	for b.Loop() {
		// BenchmarkSum-16    	256023717	         4.062 ns/op	       0 B/op	       0 allocs/o
		Sum(numbers)
	}
}

///////////////////////////////////

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 9})

	expected := []int{6, 13}

	assertSliceIsEqual(t, got, expected)
}

func TestSumTail(t *testing.T) {
	// 這樣也可以，方便用closure的方法把東西丟到assert內
	localAssert := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("Expect \"%v\", got \"%v\"", want, got)
		}
	}
	t.Run("SumTail should sum all number that are not head", func(t *testing.T) {
		got := SumTail([]int{1, 2, 3}, []int{4, 5}, []int{9})

		expected := []int{5, 5, 0}

		localAssert(t, got, expected)
	})

	t.Run("Safely sum empty slice", func(t *testing.T) {
		got := SumTail([]int{}, []int{3, 4, 5})

		expected := []int{0, 9}

		localAssert(t, got, expected)
	})
}

func assertEqual[T comparable](t testing.TB, got, want T, origin any) {
	t.Helper()

	if got != want {
		t.Errorf("Expect \"%v\", got \"%v\", input: %v", want, got, origin)
	}
}

func assertSliceIsEqual[T []E, E comparable](t testing.TB, got, want T) {
	t.Helper()

	// array 是 comparable但slice不是
	// 也可以用 !reflect.DeepEqual(got, want)
	if !slices.Equal(got, want) {
		t.Errorf("Expect \"%v\", got \"%v\"", want, got)
	}
}
