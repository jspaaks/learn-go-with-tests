package maps

import "testing"

var assertCorrectValueForKey = func(t *testing.T, want Value, got Value) {
	t.Helper()
	if got != want {
		t.Fatal("Unexpected value for key")
	}
}

var assertKeyDoesntExistDeleteError = func(t *testing.T, err error) {
	t.Helper()
	if err != KeyDoesntExistDeleteError {
		t.Fatal("Expected a KeyDoesntExistDeleteError but didn't get one")
	}
}

var assertKeyDoesntExistLookupError = func(t *testing.T, err error) {
	t.Helper()
	if err != KeyDoesntExistLookupError {
		t.Fatal("Expected a KeyDoesntExistLookupError but didn't get one")
	}
}

var assertKeyDoesntExistUpdateError = func(t *testing.T, err error) {
	t.Helper()
	if err != KeyDoesntExistUpdateError {
		t.Fatal("Expected a KeyDoesntExistUpdateError but didn't get one")
	}
}

var assertKeyAlreadyExistsError = func(t *testing.T, err error) {
	t.Helper()
	if err != KeyAlreadyExistsError {
		t.Fatal("Expected a KeyAlreadyExistsError but didn't get one")
	}
}

var assertNoErrors = func(t *testing.T, err error) {
	t.Helper()
	if err == KeyAlreadyExistsError {
		t.Fatal("Didn't expect a KeyAlreadyExistsError but got one")
	}
	if err == KeyDoesntExistDeleteError {
		t.Fatal("Didn't expect a KeyDoesntExistDeleteError but got one")
	}
	if err == KeyDoesntExistLookupError {
		t.Fatal("Didn't expect a KeyDoesntExistLookupError but got one")
	}
	if err == KeyDoesntExistUpdateError {
		t.Fatal("Didn't expect a KeyDoesntExistUpdateError but got one")
	}
	if err != nil {
		t.Fatal("Didn't expect any error but got one")
	}
}

func TestAdd(t *testing.T) {
	t.Run("adding a kv pair to an intially empty dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("testkey", "testvalue")
		got, err := dictionary.Lookup("testkey")
		assertNoErrors(t, err)
		assertCorrectValueForKey(t, "testvalue", got)
	})
	t.Run("attempting to add a kv pair to a dictionary for which the key already exists", func(t *testing.T) {
		dictionary := Dictionary{"testkey": "testvalue"}
		err := dictionary.Add("testkey", "newtestvalue")
		assertKeyAlreadyExistsError(t, err)
		got, err := dictionary.Lookup("testkey")
		assertCorrectValueForKey(t, "testvalue", got)
	})
}

func TestDelete(t *testing.T) {
	t.Run("deleting the value of an existing key in a dictionary", func(t *testing.T) {
		dictionary := Dictionary{"testkey": "testvalue"}
		err := dictionary.Delete("testkey")
		assertNoErrors(t, err)
		got, err := dictionary.Lookup("testkey")
		assertKeyDoesntExistLookupError(t, err)
		assertCorrectValueForKey(t, "", got)
	})
	t.Run("attempting to delete the value of a nonexisting key in a dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("testkey")
		assertKeyDoesntExistDeleteError(t, err)
	})
}

func TestLookup(t *testing.T) {
	t.Run("looking up an existing key in a dictionary", func(t *testing.T) {
		dictionary := Dictionary{"testkey": "testvalue"}
		got, err := dictionary.Lookup("testkey")
		assertNoErrors(t, err)
		assertCorrectValueForKey(t, "testvalue", got)
	})
	t.Run("attempting to look up a nonexisting key in a dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		got, err := dictionary.Lookup("testkey")
		assertKeyDoesntExistLookupError(t, err)
		assertCorrectValueForKey(t, "", got)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("updating the value of an existing key in a dictionary", func(t *testing.T) {
		dictionary := Dictionary{"testkey": "testvalue"}
		err := dictionary.Update("testkey", "newtestvalue")
		assertNoErrors(t, err)
		got, err := dictionary.Lookup("testkey")
		assertCorrectValueForKey(t, "newtestvalue", got)
	})
	t.Run("attempting to update the value of a nonexisting key in a dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("testkey", "testvalue")
		assertKeyDoesntExistUpdateError(t, err)
		got, err := dictionary.Lookup("testkey")
		assertCorrectValueForKey(t, "", got)
	})
}
