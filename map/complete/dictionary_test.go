package dictionary

import (
	"testing"
)

func Test_Search1(t *testing.T) {
	dict := map[string]string{
		"test": "this is test",
	}

	got := Search1(dict, "test")
	want := "this is test"

	assertStrings(t, got, want)
}

func Test_DicSearch(t *testing.T) {
	dict := Dictionary{
		"test": "this is test",
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		want := ErrNotFound
		assertError(t, err, want)
	})

}

func Test_DictAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}

		word := "test"
		definition := "this is test"

		err := dict.Add(word, definition)
		assertNoError(t, err)

		assertDefinition(t, dict, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "this is test"
		dict := Dictionary{word: definition}

		err := dict.Add(word, "new test")
		assertError(t, err, ErrWordExists)
	})
}

func Test_DictUpdate(t *testing.T) {
	t.Run("existed word", func(t *testing.T) {
		word := "test"
		definition := "this is test"
		dict := Dictionary{word: definition}

		newDefinition := "new test"
		err := dict.Update(word, newDefinition)
		assertNoError(t, err)

		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{}

		newDefinition := "new test"
		err := dict.Update(word, newDefinition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func Test_DictDelete(t *testing.T) {
	t.Run("existed word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{word: "this is test"}

		err := dict.Delete(word)
		assertNoError(t, err)

		_, err = dict.Search(word)
		assertError(t, err, ErrNotFound)
	})

	t.Run("non-existed word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{}

		err := dict.Delete(word)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("got error %q, want nil", got.Error())
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}
