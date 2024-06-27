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