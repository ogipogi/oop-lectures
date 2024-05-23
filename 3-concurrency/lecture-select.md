
# Using Select in Go

## Slide 1: Title Slide
**Title:** Using Select in Go  
**Subtitle:** Managing Concurrency with Select Statements  
**Source:** https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/select

## Slide 2: Introduction to Select
- **Objective:**  
  Create a function `WebsiteRacer` to race two URLs and return the one that responds first.
- **Requirements:**
    - Use `net/http` for HTTP requests.
    - Use `net/http/httptest` for testing.
    - Implement with goroutines and `select` for synchronization.

## Slide 3: Writing the Initial Test
- **Naive Test:**
  ```go
  func TestRacer(t *testing.T) {
      slowURL := "http://www.facebook.com"
      fastURL := "http://www.quii.dev"

      want := fastURL
      got := Racer(slowURL, fastURL)

      if got != want {
          t.Errorf("got %q, want %q", got, want)
      }
  }
  ```
- **First Attempt:**
  ```go
  func Racer(a, b string) (winner string) {
      return
  }
  ```

## Slide 4: Basic Implementation
- **Measuring Response Time:**
  ```go
  func Racer(a, b string) (winner string) {
      startA := time.Now()
      http.Get(a)
      aDuration := time.Since(startA)

      startB := time.Now()
      http.Get(b)
      bDuration := time.Since(startB)

      if aDuration < bDuration {
          return a
      }

      return b
  }
  ```

## Slide 5: Problems with Initial Implementation
- **Issues:**
    - Relies on real websites, which can be slow and unreliable.
    - Difficult to test edge cases.

## Slide 6: Using Mock Servers
- **Refactored Test:**
  ```go
  func TestRacer(t *testing.T) {
      slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
          time.Sleep(20 * time.Millisecond)
          w.WriteHeader(http.StatusOK)
      }))

      fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
          w.WriteHeader(http.StatusOK)
      }))

      slowURL := slowServer.URL
      fastURL := fastServer.URL

      want := fastURL
      got := Racer(slowURL, fastURL)

      if got != want {
          t.Errorf("got %q, want %q", got, want)
      }

      slowServer.Close()
      fastServer.Close()
  }
  ```

## Slide 7: Refactoring for Clarity
- **Improved Racer Function:**
  ```go
  func Racer(a, b string) (winner string) {
      aDuration := measureResponseTime(a)
      bDuration := measureResponseTime(b)

      if aDuration < bDuration {
          return a
      }

      return b
  }

  func measureResponseTime(url string) time.Duration {
      start := time.Now()
      http.Get(url)
      return time.Since(start)
  }
  ```

## Slide 8: Simplifying Test Setup
- **Refactored Test Setup:**
  ```go
  func TestRacer(t *testing.T) {
      slowServer := makeDelayedServer(20 * time.Millisecond)
      fastServer := makeDelayedServer(0 * time.Millisecond)

      defer slowServer.Close()
      defer fastServer.Close()

      slowURL := slowServer.URL
      fastURL := fastServer.URL

      want := fastURL
      got := Racer(slowURL, fastURL)

      if got != want {
          t.Errorf("got %q, want %q", got, want)
      }
  }

  func makeDelayedServer(delay time.Duration) *httptest.Server {
      return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
          time.Sleep(delay)
          w.WriteHeader(http.StatusOK)
      }))
  }
  ```

## Slide 9: Introducing Select
- **Using Select for Concurrency:**
  ```go
  func Racer(a, b string) (winner string) {
      select {
      case <-ping(a):
          return a
      case <-ping(b):
          return b
      }
  }

  func ping(url string) chan struct{} {
      ch := make(chan struct{})
      go func() {
          http.Get(url)
          close(ch)
      }()
      return ch
  }
  ```

## Slide 10: Explanation of Ping Function
- **Ping Function:**
    - Creates a `chan struct{}`.
    - Starts a goroutine to `http.Get(url)` and closes the channel upon completion.
- **Why `struct{}`?:**
    - Minimal memory allocation.

## Slide 11: Handling Timeouts
- **Updating Racer for Timeouts:**
  ```go
  func Racer(a, b string) (winner string, error error) {
      select {
      case <-ping(a):
          return a, nil
      case <-ping(b):
          return b, nil
      case <-time.After(10 * time.Second):
          return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
      }
  }
  ```

## Slide 12: Making Timeout Configurable
- **Configurable Racer:**
  ```go
  var tenSecondTimeout = 10 * time.Second

  func Racer(a, b string) (winner string, error error) {
      return ConfigurableRacer(a, b, tenSecondTimeout)
  }

  func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
      select {
      case <-ping(a):
          return a, nil
      case <-ping(b):
          return b, nil
      case <-time.After(timeout):
          return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
      }
  }
  ```

## Slide 13: Final Tests
- **Complete Test Suite:**
  ```go
  func TestRacer(t *testing.T) {
      t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
          slowServer := makeDelayedServer(20 * time.Millisecond)
          fastServer := makeDelayedServer(0 * time.Millisecond)

          defer slowServer.Close()
          defer fastServer.Close()

          slowURL := slowServer.URL
          fastURL := fastServer.URL

          want := fastURL
          got, err := Racer(slowURL, fastURL)

          if err != nil {
              t.Fatalf("did not expect an error but got one %v", err)
          }

          if got != want {
              t.Errorf("got %q, want %q", got, want)
          }
      })

      t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
          server := makeDelayedServer(25 * time.Millisecond)

          defer server.Close()

          _, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

          if err == nil {
              t.Error("expected an error but didn't get one")
          }
      })
  }
  ```

## Slide 14: Wrapping Up
- **Key Concepts:**
    - `select` for managing multiple channel operations.
    - Using `httptest` for reliable and controllable tests.
    - Making timeouts configurable for faster test execution.
