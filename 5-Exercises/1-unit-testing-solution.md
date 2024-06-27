# TDD Aufgaben

## Aufgabe 1: Schreibe einen einfachen Test
- **Schreibe einen Test für eine Funktion, die "Hello, world" zurückgibt.**
  - Erwartung: Die Funktion soll "Hello, world" zurückgeben.
- **Erstelle dann die Funktion.**
- **Führe den Test aus, um sicherzustellen, dass er besteht.**

## Aufgabe 2: Refaktoriere für Testbarkeit
- **Schreibe einen Test, der überprüft, ob die Funktion eine personalisierte Begrüssung zurückgeben kann.**
  - Erwartung: Die Funktion soll "Hello, [Name]" zurückgeben, wenn ein Name angegeben ist.
- **Refaktoriere die Funktion entsprechend.**
- **Führe den Test aus, um sicherzustellen, dass er besteht.**

## Aufgabe 3: Füge Unterstützung für Standardwerte hinzu
- **Schreibe einen Test, der überprüft, ob die Funktion "Hello, World" zurückgibt, wenn kein Name angegeben ist.**
  - Erwartung: Die Funktion soll "Hello, World" zurückgeben, wenn der Name leer ist.
- **Modifiziere die Funktion entsprechend.**
- **Führe den Test aus, um die Änderungen zu überprüfen.**

## Aufgabe 4: Implementiere Unterstützung für verschiedene Begrüssungen
- **Schreibe einen Test, der überprüft, ob die Funktion verschiedene Begrüssungen zurückgeben kann, z.B. "Good morning, [Name]", "Good evening, [Name]".**
  - Erwartung: Die Funktion soll verschiedene Begrüssungen basierend auf einem zusätzlichen Parameter zurückgeben.
- **Erweitere die Funktion entsprechend.**
- **Füge Subtests hinzu, um verschiedene Begrüssungen zu überprüfen.**
- **Führe die Tests aus, um sicherzustellen, dass sie bestehen.**

## Aufgabe 5: Refaktoriere mit Konstanten
- **Refaktoriere die Begrüssungsfunktion, um Konstanten für die Begrüssungsprefixe zu verwenden.**
- **Aktualisiere die Tests, um diese Änderungen zu berücksichtigen.**
- **Führe die Tests aus, um sicherzustellen, dass sie noch bestehen.**

## Aufgabe 6: Implementiere Dependency Injection
- **Schreibe einen Test für eine Funktion, die eine Nachricht in einen `io.Writer` schreibt.**
  - Erwartung: Die Funktion soll die Nachricht in den `writer` schreiben.
- **Erstelle dann die Funktion.**
- **Führe den Test aus, um sicherzustellen, dass er besteht.**

## Aufgabe 7: Erweitere die Funktionalität
- **Schreibe einen Test, der überprüft, ob die Funktion eine Liste von Nachrichten in einen `io.Writer` schreibt.**
  - Erwartung: Die Funktion soll jede Nachricht in die neue Zeile schreiben.
- **Modifiziere die Funktion entsprechend.**
- **Führe den Test aus, um sicherzustellen, dass er besteht.**

## Aufgabe 8: Verwende Mocking für Bank-Einzahlungen
- **Schreibe einen Test für eine Funktion, die eine Einzahlung in ein Bankkonto vornimmt.**
  - Erwartung: Die Funktion soll den Betrag zum Konto hinzufügen und den neuen Kontostand zurückgeben.
- **Implementiere ein Mock Bankkonto.**
- **Schreibe den Test und überprüfe das Verhalten der Funktion.**
- **Stelle sicher, dass der Test besteht.**

## Aufgabe 9: Erstelle eine konfigurierbare Funktion
- **Schreibe einen Test für eine konfigurierbare Funktion, die verschiedene Nachrichtenformate unterstützt.**
  - Erwartung: Die Funktion soll eine Nachricht mit einem konfigurierbaren Präfix und Suffix zurückgeben.
- **Erstelle die Funktion und modifiziere sie entsprechend.**
- **Schreibe Tests, um die Funktionalität zu überprüfen.**
- **Stelle sicher, dass alle Tests bestehen.**

## Markdown 2: TDD Aufgaben mit Lösungen


# TDD Aufgaben mit Lösungen

