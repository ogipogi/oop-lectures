# HTTP-Server Aufgaben

## Aufgabe 1: Einfache HTTP-Server-Funktion
- **Schreibe einen Test für eine HTTP-Server-Funktion, die den Namen eines Tieres zurückgibt.**
  - Erwartung: Die Funktion soll "Hund" für die URL "/tiere/hund" zurückgeben.
- **Erstelle dann die Funktion.**
- **Führe den Test aus, um sicherzustellen, dass er besteht.**

## Aufgabe 2: Dynamische Tiernamen
- **Schreibe einen Test, der überprüft, ob die Funktion unterschiedliche Tiernamen korrekt zurückgibt.**
  - Erwartung: Die Funktion soll "Hund" für "/tiere/hund" und "Katze" für "/tiere/katze" zurückgeben.
- **Modifiziere die Funktion, um dynamische Tiernamen zu unterstützen.**
- **Führe den Test aus, um sicherzustellen, dass er besteht.**

## Aufgabe 3: Tieralter speichern
- **Schreibe einen Test, der überprüft, ob eine POST-Anfrage das Alter eines Tieres speichert und `http.StatusAccepted` zurückgibt.**
  - Erwartung: Die Funktion soll `http.StatusAccepted` für POST-Anfragen zurückgeben.
- **Modifiziere die Funktion, um POST-Anfragen zu unterstützen.**
- **Führe den Test aus, um sicherzustellen, dass er besteht.**

## Aufgabe 4: Code-Refaktorisierung
- **Refaktoriere den Code, um ihn klarer und besser strukturiert zu machen.**
  - Erwartung: Der Code soll sauber und gut organisiert sein.
- **Führe den Test aus, um sicherzustellen, dass er besteht.**

## Aufgabe 5: Einführung des TierStore Interface
- **Erstelle ein Interface zur Handhabung der Tierdaten.**
  - Erwartung: Das Interface soll Methoden zur Abfrage und Speicherung von Tierdaten enthalten.
- **Modifiziere den Server, um dieses Interface zu verwenden.**
- **Führe den Test aus, um sicherzustellen, dass er besteht.**

## Aufgabe 6: In-Memory Tier Store Implementierung
- **Schreibe Tests für die In-Memory Implementierung des TierStores.**
  - Erwartung: Die In-Memory Implementierung soll Tierdaten speichern und abrufen können.
- **Erstelle die In-Memory Implementierung.**
- **Führe die Tests aus, um sicherzustellen, dass sie bestehen.**

## Aufgabe 7: Integrationstests
- **Schreibe Integrationstests, um sicherzustellen, dass der Server und der In-Memory TierStore zusammenarbeiten.**
  - Erwartung: Der Server soll korrekt mit dem In-Memory TierStore interagieren.
- **Führe die Integrationstests aus, um sicherzustellen, dass sie bestehen.**

# HTTP-Server Aufgaben mit Lösungen

## Aufgabe 1: Einfache HTTP-Server-Funktion
- **Schreibe einen Test für eine HTTP-Server-Funktion, die den Namen eines Tieres zurückgibt.**
  - Erwartung: Die Funktion soll "Hund" für die URL "/tiere/hund" zurückgeben.
  ```go
  // server_test.go
  package main

  import (
      "net/http"
      "net/http/httptest"
      "testing"
  )

  func TestGETAnimals(t *testing.T) {
      t.Run("returns the name of the animal", func(t *testing.T) {
          request, _ := http.NewRequest(http.MethodGet, "/tiere/hund", nil)
          response := httptest.NewRecorder()

          AnimalServer(response, request)

          got := response.Body.String()
          want := "Hund"

          if got != want {
              t.Errorf("got %q, want %q", got, want)
          }
      })
  }
  ```
- **Erstelle dann die Funktion.**
  ```go
  // server.go
  package main

  import (
      "fmt"
      "net/http"
  )

  func AnimalServer(w http.ResponseWriter, r *http.Request) {
      fmt.Fprint(w, "Hund")
  }

  func main() {
      http.ListenAndServe(":5000", http.HandlerFunc(AnimalServer))
  }
  ```
- **Führe den Test aus, um sicherzustellen, dass er besteht.**
  ```sh
  go test
  ```

## Aufgabe 2: Dynamische Tiernamen
- **Schreibe einen Test, der überprüft, ob die Funktion unterschiedliche Tiernamen korrekt zurückgibt.**
    - Erwartung: Die Funktion soll "Hund" für "/tiere/hund" und "Katze" für "/tiere/katze" zurückgeben.
  ```go
  func TestGETAnimals(t *testing.T) {
      t.Run("returns the name of the animal Hund", func(t *testing.T) {
          request, _ := http.NewRequest(http.MethodGet, "/tiere/hund", nil)
          response := httptest.NewRecorder()

          AnimalServer(response, request)

          got := response.Body.String()
          want := "Hund"

          if got != want {
              t.Errorf("got %q, want %q", got, want)
          }
      })

      t.Run("returns the name of the animal Katze", func(t *testing.T) {
          request, _ := http.NewRequest(http.MethodGet, "/tiere/katze", nil)
          response := httptest.NewRecorder()

          AnimalServer(response, request)

          got := response.Body.String()
          want := "Katze"

          if got != want {
              t.Errorf("got %q, want %q", got, want)
          }
      })
  }
  ```
