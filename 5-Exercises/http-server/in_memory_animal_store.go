// in_memory_animal_store.go
package main

type InMemoryAnimalStore struct {
	store map[string]string
}

func NewInMemoryAnimalStore() *InMemoryAnimalStore {
	return &InMemoryAnimalStore{map[string]string{}}
}

func (i *InMemoryAnimalStore) GetAnimalName(name string) string {
	return i.store[name]
}

func (i *InMemoryAnimalStore) RecordName(name string) {
	i.store[name] = name
}