## Aufgabe 1: Schreibe einen einfachen Test
- **Schreibe einen Test für eine Funktion, die "Hello, world" zurückgibt.**
  - Erwartung: Die Funktion soll "Hello, world" zurückgeben.
  ```go
  func TestHello(t *testing.T) {
      got := Hello()
      want := "Hello, world"

      if got != want {
          t.Errorf("got %q want %q", got, want)
      }
  }
  ```
- **Erstelle dann die Funktion.**
  ```go
  func Hello() string {
      return "Hello, world"
  }
  ```
- **Führe den Test aus, um sicherzustellen, dass er besteht.**
  ```sh
  go test
  ```

## Aufgabe 2: Refaktoriere für Testbarkeit
- **Schreibe einen Test, der überprüft, ob die Funktion eine personalisierte Begrüssung zurückgeben kann.**
    - Erwartung: Die Funktion soll "Hello, [Name]" zurückgeben, wenn ein Name angegeben ist.
  ```go
  func TestHello(t *testing.T) {
      t.Run("saying hello to people", func(t *testing.T) {
          got := Hello("Anna")
          want := "Hello, Anna"
          if got != want {
              t.Errorf("got %q want %q", got, want)
          }
      })
  }
  ```
- **Refaktoriere die Funktion entsprechend.**
  ```go
  func Hello(name string) string {
      return "Hello, " + name
  }
  ```
- **Führe den Test aus, um sicherzustellen, dass er besteht.**
  ```sh
  go test
  ```

## Aufgabe 3: Füge Unterstützung für Standardwerte hinzu
- **Schreibe einen Test, der überprüft, ob die Funktion "Hello, World" zurückgibt, wenn kein Name angegeben ist.**
    - Erwartung: Die Funktion soll "Hello, World" zurückgeben, wenn der Name leer ist.
  ```go
  t.Run("empty string defaults to 'World'", func(t *testing.T) {
      got := Hello("")
      want := "Hello, World"
      if got != want {
          t.Errorf("got %q want %q", got, want)
      }
  })
  ```
- **Modifiziere die Funktion entsprechend.**
    ```go
  func Hello(name string) string {
      if name == "" {
          name = "World"
      }
      return "Hello, " + name
  }
  ```
- **Führe den Test aus, um die Änderungen zu überprüfen.**
  ```sh
  go test
  ```

## Aufgabe 4: Implementiere Unterstützung für verschiedene Begrüssungen
- **Schreibe einen Test, der überprüft, ob die Funktion verschiedene Begrüssungen zurückgeben kann, z.B. "Good morning, [Name]", "Good evening, [Name]".**
    - Erwartung: Die Funktion soll verschiedene Begrüssungen basierend auf einem zusätzlichen Parameter zurückgeben.
  ```go
      t.Run("morning greeting", func(t *testing.T) {
          got := Hello("Anna", "morning")
          want := "Good morning, Anna"
          if got != want {
              t.Errorf("got %q want %q", got, want)
          }
      })

      t.Run("evening greeting", func(t *testing.T) {
          got := Hello("Anna", "evening")
          want := "Good evening, Anna"
          if got != want {
              t.Errorf("got %q want %q", got, want)
          }
      })
  ```
- **Erweitere die Funktion entsprechend.**
  ```go
  func Hello(name string, greetingType string) string {
      if name == "" {
          name = "World"
      }
      greeting := "Hello"
      if greetingType == "morning" {
          greeting = "Good morning"
      } else if greetingType == "evening" {
          greeting = "Good evening"
      }
      return greeting + ", " + name
  }
  ```
## Aufgabe 5: Refaktoriere mit Konstanten
- **Refaktoriere die Begrüssungsfunktion, um Konstanten für die Begrüssungsprefixe zu verwenden.**
  ```go
  const (
      helloPrefix = "Hello, "
      morningPrefix = "Good morning, "
      eveningPrefix = "Good evening, "
  )

  func Hello(name string, greetingType string) string {
      if name == "" {
          name = "World"
      }
      prefix := helloPrefix
      if greetingType == "morning" {
          prefix = morningPrefix
      } else if greetingType == "evening" {
          prefix = eveningPrefix
      }
      return prefix + name
  }
  ```

## Aufgabe 6: Implementiere Dependency Injection
- **Schreibe einen Test für eine Funktion

