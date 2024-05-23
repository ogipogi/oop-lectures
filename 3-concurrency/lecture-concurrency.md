
# Concurrency in Go

## Slide 1: Title Slide
**Title:** Concurrency in Go  
**Subtitle:** Enhancing Performance with Concurrent Programming  
**Source:** https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency

## Slide 2: Introduction to Concurrency
- **Definition:**  
  Concurrency in Go means "having more than one thing in progress."
- **Benefits:**  
  - Faster execution  
  - Efficient resource utilization  
  - Improved performance for I/O-bound tasks  

## Slide 3: Problem Statement
- **Task:**  
  A function, CheckWebsites, checks the status of a list of URLs.
  ```go
  package concurrency

  type WebsiteChecker func(string) bool

  func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
      results := make(map[string]bool)

      for _, url := range urls {
          results[url] = wc(url)
      }

      return results
  }
  ```

## Slide 4: Initial Test
- **Test Example:**
  ```go
  package concurrency

  import (
      "reflect"
      "testing"
  )

  func mockWebsiteChecker(url string) bool {
      if url == "waat://furhurterwe.geds" {
          return false
      }
      return true
  }

  func TestCheckWebsites(t *testing.T) {
      websites := []string{
          "http://google.com",
          "http://blog.gypsydave5.com",
          "waat://furhurterwe.geds",
      }

      want := map[string]bool{
          "http://google.com":          true,
          "http://blog.gypsydave5.com": true,
          "waat://furhurterwe.geds":    false,
      }

      got := CheckWebsites(mockWebsiteChecker, websites)

      if !reflect.DeepEqual(want, got) {
          t.Fatalf("wanted %v, got %v", want, got)
      }
  }
  ```

## Slide 5: Benchmarking
- **Benchmark Test:**
  ```go
  package concurrency

  import (
      "testing"
      "time"
  )

  func slowStubWebsiteChecker(_ string) bool {
      time.Sleep(20 * time.Millisecond)
      return true
  }

  func BenchmarkCheckWebsites(b *testing.B) {
      urls := make([]string, 100)
      for i := 0; i < len(urls); i++ {
          urls[i] = "a url"
      }
      b.ResetTimer()
      for i := 0; i < b.N; i++ {
          CheckWebsites(slowStubWebsiteChecker, urls)
      }
  }
  ```

## Slide 6: Running the Benchmark
- **Benchmark Result:**
  ```sh
  pkg: github.com/gypsydave5/learn-go-with-tests/concurrency/v0
  BenchmarkCheckWebsites-4               1        2249228637 ns/op
  PASS
  ok      github.com/gypsydave5/learn-go-with-tests/concurrency/v0        2.268s
  ```
- **Observation:**  
  The function takes about two and a quarter seconds.

## Slide 7: Introducing Concurrency
- **Concept:**  
  Use goroutines to perform multiple tasks concurrently.
- **Example:**
  ```go
  package concurrency

  type WebsiteChecker func(string) bool

  func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
      results := make(map[string]bool)

      for _, url := range urls {
          go func() {
              results[url] = wc(url)
          }()
      }

      return results
  }
  ```

## Slide 8: Handling Concurrency Issues
- **Problem:**  
  Goroutines might not finish before the function returns.
- **Solution:**  
  Use a channel to synchronize.
  ```go
  package concurrency

  import (
      "time"
  )

  type WebsiteChecker func(string) bool

  func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
      results := make(map[string]bool)

      for _, url := range urls {
          go func(u string) {
              results[u] = wc(u)
          }(url)
      }

      time.Sleep(2 * time.Second)

      return results
  }
  ```

## Slide 9: Fixing Data Races
- **Race Condition:**  
  Concurrent writes to a map can cause a race condition.
- **Solution:**  
  Use a channel to ensure safe access.
  ```go
  package concurrency

  type WebsiteChecker func(string) bool
  type result struct {
      string
      bool
  }

  func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
      results := make(map[string]bool)
      resultChannel := make(chan result)

      for _, url := range urls {
          go func(u string) {
              resultChannel <- result{u, wc(u)}
          }(url)
      }

      for i := 0; i < len(urls); i++ {
          r := <-resultChannel
          results[r.string] = r.bool
      }

      return results
  }
  ```

## Slide 10: Using Channels
- **Concept:**  
  Channels allow communication between goroutines.
- **Example:**  
  ```go
  resultChannel := make(chan result)
  resultChannel <- result{u, wc(u)}
  r := <-resultChannel
  ```

## Slide 11: Final Benchmark
- **Result:**  
  ```sh
  pkg: github.com/gypsydave5/learn-go-with-tests/concurrency/v2
  BenchmarkCheckWebsites-8             100          23406615 ns/op
  PASS
  ok      github.com/gypsydave5/learn-go-with-tests/concurrency/v2        2.377s
  ```
- **Observation:**  
  The function is now about one hundred times faster.

