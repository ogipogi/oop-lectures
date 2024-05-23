
# Unit Testing in Go

## Slide 1: Title Slide
**Title:** Unit Testing in Go  
**Subtitle:** Ensuring Code Quality with Tests  
**Source:** https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

## Slide 2: Introduction to Unit Testing
- **Objective:**  
  Learn how to write and run unit tests in Go.
- **Benefits:**  
  - Ensures code correctness
  - Facilitates refactoring
  - Documents code behavior

## Slide 3: Hello, World
- **First Program:**
  ```go
  package main

  import "fmt"

  func main() {
      fmt.Println("Hello, world")
  }
  ```
- **Run:**  
  ```sh
  go run hello.go
  ```

## Slide 4: Separating Concerns
- **Refactor for Testability:**
  ```go
  package main

  import "fmt"

  func Hello() string {
      return "Hello, world"
  }

  func main() {
      fmt.Println(Hello())
  }
  ```

## Slide 5: Writing the Test
- **First Test:**
  ```go
  package main

  import "testing"

  func TestHello(t *testing.T) {
      got := Hello()
      want := "Hello, world"

      if got != want {
          t.Errorf("got %q want %q", got, want)
      }
  }
  ```
- **Run:**  
  ```sh
  go test
  ```

## Slide 6: Handling Modules
- **Module Initialization:**
  ```sh
  go mod init hello
  ```
- **Module File:**
  ```go
  module hello

  go 1.16
  ```

## Slide 7: Enhancing Tests
- **Adding Subtests:**
  ```go
  func TestHello(t *testing.T) {
      t.Run("saying hello to people", func(t *testing.T) {
          got := Hello("Chris")
          want := "Hello, Chris"

          if got != want {
              t.Errorf("got %q want %q", got, want)
          }
      })

      t.Run("empty string defaults to 'World'", func(t *testing.T) {
          got := Hello("")
          want := "Hello, World"

          if got != want {
              t.Errorf("got %q want %q", got, want)
          }
      })
  }
  ```

## Slide 8: Updating the Hello Function
- **Support for Names:**
  ```go
  const englishHelloPrefix = "Hello, "

  func Hello(name string) string {
      if name == "" {
          name = "World"
      }
      return englishHelloPrefix + name
  }
  ```
- **Run Tests:**  
  ```sh
  go test
  ```

## Slide 9: Refactoring Tests
- **Helper Function:**
  ```go
  func assertCorrectMessage(t testing.TB, got, want string) {
      t.Helper()
      if got != want {
          t.Errorf("got %q want %q", got, want)
      }
  }
  ```

## Slide 10: Introducing Languages
- **New Requirement:**
  ```go
  const spanish = "Spanish"
  const spanishHelloPrefix = "Hola, "

  func Hello(name string, language string) string {
      if name == "" {
          name = "World"
      }

      if language == spanish {
          return spanishHelloPrefix + name
      }

      return englishHelloPrefix + name
  }
  ```

## Slide 11: Adding More Languages
- **French Support:**
  ```go
  const french = "French"
  const frenchHelloPrefix = "Bonjour, "

  func Hello(name string, language string) string {
      if name == "" {
          name = "World"
      }

      switch language {
      case spanish:
          return spanishHelloPrefix + name
      case french:
          return frenchHelloPrefix + name
      default:
          return englishHelloPrefix + name
      }
  }
  ```

## Slide 12: Refactoring with Constants
- **Improved Code:**
  ```go
  const (
      spanish           = "Spanish"
      french            = "French"
      englishHelloPrefix = "Hello, "
      spanishHelloPrefix = "Hola, "
      frenchHelloPrefix  = "Bonjour, "
  )

  func Hello(name string, language string) string {
      if name == "" {
          name = "World"
      }

      return greetingPrefix(language) + name
  }

  func greetingPrefix(language string) (prefix string) {
      switch language {
      case french:
          prefix = frenchHelloPrefix
      case spanish:
          prefix = spanishHelloPrefix
      default:
          prefix = englishHelloPrefix
      }
      return
  }
  ```

## Slide 13: Summary
- **Key Concepts:**
  - Writing and running tests
  - Using subtests for better organization
  - Refactoring code and tests for clarity and maintainability
