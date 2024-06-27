# Concurrency Aufgaben

## Aufgabe 1: Dummy-Funktion mit Sleep
- **Schreibe einen Test für eine Dummy-Funktion, die 1 Sekunde schläft und dann den Wert verdoppelt zurückgibt.**
    - Erwartung: Die Funktion soll den doppelten Wert der gegebenen Zahl nach 1 Sekunde zurückgeben.
- **Erstelle dann die Funktion.**
- **Führe den Test aus, um sicherzustellen, dass er besteht.**

## Aufgabe 2: Synchronisierung mit Channels
- **Refaktorisiere die Funktion so, dass Channels verwendet werden, um die Berechnungen zu synchronisieren.**
- **Führe den Test aus, um sicherzustellen, dass er weiterhin besteht.**

## Aufgabe 3: Parallelisierung und Synchronisierung testen
- **Testen der Parallelisierung:** Schreibe einen Test, der die Ausführungszeit der parallelen Version misst und sicherstellt, dass sie weniger als die Summe der Sleep-Zeiten der einzelnen Aufgaben beträgt.
    - Erwartung: Die parallele Version soll schneller als die sequentielle Version sein.

## Aufgabe 4: Sleep wegmocken
- **Schreibe einen Test, der überprüft, ob die Funktion die Werte einer Liste von Zahlen verdoppelt, ohne tatsächlich zu schlafen.**
    - Erwartung: Die Funktion soll die Werte der gegebenen Zahlen verdoppeln, ohne tatsächlich zu schlafen.
- **Führe ein Sleeper Interface ein, um das Schlafen zu mocken.**

## Aufgabe 5: Paralleles Ausführen der Dummy-Funktion mit WaitGroup (nicht Prüfungsrelevant)
- **Schreibe einen Test, der überprüft, ob die Funktion die Werte einer Liste von Zahlen gleichzeitig verdoppelt.**
  - Erwartung: Die Funktion soll eine Liste der verdoppelten Werte der gegebenen Zahlen zurückgeben, wobei jede Berechnung parallel erfolgt.
- **Erstelle die Funktion, die Goroutines verwendet, um die Berechnung parallel durchzuführen.**
- **Führe den Test aus, um sicherzustellen, dass er besteht.**
