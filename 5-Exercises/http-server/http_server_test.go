// server_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubAnimalStore struct {
	names    map[string]string
	ageCalls []string
}

func (s *StubAnimalStore) GetAnimalName(name string) string {
	return s.names[name]
}

func (s *StubAnimalStore) RecordName(name string) {
	s.ageCalls = append(s.ageCalls, name)
}

func TestGETAnimals(t *testing.T) {
	store := StubAnimalStore{
		names: map[string]string{
			"hund":  "Hund",
			"katze": "Katze",
		},
	}
	server := &AnimalServer{&store}

	t.Run("returns the name of the animal Hund", func(t *testing.T) {
		request := newGetNameRequest("hund")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "Hund")
	})

	t.Run("returns the name of the animal Katze", func(t *testing.T) {
		request := newGetNameRequest("katze")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "Katze")
	})
}

func newGetNameRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/tiere/"+name, nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
