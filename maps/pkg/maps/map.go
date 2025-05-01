package maps

type Key string
type Value string
type Dictionary map[Key]Value

func (d Dictionary) Add(k Key, v Value) error {

	_, err := d.Lookup(k)

	switch err {
	case KeyError:
		d[k] = v
	case nil:
		return KeyExistsError
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(k Key) error {

	_, err := d.Lookup(k)

	switch err {
	case KeyError:
		return KeyDoesntExistDeleteError
	case nil:
		delete(d, k)
		return nil
	default:
		return err
	}
	return nil
}

func (d Dictionary) Lookup(k Key) (Value, error) {
	v, found := d[k]
	if !found {
		return "", KeyError
	}
	return v, nil
}

func (d Dictionary) Update(k Key, v Value) error {

	_, err := d.Lookup(k)

	switch err {
	case KeyError:
		return KeyDoesntExistUpdateError
	case nil:
		d[k] = v
		return nil
	default:
		return err
	}
	return nil
}
