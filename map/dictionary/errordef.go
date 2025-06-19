package dictionary

const (
	// ErrNotFound 表示 Search 時找不到key
	ErrNotFound = DictionaryErr("could not find the word you were looking for")

	// ErrWordExists 表示 Add 時 已存在 key
	ErrWordExists = DictionaryErr("cannot add word because it already exists")

	// ErrWordDoesNotExist 表示 Update, Delete 的時候 key 不存在
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

// DictionaryErr 實作 `error` interface,
//
// https://dave.cheney.net/2016/04/07/constant-errors
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
