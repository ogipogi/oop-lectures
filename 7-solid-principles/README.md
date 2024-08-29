# SOLID Prinzipien in Go

## Einführung in SOLID
Gute Softwaresysteme beginnen mit sauberem Code. Wenn die Bausteine nicht gut gemacht sind, spielt die Architektur des Gebäudes keine große Rolle. Umgekehrt kann man mit gut gemachten Bausteinen ein erhebliches Durcheinander anrichten. Hier kommen die SOLID-Prinzipien ins Spiel.

Die SOLID-Prinzipien sagen uns, wie wir unsere Funktionen und Daten anordnen sollen, und das Ziel der Prinzipien ist die Schaffung von Softwarestrukturen, die:
- Änderungen tolerieren
- Einfach zu verstehen sind
- Die Grundlage von Komponenten bilden, die in vielen Softwaresystemen verwendet werden können

## Single Responsibility Principle (SRP)
**Definition**: "A module should have one, and only one reason to change"

### Problem
```go
package main

import (
    "fmt"
    "math"
)

type Square struct {
    side float64
}

type Circle struct {
    radius float64
}

func (s Square) Area() {
    area := s.side * s.side
    fmt.Printf("The area of the square is: %f\n", area)
}

func (c Circle) Area() {
    area := math.Pi * c.radius * c.radius
    fmt.Printf("The area of the circle is: %f\n", area)
}

func main() {
    square := Square{side: 2}
    circle := Circle{radius: 3}

    square.Area()
    circle.Area()
}
```
### Erklärung
In diesem Beispiel berechnen die `Area`-Methoden sowohl die Fläche als auch die Ausgabe des Ergebnisses. Dies verletzt das SRP, da die Methoden mehr als eine Verantwortung haben.

### Lösung
```go
package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
}

type Square struct {
    side float64
}

func (s Square) Area() float64 {
    return s.side * s.side
}

type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}

type OutPrinter struct{}

func (op OutPrinter) PrintArea(shape Shape) {
    fmt.Printf("The area is: %f\n", shape.Area())
}

func main() {
    square := Square{side: 2}
    circle := Circle{radius: 3}
    printer := OutPrinter{}

    printer.PrintArea(square)
    printer.PrintArea(circle)
}
```
### Erklärung
Hier haben wir die Berechnung der Fläche und die Ausgabe des Ergebnisses getrennt. Die `Area`-Methoden berechnen nur die Fläche, während die `OutPrinter`-Klasse für die Ausgabe verantwortlich ist.

### Übung
1. Implementieren Sie eine neue Form `Rectangle` und fügen Sie sie in das bestehende System ein.
2. Schreiben Sie eine Funktion, die die Flächen mehrerer Formen summiert und das Ergebnis ausgibt.

## Open-Closed Principle (OCP)
**Definition**: "A software artifact should be open for extension but closed for modifications"

### Problem
```go
package main

import (
    "fmt"
    "math"
)

type Square struct {
    side float64
}

type Circle struct {
    radius float64
}

type Calculator struct{}

func (c Calculator) SumAreas(shapes []interface{}) float64 {
    var total float64
    for _, shape := range shapes {
        switch s := shape.(type) {
        case Square:
            total += s.side * s.side
        case Circle:
            total += math.Pi * s.radius * s.radius
        }
    }
    return total
}

func main() {
    shapes := []interface{}{
        Square{side: 2},
        Circle{radius: 3},
    }
    calculator := Calculator{}
    fmt.Println("Total area:", calculator.SumAreas(shapes))
}
```
### Erklärung
In diesem Beispiel verletzt die `SumAreas`-Methode das OCP, da sie bei jeder neuen Form geändert werden muss.

