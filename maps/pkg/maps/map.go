package maps

// A dictionary key.
type Key string

// A dictionary value.
type Value string

// Definition of dictionary.
type Dictionary map[Key]Value

// Adds a key-value pair to the receiver dictionary unless the key already exists.
func (d Dictionary) Add(k Key, v Value) error {
	_, err := d.Lookup(k)
	switch err {
	case KeyDoesntExistLookupError:
		d[k] = v
	case nil:
		return KeyAlreadyExistsError
	default:
		return err
	}
	return nil
}

// Deletes a key-value pair from the receiver dictionary unless the provided key doesn't exist.
func (d Dictionary) Delete(k Key) error {
	_, err := d.Lookup(k)
	switch err {
	case KeyDoesntExistLookupError:
		return KeyDoesntExistDeleteError
	case nil:
		delete(d, k)
		return nil
	default:
		return err
	}
	return nil
}

// Retrieves the value from the receiver dictionary given a key, unless the key doesn't exist.
func (d Dictionary) Lookup(k Key) (Value, error) {
	v, found := d[k]
	if !found {
		return "", KeyDoesntExistLookupError
	}
	return v, nil
}

// Updates the value from the receiver dictionary given a key, unless the key doesn't exist.
func (d Dictionary) Update(k Key, v Value) error {
	_, err := d.Lookup(k)
	switch err {
	case KeyDoesntExistLookupError:
		return KeyDoesntExistUpdateError
	case nil:
		d[k] = v
		return nil
	default:
		return err
	}
	return nil
}
