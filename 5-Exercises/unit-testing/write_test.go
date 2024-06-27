package main

import (
	"bytes"
	"testing"
)

func TestWriteMessages(t *testing.T) {
	buffer := bytes.Buffer{}
	messages := []string{"Hello, Go", "Hello, World"}
	WriteMessages(&buffer, messages)

	got := buffer.String()
	want := "Hello, Go\nHello, World\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
