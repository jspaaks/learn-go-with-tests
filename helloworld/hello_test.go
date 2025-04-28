package helloworld

import "testing"

func assertCorrectMessage(t *testing.T, got string, want string) {
	t.Helper() // affects the reported line number in case of failure
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSayHello(t *testing.T) {

	t.Run("saying hello to people in dutch", func(t *testing.T) {
		got := SayHello("Chris", "Dutch")
		want := "Hallo, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in english", func(t *testing.T) {
		got := SayHello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in french", func(t *testing.T) {
		got := SayHello("Chris", "French")
		want := "Bonjour, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in spanish", func(t *testing.T) {
		got := SayHello("Chris", "Spanish")
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in unknown language", func(t *testing.T) {
		got := SayHello("Chris", "wdfg245245")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, world' when an empty string is suplied", func(t *testing.T) {
		got := SayHello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

}
