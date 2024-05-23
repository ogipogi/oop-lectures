
# Dependency Injection in Go

## Slide 1: Title Slide
**Title:** Dependency Injection in Go  
**Subtitle:** Making Your Code More Testable and Flexible  
**Source:** https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection 

## Slide 2: Introduction to Dependency Injection
- **Definition:**  
  Dependency Injection (DI) is a design pattern used to implement IoC (Inversion of Control).
- **Benefits of DI:**
    - No need for a framework
    - Simplifies design
    - Facilitates testing
    - Enables writing reusable, general-purpose functions

## Slide 3: Dependency Injection Basics
- **Function Example Without DI:**
  ```go
  func Greet(name string) {
      fmt.Printf("Hello, %s", name)
  }
  ```
- **Problem:**  
  Directly prints to `stdout`, hard to test.

## Slide 4: Injecting Dependencies
- **Concept:**  
  Inject dependencies to make code more flexible and testable.
- **Use an Interface:**  
  Accept an `io.Writer` interface instead of a concrete type.
  ```go
  func Greet(writer io.Writer, name string) {
      fmt.Fprintf(writer, "Hello, %s", name)
  }
  ```

## Slide 5: Writing the Test
- **Test Example:**
  ```go
  func TestGreet(t *testing.T) {
      buffer := bytes.Buffer{}
      Greet(&buffer, "Chris")

      got := buffer.String()
      want := "Hello, Chris"

      if got != want {
          t.Errorf("got %q want %q", got, want)
      }
  }
  ```
- **Explanation:**  
  Use `bytes.Buffer` to capture output and verify.

## Slide 6: Fixing the Code
- **Initial Code:**
  ```go
  func Greet(writer *bytes.Buffer, name string) {
      fmt.Printf("Hello, %s", name)
  }
  ```
- **Compiler Error:**  
  Requires `*bytes.Buffer`, but not useful for `os.Stdout`.

## Slide 7: Using io.Writer
- **Refactor Code:**
  ```go
  func Greet(writer io.Writer, name string) {
      fmt.Fprintf(writer, "Hello, %s", name)
  }
  ```
- **Benefit:**  
  Works with both `*bytes.Buffer` and `os.Stdout`.

## Slide 8: Practical Application
- **Main Function Example:**
  ```go
  func main() {
      Greet(os.Stdout, "Elodie")
  }
  ```
- **Flexible Use:**  
  Can also be used in HTTP handlers.

## Slide 9: HTTP Server Example
- **Code Example:**
  ```go
  func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
      Greet(w, "world")
  }

  func main() {
      log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
  }
  ```
- **Explanation:**  
  `http.ResponseWriter` implements `io.Writer`.

## Slide 10: Wrapping Up
- **Summary:**
    - **Testing:**  
      DI makes code testable by decoupling dependencies.
    - **Separation of Concerns:**  
      DI separates how data is generated from where it is sent.
    - **Reusability:**  
      Code can be reused in different contexts.

- **Mocking:**  
  Will be covered later; helps replace real dependencies in tests.

## Slide 11: Conclusion
- **Key Takeaway:**  
  Use Dependency Injection to make your Go code more testable, flexible, and reusable.
- **Next Steps:**  
  Study the Go standard library for useful interfaces like `io.Writer`.

## Slide 12: Questions & Discussion
- **Open Floor:**  
  Encourage questions and discussion on the topic.
