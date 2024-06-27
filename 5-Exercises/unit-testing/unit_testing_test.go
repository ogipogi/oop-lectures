package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Anna", "")
		want := "Hello, Anna"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("morning greeting", func(t *testing.T) {
		got := Hello("Anna", "morning")
		want := "Good morning, Anna"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("evening greeting", func(t *testing.T) {
		got := Hello("Anna", "evening")
		want := "Good evening, Anna"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
