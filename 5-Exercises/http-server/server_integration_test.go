// server_integration_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingAgesAndRetrievingThem(t *testing.T) {
	store := NewInMemoryAnimalStore()
	server := &AnimalServer{store}
	animal := "hund"

	server.ServeHTTP(httptest.NewRecorder(), newPostAgeRequest(animal))
	server.ServeHTTP(httptest.NewRecorder(), newPostAgeRequest(animal))
	server.ServeHTTP(httptest.NewRecorder(), newPostAgeRequest(animal))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetNameRequest(animal))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "hund")
}

func newPostAgeRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/tiere/"+name, nil)
	return req
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
