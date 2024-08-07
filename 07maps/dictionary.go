package maps

const NotFoundError = DictionaryErr("couldn't find the word you were looking for")
const ErrWordExists = DictionaryErr("word already exists")
const ErrWordDoesNotExists = DictionaryErr("can't update because the word doesn't exist")

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", NotFoundError
	}
	return def, nil
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case NotFoundError:
		d[word] = def
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, newDef string) error {
	_, err := d.Search(word)
	switch err {
	case NotFoundError:
		return ErrWordDoesNotExists
	case nil:
		d[word] = newDef
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
