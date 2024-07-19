package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := hello("", "English")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := hello("Mbappe", "French")
		want := "Bonjour, Mbappe"
		assertCorrectMessage(t, got, want)
	})
}

/*
This is a helper function.
Calling t.Helper() is required so that the test suite recognises it's a helper
(as regular functions can't be called in this file)
*/

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
