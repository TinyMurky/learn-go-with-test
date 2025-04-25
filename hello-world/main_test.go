package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Murky")
		want := "Hello, Murky"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")

		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got string, want string) {
	// TB 可以滿足 *testing.T and *testing.B
	t.Helper() // 這個是interface, 專門告訴tester這個function只是helper不是test
	// 這樣錯誤的時候顯示的錯誤行數就會是test case的，不會跑進來

	if got != want {
		t.Errorf("got %q, expect %q", got, want)
	}
}
