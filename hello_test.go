package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := HelloWorld("Karan")
	want := "Hello, Karan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
