### Goal 1: Create a Basic HTTP Server
**Short Goal:** Set up a basic HTTP server in Go.

#### Test
```go
// server_test.go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestGETPlayers(t *testing.T) {
    t.Run("returns Pepper's score", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
        response := httptest.NewRecorder()

        PlayerServer(response, request)

        got := response.Body.String()
        want := "20"

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
    })
}
```
**Explanation:** We start by writing a test for a GET request to `/players/{name}` which should return the number of wins for that player. This test uses `httptest` to simulate the HTTP request and response.

#### Implementation
```go
// server.go
package main

import (
    "fmt"
    "net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "20")
}

func main() {
    http.ListenAndServe(":5000", http.HandlerFunc(PlayerServer))
}
```
**Explanation:** The implementation defines the `PlayerServer` function to respond with "20" for any GET request. The `main` function starts the HTTP server on port 5000.

### Goal 2: Handle Dynamic Player Names
**Short Goal:** Modify the server to handle dynamic player names.

#### Test
```go
func TestGETPlayers(t *testing.T) {
    t.Run("returns Pepper's score", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
        response := httptest.NewRecorder()

        PlayerServer(response, request)

        got := response.Body.String()
        want := "20"

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
    })

    t.Run("returns Floyd's score", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
        response := httptest.NewRecorder()

        PlayerServer(response, request)

        got := response.Body.String()
        want := "10"

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
    })
}
```
**Explanation:** The test is extended to check responses for two different player names, expecting different scores.

#### Implementation
```go
func PlayerServer(w http.ResponseWriter, r *http.Request) {
    player := r.URL.Path[len("/players/"):]
    if player == "Pepper" {
        fmt.Fprint(w, "20")
    } else if player == "Floyd" {
        fmt.Fprint(w, "10")
    }
}
```
**Explanation:** The `PlayerServer` function is updated to return different scores based on the player name extracted from the URL.

### Goal 3: Record Player Wins
**Short Goal:** Implement the POST method to record player wins.

#### Test
```go
func TestStoreWins(t *testing.T) {
    request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
    response := httptest.NewRecorder()

    PlayerServer(response, request)

    if response.Code != http.StatusAccepted {
        t.Errorf("got status %d want %d", response.Code, http.StatusAccepted)
    }
}
```
**Explanation:** A new test checks if a POST request to `/players/{name}` returns `http.StatusAccepted`.

#### Implementation
```go
func PlayerServer(w http.ResponseWriter, r *http.Request) {
    player := r.URL.Path[len("/players/"):]
    if r.Method == http.MethodPost {
        w.WriteHeader(http.StatusAccepted)
    } else {
        if player == "Pepper" {
            fmt.Fprint(w, "20")
        } else if player == "Floyd" {
            fmt.Fprint(w, "10")
        }
    }
}
```
**Explanation:** The `PlayerServer` function is updated to handle POST requests by returning `http.StatusAccepted`.

### Goal 4: Refactor Code for Clarity
**Short Goal:** Refactor the server code for better structure and clarity.

#### Implementation
```go
// server.go
package main

import (
    "fmt"
    "net/http"
    "strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
    player := strings.TrimPrefix(r.URL.Path, "/players/")
    switch r.Method {
    case http.MethodPost:
        w.WriteHeader(http.StatusAccepted)
    case http.MethodGet:
        fmt.Fprint(w, GetPlayerScore(player))
    }
}

func GetPlayerScore(name string) string {
    if name == "Pepper" {
        return "20"
    }
    if name == "Floyd" {
        return "10"
    }
    return ""
}

func main() {
    http.ListenAndServe(":5000", http.HandlerFunc(PlayerServer))
}
```
**Explanation:** The code is refactored to extract player score logic into a separate function `GetPlayerScore`, improving readability and maintainability.

### Goal 5: Introduce PlayerStore Interface
**Short Goal:** Use an interface to handle player data storage and retrieval.

#### Implementation
```go
// server.go
package main

import (
    "fmt"
    "net/http"
    "strings"
)

type PlayerStore interface {
    GetPlayerScore(name string) int
    RecordWin(name string)
}

type PlayerServer struct {
    store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    player := strings.TrimPrefix(r.URL.Path, "/players/")
    switch r.Method {
    case http.MethodPost:
        w.WriteHeader(http.StatusAccepted)
    case http.MethodGet:
        p.showScore(w, player)
    }
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
    score := p.store.GetPlayerScore(player)
    if score == 0 {
        w.WriteHeader(http.StatusNotFound)
    }
    fmt.Fprint(w, score)
}
```
**Explanation:** The `PlayerStore` interface abstracts data storage, and `PlayerServer` now depends on this interface. The `ServeHTTP` method routes requests appropriately.

### Goal 6: Implement In-Memory Player Store
**Short Goal:** Create an in-memory implementation of `PlayerStore`.

#### Test
```go
// server_test.go
package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

type StubPlayerStore struct {
    scores   map[string]int
    winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
    return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
    s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
    store := StubPlayerStore{
        scores: map[string]int{
            "Pepper": 20,
            "Floyd":  10,
        },
    }
    server := &PlayerServer{&store}

    t.Run("returns Pepper's score", func(t *testing.T) {
        request := newGetScoreRequest("Pepper")
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        assertResponseBody(t, response.Body.String(), "20")
    })

    t.Run("returns Floyd's score", func(t *testing.T) {
        request := newGetScoreRequest("Floyd")
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        assertResponseBody(t, response.Body.String(), "10")
    })
}

func newGetScoreRequest(name string) *http.Request {
    req, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
    return req
}

func assertResponseBody(t *testing.T, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("response body is wrong, got %q want %q", got, want)
    }
}
```

**Implementation:**
```go
// in_memory_player_store.go
package main

type InMemoryPlayerStore struct {
    store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
    return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
    return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
    i.store[name]++
}
```

### Goal 7: Integration Tests
**Short Goal:** Ensure the integration between `PlayerServer` and `InMemoryPlayerStore` works.

#### Test
```go
// server_integration_test.go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
    store := NewInMemoryPlayerStore()
    server := &PlayerServer{store}
    player := "Pepper"

    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

    response := httptest.NewRecorder()
    server.ServeHTTP(response, newGetScoreRequest(player))
    assertStatus(t, response.Code, http.StatusOK)

    assertResponseBody(t, response.Body.String(), "3")
}

func newPostWinRequest(name string) *http.Request {
    req, _ := http.NewRequest(http.MethodPost, "/players/"+name, nil)
    return req
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
```

**Implementation:**
```go
// main.go
package main

import (
    "log"
    "net/http"
)

func main() {
    server := &PlayerServer{NewInMemoryPlayerStore()}
    log.Fatal(http.ListenAndServe(":5000", server))
}
```

### Summary
By following these iterative steps, we ensure that each part of the server is implemented and tested before moving on to the next. This approach helps maintain focus and clarity, making the tutorial more digestible for students. Each step matches the expected outcomes from the original tutorial.
