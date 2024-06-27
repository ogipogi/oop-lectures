# Concurrency Aufgaben mit Lösungen

## Aufgabe 1: Dummy-Funktion mit Sleep
- **Schreibe einen Test für eine Dummy-Funktion, die 1 Sekunde schläft und dann den Wert verdoppelt zurückgibt.**
  - Erwartung: Die Funktion soll den doppelten Wert der gegebenen Zahl nach 1 Sekunde zurückgeben.
  ```go
  func TestDummyFunction(t *testing.T) {
      got := DummyFunction(2)
      want := 4

      if got != want {
          t.Errorf("got %d want %d", got, want)
      }
  }
  ```
- **Erstelle dann die Funktion.**
  ```go
  func DummyFunction(n int) int {
      time.Sleep(1 * time.Second)
      return n * 2
  }
  ```
- **Führe den Test aus, um sicherzustellen, dass er besteht.**
  ```sh
  go test
  ```

## Aufgabe 2: Paralleles Ausführen der Dummy-Funktion mit Channels
- **Schreibe einen Test, der überprüft, ob die Funktion die Werte einer Liste von Zahlen verdoppelt.**
  - Erwartung: Die Funktion soll eine Liste der verdoppelten Werte der gegebenen Zahlen zurückgeben.
  ```go
  func TestConcurrentDummyFunction(t *testing.T) {
      numbers := []int{1, 2, 3}
      got := ConcurrentDummyFunction(numbers)
      want := []int{2, 4, 6}

      if !reflect.DeepEqual(got, want) {
          t.Errorf("got %v want %v", got, want)
      }
  }
  ```
- **Erstelle die Funktion, die Goroutines verwendet, um die Berechnung parallel durchzuführen.**
  ```go
  func ConcurrentDummyFunctionWithChannels(numbers []int) []int {
      results := make([]int, len(numbers))
      resultChannel := make(chan struct {
          index int
          value int
      })
      for i, n := range numbers {
          go func(i, n int) {
              result := DummyFunction(n)
              resultChannel <- struct {
                  index int
                  value int
              }{i, result}
          }(i, n)
      }
      for range numbers {
          result := <-resultChannel
          results[result.index] = result.value
      }
      return results
  }
  ```
- **Führe den Test aus, um sicherzustellen, dass er besteht.**
  ```sh
  go test
  ```

## Aufgabe 3: Parallelisierung und Synchronisierung testen

### Testen der Parallelisierung
- **Schreibe einen Test, der die Ausführungszeit der parallelen Version misst und sicherstellt, dass sie weniger als die Summe der Sleep-Zeiten der einzelnen Aufgaben beträgt.**
  - Erwartung: Die parallele Version soll schneller als die sequentielle Version sein.
  ```go
  func TestParallelExecution(t *testing.T) {
      numbers := []int{1, 2, 3}
      start := time.Now()
      ConcurrentDummyFunction(numbers)
      duration := time.Since(start)

      if duration >= 3*time.Second {
          t.Errorf("expected function to run in less than 3 seconds, but it took %v", duration)
      }
  }
  ```

## Aufgabe 4: Sleep wegmocken
- **Schreibe einen Test, der überprüft, ob die Funktion die Werte einer Liste von Zahlen verdoppelt, ohne tatsächlich zu schlafen.**
  - Erwartung: Die Funktion soll die Werte der gegebenen Zahlen verdoppeln, ohne tatsächlich zu schlafen.
  ```go
    type SleeperMock struct {
      Calls int
    }
    
    func (s *SleeperMock) Sleep() {
      s.Calls++
    }
  
    
  ```

  - **Führe ein Sleeper Interface ein, um das Schlafen zu mocken.**
    ```go
    type Sleeper interface {
        Sleep(time.Duration)
    }

    func DummyFunctionWithSleeper(n int, s Sleeper) int {
        s.Sleep(1 * time.Second)
        return n * 2
    }

    func TestParallelExecution(t *testing.T) {
      numbers := []int{1, 2, 3}
      start := time.Now()
      sleeper := SleeperMock{}
      ConcurrentDummyFunction(numbers, &sleeper)
      duration := time.Since(start)

      if duration >= 3*time.Second {
          t.Errorf("expected function to run in less than 3 seconds, but it took %v", duration)
      }
    }
    
    func TestParallelExecutionWithMock(t *testing.T) {
      numbers := []int{1, 2, 3}
      sleeper := SleeperMock{}
      ConcurrentDummyFunction(numbers, &sleeper)
      if sleeper.Calls != 3 {
        t.Errorf("expected function to run 3 times, but it ran %d times", sleeper.Calls)
      }
    }
    func TestConcurrentDummyFunction(t *testing.T) {
      numbers := []int{1, 2, 3}
      sleeper := SleeperMock{}
      got := ConcurrentDummyFunction(numbers, &sleeper)
      want := []int{2, 4, 6}
    
      if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
      }
    }
  ```


## Aufgabe 5: Parallelisierung mit WaitGroup
- **Refaktorisiere die Funktion so, dass Waitgroups verwendet werden, um die Berechnungen zu synchronisieren.**
  ```go
  func ConcurrentDummyFunction(numbers []int) []int {
      results := make([]int, len(numbers))
      var wg sync.WaitGroup
      for i, n := range numbers {
          wg.Add(1)
          go func(i, n int) {
              defer wg.Done()
              results[i] = DummyFunction(n)
          }(i, n)
      }
      wg.Wait()
      return results
  }
  ```
  - **Test**
    ```go
    func TestParallelExecutionWithWaitGroup(t *testing.T) {
      numbers := []int{1, 2, 3}
      start := time.Now()
      sleeper := SleeperMock{}
      ConcurrentDummyFunctionWithWaitGroup(numbers, &sleeper)
      duration := time.Since(start)

      if duration >= 3*time.Second {
          t.Errorf("expected function to run in less than 3 seconds, but it took %v", duration)
    }
  }
  ```
