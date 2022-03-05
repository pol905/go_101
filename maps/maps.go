package maps

const (
	ErrNotFound      = DictionaryErr("could not find the word you were looking for")
	ErrAlreadyExists = DictionaryErr("the provided key already exists")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

type Dictionary map[string]string

func (d Dictionary) Search(searchString string) (string, error) {
	definition, ok := d[searchString]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrNotFound
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
