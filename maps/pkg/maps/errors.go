package maps

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

// Returned when attempting to look up a key that doesn't exist
const KeyError = DictionaryError("Can't look up a key that doesn't already exist")

// Returned when attempting to add a key-value pair for which the key already exists
const KeyExistsError = DictionaryError("Can't add a key that already exists")

// Returned when attempting to update a key that does not exist
const KeyDoesntExistDeleteError = DictionaryError("Can't delete a key that doesn't exist")

// Returned when attempting to update a key that does not exist
const KeyDoesntExistUpdateError = DictionaryError("Can't update a key that doesn't exist")
