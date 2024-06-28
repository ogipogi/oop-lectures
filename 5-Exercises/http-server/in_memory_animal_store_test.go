package main

import "testing"

func TestInMemoryAnimalStore(t *testing.T) {
	store := NewInMemoryAnimalStore()
	store.RecordName("hund")
	store.RecordName("katze")

	assertAnimalName(t, store, "hund", "hund")
	assertAnimalName(t, store, "katze", "katze")
}

func assertAnimalName(t *testing.T, store *InMemoryAnimalStore, name, want string) {
	t.Helper()
	got := store.GetAnimalName(name)
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
