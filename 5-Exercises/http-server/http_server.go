// server.go
package main

import (
	"fmt"
	"net/http"
	"strings"
)

type AnimalStore interface {
	GetAnimalName(name string) string
	RecordName(name string)
}

type AnimalServer struct {
	store AnimalStore
}

func (a *AnimalServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	animal := strings.TrimPrefix(r.URL.Path, "/tiere/")
	switch r.Method {
	case http.MethodPost:
		a.store.RecordName(animal)
		w.WriteHeader(http.StatusAccepted)
	case http.MethodGet:
		a.showName(w, animal)
	}
}

func (a *AnimalServer) showName(w http.ResponseWriter, animal string) {
	name := a.store.GetAnimalName(animal)
	if name == "" {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, name)
}
