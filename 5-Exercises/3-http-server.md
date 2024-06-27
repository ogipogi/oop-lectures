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

