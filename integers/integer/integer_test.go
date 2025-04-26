package integer

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("Add two integer", func(t *testing.T) {
		got := Add(2, 2)
		want := 4

		assetEqual(t, got, want)
	})
}

func ExampleAdd() {
	// 最下面那個Output go是真的會檢查的
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6

}

func assetEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got \"%v\", expect \"%v\"", got, want)
	}
}
