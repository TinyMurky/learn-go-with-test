package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{
		"test": "this is test",
	}

	// 找到 key = "test" 要回傳 value
	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is test"
		assertStrings(t, got, want)
	})

	// 找不到 key = "unknown" 則回傳 ErrNotFound
	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		want := ErrNotFound
		assertError(t, err, want)
	})

}

func TestAdd(t *testing.T) {
	// 如果 key = "test" 不存在則加入 Dictionary,
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}

		word := "test"
		definition := "this is test"

		err := dict.Add(word, definition)
		assertNoError(t, err)

		assertDefinition(t, dict, word, definition)
	})

	// 如果 key = "test" 存在回傳 ErrWordExists
	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "this is test"
		dict := Dictionary{word: definition}

		err := dict.Add(word, "new test")
		assertError(t, err, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {
	// 如果 key = "test" 存在則 update 其中的value,
	// 不存在則回傳 ErrWordDoesNotExist
	t.Run("existed word", func(t *testing.T) {
		word := "test"
		definition := "this is test"
		dict := Dictionary{word: definition}

		newDefinition := "new test"
		err := dict.Update(word, newDefinition)
		assertNoError(t, err)

		assertDefinition(t, dict, word, newDefinition)
	})

	// 如果 key = "test" 不存在則回傳 ErrWordDoesNotExist
	t.Run("new word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{}

		newDefinition := "new test"
		err := dict.Update(word, newDefinition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	// 如果 key = "test" 的話就刪除,
	t.Run("existed word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{word: "this is test"}

		err := dict.Delete(word)
		assertNoError(t, err)

		_, err = dict.Search(word)
		assertError(t, err, ErrNotFound)
	})

	// key = "test" 不存在則回傳 ErrWordDoesNotExist
	t.Run("non-existed word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{}

		err := dict.Delete(word)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

// assertStrings 檢查 是否為期望的string
func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// assertError 檢查 error 是期望的 constant error
func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q, want %q", got, want)
	}
}

// assertNoError 檢查 error 是 nil
func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("got error %q, want nil", got.Error())
	}
}

// assertDefinition 從dictionary 中找到一個value, 並檢查是否與期望值相符
// 1. 放入dictionary,
//
// 2. dictionary.search(word) 找尋 "存在" key = "word"
//
// 3. 比較 dictionary.search(word) 與 definition 是否相符
func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}
