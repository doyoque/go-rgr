package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrentMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Doyoque", "")
		want := "Hello Doyoque!"

		assertCorrentMessage(t, got, want)
	})

	t.Run("Saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Doyoque", "Spanish")
		want := "Hola Doyoque!"

		assertCorrentMessage(t, got, want)
	})

	t.Run("Say 'hello world!' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world!"
		assertCorrentMessage(t, got, want)
	})
}