### Lösung
```go
package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
}

type Square struct {
    side float64
}

func (s Square) Area() float64 {
    return s.side * s.side
}

type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}

type Calculator struct{}

func (c Calculator) SumAreas(shapes []Shape) float64 {
    var total float64
    for _, shape := range shapes {
        total += shape.Area()
    }
    return total
}

func main() {
    shapes := []Shape{
        Square{side: 2},
        Circle{radius: 3},
    }
    calculator := Calculator{}
    fmt.Println("Total area:", calculator.SumAreas(shapes))
}
```
### Erklärung
Hier haben wir eine `Shape`-Schnittstelle definiert, die eine `Area`-Methode enthält. Die `SumAreas`-Methode arbeitet nun mit dieser Schnittstelle, sodass sie nicht geändert werden muss, wenn neue Formen hinzugefügt werden.

### Übung
1. Fügen Sie eine neue Form `Triangle` hinzu und passen Sie die `SumAreas` Methode an, ohne den bestehenden Code zu ändern.

## Liskov Substitution Principle (LSP)
**Definition**: "If for each object o1 of type S there is an object o2 of type T such that for all programs P defined in terms of T, the behavior of P is unchanged when o1 is substituted for o2 then S is a subtype of T"

### Problem
```go
package main

import "fmt"

type Vehicle struct {
    name string
}

func (v Vehicle) GetName() string {
    return v.name
}

type Car struct {
    Vehicle
}

type Motorcycle struct {
    Vehicle
}

func PrintVehicleName(v Vehicle) {
    fmt.Println(v.GetName())
}

func main() {
    car := Car{Vehicle{name: "Car"}}
    motorcycle := Motorcycle{Vehicle{name: "Motorcycle"}}

    PrintVehicleName(car.Vehicle)
    PrintVehicleName(motorcycle.Vehicle)
}
```
### Erklärung
In diesem Beispiel verwenden wir Komposition anstelle von Vererbung, um das LSP zu erfüllen. `Car` und `Motorcycle` können durch `Vehicle` ersetzt werden, ohne dass das Verhalten des Programms geändert wird.

### Lösung
```go
package main

import "fmt"

type Vehicle interface {
    GetName() string
}

type Car struct {
    name string
}

func (c Car) GetName() string {
    return c.name
}

type Motorcycle struct {
    name string
}

func (m Motorcycle) GetName() string {
    return m.name
}

func PrintVehicleName(v Vehicle) {
    fmt.Println(v.GetName())
}

func main() {
    car := Car{name: "Car"}
    motorcycle := Motorcycle{name: "Motorcycle"}

    PrintVehicleName(car)
    PrintVehicleName(motorcycle)
}
```
### Erklärung
Hier haben wir eine `Vehicle`-Schnittstelle definiert, die eine `GetName`-Methode enthält. `Car` und `Motorcycle` implementieren diese Schnittstelle, sodass sie durch `Vehicle` ersetzt werden können.

### Übung
1. Implementieren Sie eine neue Fahrzeugart `Bicycle` und testen Sie die `PrintVehicleName` Funktion damit.

## Interface Segregation Principle (ISP)
**Definition**: "Clients should not be forced to depend on methods they don’t use"

### Problem
```go
package main

import "fmt"

type Shape interface {
    Area() float64
    Volume() float64
}

type Square struct {
    side float64
}

func (s Square) Area() float64 {
    return s.side * s.side
}

func (s Square) Volume() float64 {
    return 0 // Square has no volume
}

type Cube struct {
    side float64
}

func (c Cube) Area() float64 {
    return 6 * c.side * c.side
}

func (c Cube) Volume() float64 {
    return c.side * c.side * c.side
}

func SumAreas(shapes []Shape) float64 {
    var total float64
    for _, shape := range shapes {
        total += shape.Area()
    }
    return total
}

func SumVolumes(shapes []Shape) float64 {
    var total float64
    for _, shape := range shapes {
        total += shape.Volume()
    }
    return total
}

func main() {
    shapes := []Shape{
        Square{side: 2},
        Cube{side: 3},
    }

    fmt.Println("Total area:", SumAreas(shapes))
    fmt.Println("Total volume:", SumVolumes(shapes))
}
```
### Erklärung
In diesem Beispiel wird das ISP verletzt, da `Square` gezwungen ist, die `Volume`-Methode zu implementieren, obwohl es keine Volumen hat.

