package main

import "testing"

func TestHelloWorld(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, want, got string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("say Hello in English", func(t *testing.T) {
		got := HelloWorld("Karan", "")
		want := "Hello, Karan"

		assertCorrectMessage(t, want, got)
	})

	t.Run("say 'Hello, World' if an empty string is passed", func(t *testing.T) {
		got := HelloWorld("", "")
		want := "Hello, World"

		assertCorrectMessage(t, want, got)
	})

	t.Run("say Hello in Spanish", func(t *testing.T) {
		got := HelloWorld("Jose", "Spanish")
		want := "Hola, Jose"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello in French", func(t *testing.T) {
		got := HelloWorld("Francois", "French")
		want := "Bonjour, Francois"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello in Italian", func(t *testing.T) {
		got := HelloWorld("Francesco", "Italian")
		want := "Buongiorno, Francesco"

		assertCorrectMessage(t, got, want)
	})
}