- **Modifiziere die Funktion, um dynamische Tiernamen zu unterstützen.**
  ```go
  func AnimalServer(w http.ResponseWriter, r *http.Request) {
      animal := r.URL.Path[len("/tiere/"):]
      if animal == "hund" {
          fmt.Fprint(w, "Hund")
      } else if animal == "katze" {
          fmt.Fprint(w, "Katze")
      }
  }
  ```
- **Führe den Test aus, um sicherzustellen, dass er besteht.**
  ```sh
  go test
  ```

## Aufgabe 3: Tieralter speichern
- **Schreibe einen Test, der überprüft, ob eine POST-Anfrage das Alter eines Tieres speichert und `http.StatusAccepted` zurückgibt.**
    - Erwartung: Die Funktion soll `http.StatusAccepted` für POST-Anfragen zurückgeben.
  ```go
  func TestStoreAge(t *testing.T) {
      request, _ := http.NewRequest(http.MethodPost, "/tiere/hund", nil)
      response := httptest.NewRecorder()

      AnimalServer(response, request)

      if response.Code != http.StatusAccepted {
          t.Errorf("got status %d want %d", response.Code, http.StatusAccepted)
      }
  }
  ```
- **Modifiziere die Funktion, um POST-Anfragen zu unterstützen.**
  ```go
  func AnimalServer(w http.ResponseWriter, r *http.Request) {
      if r.Method == http.MethodPost {
          w.WriteHeader(http.StatusAccepted)
      } else {
          animal := r.URL.Path[len("/tiere/"):]
          if animal == "hund" {
              fmt.Fprint(w, "Hund")
          } else if animal == "katze" {
              fmt.Fprint(w, "Katze")
          }
      }
  }
  ```
- **Führe den Test aus, um sicherzustellen, dass er besteht.**
  ```sh
  go test
  ```

## Aufgabe 4: Code-Refaktorisierung
- **Refaktoriere den Code, um ihn klarer und besser strukturiert zu machen.**
    - Erwartung: Der Code soll sauber und gut organisiert sein.
  ```go
  // server.go
  package main

  import (
      "fmt"
      "net/http"
      "strings"
  )

  func AnimalServer(w http.ResponseWriter, r *http.Request) {
      animal := strings.TrimPrefix(r.URL.Path, "/tiere/")
      switch r.Method {
      case http.MethodPost:
          w.WriteHeader(http.StatusAccepted)
      case http.MethodGet:
          fmt.Fprint(w, GetAnimalName(animal))
      }
  }

  func GetAnimalName(name string) string {
      if name == "hund" {
          return "Hund"
      }
      if name == "katze" {
          return "Katze"
      }
      return ""
  }

  func main() {
      http.ListenAndServe(":5000", http.HandlerFunc(AnimalServer))
  }
  ```
- **Führe den Test aus, um sicherzustellen, dass er besteht.**
  ```sh
  go test
  ```

## Aufgabe 5: Einführung des TierStore Interface
- **Erstelle ein Interface zur Handhabung der Tierdaten.**
    - Erwartung: Das Interface soll Methoden zur Abfrage und Speicherung von Tierdaten enthalten.
  ```go
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
  ```
- **Modifiziere den Server, um dieses Interface zu verwenden.**
- **Führe den Test aus, um sicherzustellen, dass er besteht.**
  ```go
  // server_test.go
  package main

  import (
      "net/http"
      "net/http/httptest"
      "testing"
  )

  type StubAnimalStore struct {
      names map[string]string
      nameCalls []string
  }

  func (s *StubAnimalStore) GetAnimalName(name string) string {
      return s.names[name]
  }

  func (s *StubAnimalStore) RecordName(name string) {
      s.nameCalls = append(s.nameCalls, name)
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
  ```

## Aufgabe 6: In-Memory Tier Store Implementierung
- **Schreibe Tests für die In-Memory Implementierung des TierStores.**
    - Erwartung: Die In-Memory Implementierung soll Tierdaten speichern und abrufen können.
  ```go
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
  ```

- **Erstelle die In-Memory Implementierung.**
- **Führe die Tests aus, um sicherzustellen, dass sie bestehen.**
  ```go
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
  ```

## Aufgabe 7: Integrationstests
- **Schreibe Integrationstests, um sicherzustellen, dass der Server und der In-Memory TierStore zusammenarbeiten.**
    - Erwartung: Der Server soll korrekt mit dem In-Memory TierStore interagieren.
  ```go
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

      server.ServeHTTP(httptest.NewRecorder(), newPostAgeRequest("hund"))
      server.ServeHTTP(httptest.NewRecorder(), newPostAgeRequest("katze"))
      server.ServeHTTP(httptest.NewRecorder(), newPostAgeRequest("maus"))

      response := httptest.NewRecorder()
      server.ServeHTTP(response, newGetNameRequest("hund"))
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
  ```

- **Führe die Integrationstests aus, um sicherzustellen, dass sie bestehen.**
  ```go
  // main.go
  package main

  import (
      "log"
      "net/http"
  )

  func main() {
      server := &AnimalServer{NewInMemoryAnimalStore()}
      log.Fatal(http.ListenAndServe(":5000", server))
  }
  ```