// Package dictionary is about map
package dictionary

// 這邊可以介紹 Uber style
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

// DictionaryErr https://dave.cheney.net/2016/04/07/constant-errors
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// 備註： 想要講的重點有

// 哪些 type 可以當作key (comparable才可以, array裡面全部要comparble)

// map 回 struct or true

// Search1 直接用 map return
func Search1(dict map[string]string, word string) string {
	return dict[word]
}

// Dictionary 就是map
type Dictionary map[string]string

// Search 是 dictionary 用的 Search
func (d Dictionary) Search(word string) (string, error) {
	// 這是可以直接修改到 map 本身的值
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add add to dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

// Update to map
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
}

// Delete remove element
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		// If map is nil or there is no such element, delete is a no-op.
		delete(d, word)
		return nil
	default:
		return err
	}
}