### Lösung
```go
package main

import "fmt"

type Shape interface {
    Area() float64
}

type Object interface {
    Shape
    Volume() float64
}

type Square struct {
    side float64
}

func (s Square) Area() float64 {
    return s.side * s.side
}

type Cube struct {
    side float64
}

func (c Cube) Area() float64 {
    return 6 * c.side * c.side
}

func (c Cube) Volume() float64 {
    return c.side * c.side * c.side
}

func SumAreas(shapes []Shape) float64 {
    var total float64
    for _, shape := range shapes {
        total += shape.Area()
    }
    return total
}

func SumVolumes(objects []Object) float64 {
    var total float64
    for _, object := range objects {
        total += object.Volume()
    }
    return total
}

func main() {
    shapes := []Shape{
        Square{side: 2},
        Cube{side: 3},
    }
    objects := []Object{
        Cube{side: 3},
    }

    fmt.Println("Total area:", SumAreas(shapes))
    fmt.Println("Total volume:", SumVolumes(objects))
}
```
### Erklärung
Hier haben wir die `Shape`-Schnittstelle in zwei Schnittstellen aufgeteilt: `Shape` für flache Formen und `Object` für Objekte mit Volumen. Dies erfüllt das ISP, da `Square` nicht gezwungen ist, die `Volume`-Methode zu implementieren.

### Übung
1. Implementieren Sie eine neue Form `Sphere` und passen Sie die bestehenden Funktionen an, ohne unnötige Methoden zu erzwingen.

## Dependency Inversion Principle (DIP)
**Definition**: "High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details. Details should not depend on abstractions"

### Problem
```go
package main

import "fmt"

type MySQL struct{}

func (db MySQL) Query(query string) string {
    return "MySQL result"
}

type Repository struct {
    db MySQL
}

func (r Repository) GetData() string {
    return r.db.Query("SELECT * FROM table")
}

func main() {
    repo := Repository{db: MySQL{}}
    fmt.Println(repo.GetData())
}
```
### Erklärung
In diesem Beispiel verletzt das Repository das DIP, da es direkt von der `MySQL`-Struktur abhängt.

### Lösung
```go
package main

import "fmt"

type Database interface {
    Query(query string) string
}

type MySQL struct{}

func (db MySQL) Query(query string) string {
    return "MySQL result"
}

type PostgreSQL struct{}

func (db PostgreSQL) Query(query string) string {
    return "PostgreSQL result"
}

type Repository struct {
    db Database
}

func (r Repository) GetData() string {
    return r.db.Query("SELECT * FROM table")
}

func main() {
    mysql := MySQL{}
    postgres := PostgreSQL{}

    repo1 := Repository{db: mysql}
    repo2 := Repository{db: postgres}

    fmt.Println(repo1.GetData())
    fmt.Println(repo2.GetData())
}
```
### Erklärung
Hier haben wir eine `Database`-Schnittstelle definiert, die eine `Query`-Methode enthält. Das Repository hängt nun von dieser Schnittstelle ab, sodass es nicht geändert werden muss, wenn eine neue Datenbank hinzugefügt wird.

### Übung
1. Implementieren Sie eine neue Datenbank `SQLite` und passen Sie das Repository an, um die neue Datenbank zu verwenden.

## Fazit
Als Softwareentwickler sollten wir immer nach der besten Möglichkeit suchen, unseren Code zu schreiben. Die Anwendung der SOLID-Prinzipien in unseren Projekten hilft uns, Code zu schreiben, der robuster, skalierbarer und toleranter gegenüber Änderungen ist, da wir die Grundlagen unserer Anwendung korrekt definieren.

## Quellen
https://blog.devgenius.io/golang-solid-principles-e7641dee90b0