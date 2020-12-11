package maps

// // ErrNotFound is an error saying that the word was not found
// var ErrNotFound = errors.New("can't find unknown")

// // ErrWordExists is an error to pop when the same word is added again
// var ErrWordExists = errors.New("Word already exists")

const (
	// ErrNotFound is an error saying that the word was not found
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists is an error to pop when the same word is added again
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	// ErrWordDoesNotExist is an error that says the word doesn't exist
	ErrWordDoesNotExist = DictionaryErr("cannot update word cause it doesn't exist")
)

// DictionaryErr is an error related to the dictionary
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary is a mapping from string to string
type Dictionary map[string]string

// Search is used to return the get from the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add is used to add a key value mapping to a dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update is used to update the key value mapping
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

// Delete is used to delete the key value mapping
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
