
# Mocking in Go

## Slide 1: Title Slide
**Title:** Mocking in Go  
**Subtitle:** Enhancing Your Tests with Mocking Techniques  
**Source:** https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking

## Slide 2: Introduction to Mocking
- **Definition:**  
  Mocking is a technique used in testing to replace real objects with mock objects to simulate behavior.
- **Purpose:**  
  - Verify interactions between components  
  - Control external dependencies  
  - Facilitate testing complex scenarios  

## Slide 3: Problem Statement
- **Goal:**  
  Write a program to count down from 3, printing each number on a new line with a 1-second pause, and finally print "Go!"  
  ```
  3
  2
  1
  Go!
  ```

## Slide 4: Initial Function
- **Basic Structure:**  
  ```go
  package main

  func main() {
      Countdown()
  }
  ```
- **Iterative Approach:**  
  Take small steps to build the functionality incrementally.

## Slide 5: Writing the Test
- **First Test:**  
  ```go
  func TestCountdown(t *testing.T) {
      buffer := &bytes.Buffer{}

      Countdown(buffer)

      got := buffer.String()
      want := "3"

      if got != want {
          t.Errorf("got %q want %q", got, want)
      }
  }
  ```
- **Explanation:**  
  Use `bytes.Buffer` to capture output for verification.

## Slide 6: Minimal Implementation
- **Define Countdown:**  
  ```go
  func Countdown(out *bytes.Buffer) {}
  ```
- **Run the Test:**  
  ```go
  ./countdown_test.go:11:11: too many arguments in call to Countdown
      have (*bytes.Buffer)
      want ()
  ```

## Slide 7: Adjust Function Signature
- **Update Countdown:**  
  ```go
  func Countdown(out *bytes.Buffer) {}
  ```
- **Test Result:**  
  ```go
  countdown_test.go:17: got '' want '3'
  ```

## Slide 8: Implement Functionality
- **Make the Test Pass:**  
  ```go
  func Countdown(out *bytes.Buffer) {
      fmt.Fprint(out, "3")
  }
  ```
- **Use fmt.Fprint:**  
  Takes an `io.Writer` and sends a string to it.

## Slide 9: Refactor to io.Writer
- **Generalize Implementation:**  
  ```go
  func Countdown(out io.Writer) {
      fmt.Fprint(out, "3")
  }
  ```
- **Re-run Tests:**  
  Ensure tests still pass.

## Slide 10: Expand Functionality
- **Update Main:**  
  ```go
  package main

  import (
      "fmt"
      "io"
      "os"
  )

  func Countdown(out io.Writer) {
      fmt.Fprint(out, "3")
  }

  func main() {
      Countdown(os.Stdout)
  }
  ```
- **Print 3, 2, 1, Go!**  
  ```go
  func Countdown(out io.Writer) {
      for i := 3; i > 0; i-- {
          fmt.Fprintln(out, i)
      }
      fmt.Fprint(out, "Go!")
  }
  ```

## Slide 11: Adding Sleep
- **Pause Between Prints:**  
  ```go
  func Countdown(out io.Writer) {
      for i := 3; i > 0; i-- {
          fmt.Fprintln(out, i)
          time.Sleep(1 * time.Second)
      }
      fmt.Fprint(out, "Go!")
  }
  ```

## Slide 12: Mocking the Sleeper
- **Dependency on Sleep:**  
  Slow tests ruin productivity.
- **Extract Dependency:**  
  ```go
  type Sleeper interface {
      Sleep()
  }
  ```

## Slide 13: Implement Spy
- **SpySleeper:**  
  ```go
  type SpySleeper struct {
      Calls int
  }

  func (s *SpySleeper) Sleep() {
      s.Calls++
  }
  ```
- **Update Test:**  
  ```go
  func TestCountdown(t *testing.T) {
      buffer := &bytes.Buffer{}
      spySleeper := &SpySleeper{}

      Countdown(buffer, spySleeper)

      got := buffer.String()
      want := `3
  2
  1
  Go!`

      if got != want {
          t.Errorf("got %q want %q", got, want)
      }

      if spySleeper.Calls != 3 {
          t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
      }
  }
  ```

## Slide 14: Adjust Countdown Function
- **Update Signature:**  
  ```go
  func Countdown(out io.Writer, sleeper Sleeper) {
      for i := 3; i > 0; i-- {
          fmt.Fprintln(out, i)
          sleeper.Sleep()
      }
      fmt.Fprint(out, "Go!")
  }
  ```
- **Main Function:**  
  ```go
  func main() {
      sleeper := &DefaultSleeper{}
      Countdown(os.Stdout, sleeper)
  }
  ```

## Slide 15: Refactor to Configurable Sleeper
- **ConfigurableSleeper:**  
  ```go
  type ConfigurableSleeper struct {
      duration time.Duration
      sleep    func(time.Duration)
  }

  func (c *ConfigurableSleeper) Sleep() {
      c.sleep(c.duration)
  }
  ```

## Slide 16: Update Main
- **Use ConfigurableSleeper:**  
  ```go
  func main() {
      sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
      Countdown(os.Stdout, sleeper)
  }
  ```

