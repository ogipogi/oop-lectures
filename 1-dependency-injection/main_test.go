package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
