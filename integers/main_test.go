package main

import "testing"

func TestAdd(t *testing.T) {
	t.Run("Add two integer", func(t *testing.T) {
		got := Add(2, 2)
		want := 4

		assetEqual(t, got, want)
	})
}

func assetEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got \"%v\", expect \"%v\"", got, want)
	}
}
