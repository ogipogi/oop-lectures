package main

import (
	"bytes"
	"testing"
)

func TestWriteFormattedMessages(t *testing.T) {
	buffer := bytes.Buffer{}
	messages := []string{"Hello, Go", "Hello, World"}
	formatter := &MessageFormatter{Prefix: ">>", Suffix: "<<"}
	WriteFormattedMessages(&buffer, messages, formatter)

	got := buffer.String()
	want := ">>Hello, Go<<\n>>Hello, World<<\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