, die eine Nachricht in einen `io.Writer` schreibt.**
- Erwartung: Die Funktion soll die Nachricht in den `writer` schreiben.
  ```go
  func TestWriteMessage(t *testing.T) {
      buffer := bytes.Buffer{}
      WriteMessage(&buffer, "Hello, Go")

      got := buffer.String()
      want := "Hello, Go\n"

      if got != want {
          t.Errorf("got %q want %q", got, want)
      }
  }
  ```
- **Erstelle dann die Funktion.**
  ```go
  func WriteMessage(writer io.Writer, message string) {
      fmt.Fprintln(writer, message)
  }
  ```
- **Führe den Test aus, um sicherzustellen, dass er besteht.**
  ```sh
  go test
  ```

## Aufgabe 7: Erweitere die Funktionalität
- **Schreibe einen Test, der überprüft, ob die Funktion eine Liste von Nachrichten in einen `io.Writer` schreibt.**
    - Erwartung: Die Funktion soll jede Nachricht in die neue Zeile schreiben.
  ```go
  func TestWriteMessages(t *testing.T) {
      buffer := bytes.Buffer{}
      messages := []string{"Hello, Go", "Hello, World"}
      WriteMessages(&buffer, messages)

      got := buffer.String()
      want := "Hello, Go\nHello, World\n"

      if got != want {
          t.Errorf("got %q want %q", got, want)
      }
  }
  ```
- **Modifiziere die Funktion entsprechend.**
  ```go
  func WriteMessages(writer io.Writer, messages []string) {
      for _, message := range messages {
          fmt.Fprintln(writer, message)
      }
  }
  ```
- **Führe den Test aus, um sicherzustellen, dass er besteht.**
  ```sh
  go test
  ```

## Aufgabe 8: Verwende Mocking für Bank-Einzahlungen
- **Schreibe einen Test für eine Funktion, die eine Einzahlung in ein Bankkonto vornimmt.**
    - Erwartung: Die Funktion soll den Betrag zum Konto hinzufügen und den neuen Kontostand zurückgeben.
  ```go
  func TestDeposit(t *testing.T) {
      account := &MockBankAccount{}
      newBalance := Deposit(account, 100)

      want := 100
      if newBalance != want {
          t.Errorf("got %d, want %d", newBalance, want)
      }
  }
  ```
- **Implementiere ein Mock Bankkonto.**
  ```go
  type MockBankAccount struct {
      balance int
  }

  func (m *MockBankAccount) Deposit(amount int) {
      m.balance += amount
  }

  func (m *MockBankAccount) Balance() int {
      return m.balance
  }
  ```
- **Erstelle die Funktion, die die Einzahlung vornimmt.**
  ```go
  type BankAccount interface {
      Deposit(amount int)
      Balance() int
  }

  func Deposit(account BankAccount, amount int) int {
      account.Deposit(amount)
      return account.Balance()
  }
  ```
- **Schreibe den Test und überprüfe das Verhalten der Funktion.**
  Bereits im Test enthalten.
- **Stelle sicher, dass der Test besteht.**
  ```sh
  go test
  ```

## Aufgabe 9: Erstelle eine konfigurierbare Funktion
- **Schreibe einen Test für eine konfigurierbare Funktion, die verschiedene Nachrichtenformate unterstützt.**
    - Erwartung: Die Funktion soll eine Nachricht mit einem konfigurierbaren Präfix und Suffix zurückgeben.
  ```go
  func TestWriteFormattedMessages(t *testing.T) {
      buffer := bytes.Buffer{}
      messages := []string{"Hello, Go", "Hello, World"}
      formatter := &MessageFormatter{Prefix: ">>", Suffix: "<<"}
      WriteFormattedMessages(&buffer, messages, formatter)

      got := buffer.String()
      want := ">>Hello, Go<<\n>>Hello, World<<\n"

      if got != want {
          t.Errorf("got %q want %q", got, want)
      }
  }
  ```
- **Erstelle die Funktion und modifiziere sie entsprechend.**
  ```go
  type MessageFormatter struct {
      Prefix string
      Suffix string
  }

  func (mf *MessageFormatter) FormatMessage(message string) string {
      return mf.Prefix + message + mf.Suffix
  }

  func WriteFormattedMessages(writer io.Writer, messages []string, formatter *MessageFormatter) {
      for _, message := range messages {
          fmt.Fprintln(writer, formatter.FormatMessage(message))
      }
  }
  ```
