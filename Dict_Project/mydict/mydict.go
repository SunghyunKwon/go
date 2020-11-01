package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWOrdExists = errors.New("That word already exists")
	errCantUpdate = errors.New("Can't update the word")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}

	return "", errNotFound
}

// Add for a word
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
		break
	case nil:
		return errWOrdExists
	default:
		break
	}

	return nil
}

// Update for a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
		break
	case errNotFound:
		return errCantUpdate
	}

	return nil
}

// Delete the word in dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
