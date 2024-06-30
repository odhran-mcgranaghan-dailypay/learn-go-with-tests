package maps

type Dictionary map[string]string

const (
	ErrNotFound               = DictionaryErr("could not find the word you were looking for")
	ErrWordExists             = DictionaryErr("this word already exists in the dictionary")
	ErrUpdateWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
	ErrDeleteWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (de DictionaryErr) Error() string {
	return string(de)
}

// If a map isnt a reference variable, what is it?
// https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it

// Never initialise a nil map e.g. var m map[string]string
// You should either initialize an empty map or use make to create a new map
// var m = make(map[string]string) or m := map[string]string{}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Map will not throw an error if the value already exists, they will just overwrite the value
// In order to satisfy the Add function, we only want to add a word if it doesn't already exist
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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrUpdateWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err != nil {
		return ErrDeleteWordDoesNotExist
	}
	delete(d, word)
	return nil
}
