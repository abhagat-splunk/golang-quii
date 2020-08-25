package helloworld

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q wanted %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {

		got := Hello("Ankit", "English")
		want := "Hello, Ankit!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World' if empty string is supplied", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, world!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})
}
