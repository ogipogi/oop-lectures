# Aufbau einer skalierbaren RESTFUL API mit Golang Gin Gonic Framework

## Voraussetzungen
- Installation der Go-Programmiersprache auf Ihrem Betriebssystem
- Grundkenntnisse in Go
- Vertrautheit mit RESTFUL API und JSON
- Docker und Docker Compose installiert

## Was und warum Gin Gonic verwenden?
Gin Gonic ist ein leichtgewichtiges HTTP-Framework und eines der bekanntesten Golang-Web-Frameworks. Es bietet eine hohe Leistung und Effizienz beim Routing von HTTP-Anfragen und behauptet, 40-mal schneller als Martini (ein ähnliches Framework in Go) zu sein. Weitere Details zu den Benchmarks finden Sie hier: [Gin Gonic Benchmarks](https://github.com/gin-gonic/gin/blob/master/BENCHMARKS.md).

## Implementierung

### 1. Projekt und Abhängigkeiten einrichten

Beginnen Sie mit der Initialisierung des Projekts mit Go Module, um Abhängigkeiten zu verwalten. 

```bash
go mod init <Projektname>
```

#### Makefile für Abhängigkeiten:
Erstellen Sie eine Makefile, um Abhängigkeiten zu installieren und die Projektstruktur einzurichten.

```makefile
init-dependency:
 go get -u github.com/antonfisher/nested-logrus-formatter
 go get -u github.com/gin-gonic/gin
 go get -u golang.org/x/crypto
 go get -u gorm.io/gorm
 go get -u gorm.io/driver/postgres
 go get -u github.com/sirupsen/logrus
 go get -u github.com/joho/godotenv
```

Führen Sie anschließend diesen Befehl aus, um die Abhängigkeiten herunterzuladen:
```bash
make init-dependency
```

#### Ordnerstruktur für das Projekt:
Nach der Einrichtung sieht die Struktur wie folgt aus:
```
.
├── Makefile
├── go.mod
├── main.go
├── config
│   └── config.go
├── controller
│   └── user_controller.go
├── dao
│   └── user.go
├── dto
│   └── api_response.go
├── pkg
│   └── response.go
├── repository
│   └── user_repository.go
├── router
│   └── router.go
├── service
│   └── user_service.go
└── .env
```

### **Exercise 1: Projekt Setup**
- Initialisieren Sie das Projekt mit `go mod init`.
- Erstellen Sie die Verzeichnisstruktur manuell und überprüfen Sie die Installation der Abhängigkeiten.
- Erstellen Sie eine `main.go`-Datei, die "Hello, Gin!" ausgibt, um sicherzustellen, dass alles korrekt eingerichtet ist.

---

### 2. .env Datei Konfiguration
Erstellen Sie eine `.env` Datei für Umgebungsvariablen wie Port, Datenbank-Verbindung und Logging-Ebene.

```env
PORT=8080
APPLICATION_NAME=simple-restful-api
DB_DSN="host=db user=root password=root dbname=gin-gonic-api port=5432"
LOG_LEVEL=DEBUG
```

> **Hinweis:** Beachten Sie, dass hier `host=db` steht. Dies ist wichtig, weil der Datenbank-Hostname in Docker Compose als `db` definiert wird.

### 3. Logger mit Logrus initialisieren
Es ist wichtig, ein Logging-System zu haben, um Fehler und Events in Ihrer Anwendung zu verfolgen. Logrus ist eine leistungsstarke Logging-Bibliothek für Go.

```go
package config

import (
 nested "github.com/antonfisher/nested-logrus-formatter"
 "github.com/joho/godotenv"
 log "github.com/sirupsen/logrus"
 "os"
)

func InitLog() {
 log.SetLevel(getLoggerLevel(os.Getenv("LOG_LEVEL")))
 log.SetReportCaller(true)
 log.SetFormatter(&nested.Formatter{
  HideKeys:        true,
  FieldsOrder:     []string{"component", "category"},
  TimestampFormat: "2006-01-02 15:04:05",
  ShowFullLevel:   true,
  CallerFirst:     true,
 })
}

func getLoggerLevel(value string) log.Level {
 switch value {
 case "DEBUG":
  return log.DebugLevel
 case "TRACE":
  return log.TraceLevel
 default:
  return log.InfoLevel
 }
}
```

### **Exercise 2: Logrus erkunden**
- Fügen Sie `log.Info` oder `log.Error` an verschiedenen Stellen in Ihrem Code hinzu, um zu sehen, wie das Loggingsystem funktioniert.
- Experimentieren Sie mit verschiedenen Log-Ebenen (DEBUG, INFO, etc.) in Ihrer `.env` Datei.

---

### 4. Verbindung zur Datenbank einrichten (PostgreSQL)
Für die Speicherung von Daten verwenden wir PostgreSQL in Verbindung mit GORM (ORM für Go).

```go
package config

import (
 "gorm.io/driver/postgres"
 "gorm.io/gorm"
 "log"
 "os"
)

func ConnectToDB() *gorm.DB {
 dsn := os.Getenv("DB_DSN")
 db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
 if err != nil {
  log.Fatal("Error connecting to database. Error: ", err)
 }
 return db
}
```

### **Exercise 3: Datenbankverbindung testen**
- Stellen Sie sicher, dass PostgreSQL auf Ihrem lokalen Rechner läuft oder in einem Container (siehe nächster Abschnitt).
- Fügen Sie Testdaten in Ihre PostgreSQL-Datenbank ein und testen Sie die Verbindung.

---

### 5. Dockerfile erstellen
Um die Anwendung zu containerisieren, erstellen Sie eine `Dockerfile`.

```dockerfile
# Dockerfile

# Stage 1: Build the Go application
FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /simple-restful-api

# Stage 2: A minimal Docker image to run the Go application
FROM alpine:latest

WORKDIR /root/

COPY --from=build /simple-restful-api .

COPY .env ./

CMD ["./simple-restful-api"]
```

---

### 6. Docker Compose Datei erstellen
Jetzt konfigurieren wir Docker Compose, um sowohl die API als auch die PostgreSQL-Datenbank zu starten.

```yaml
# docker-compose.yml

version: '3.7'

services:
  api:
    build: .
    ports:
      - "8080:8080"
    env_file:
      - .env
    depends_on:
      - db
    networks:
      - api-network

  db:
    image: postgres:13
    environment:
      POSTGRES_USER: root
      POSTGRES_PASSWORD: root
      POSTGRES_DB: gin-gonic-api
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - api-network

networks:
  api-network:

volumes:
  postgres_data:
```

### **Exercise 4: Containerisierung testen**
1. Erstellen Sie die Docker-Container:
   ```bash
   docker-compose up --build
   ```
2. Überprüfen Sie, ob die API unter `localhost:8080` läuft und die Verbindung zur Datenbank funktioniert.

---

### 7. Routing und Controller einrichten
Wir verwenden Gin Gonic für das Routing und setzen Controller ein, um die API-Anfragen zu verarbeiten.

```go
package router

import (
 "gin-gonic-api/config"
 "github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {
 router := gin.New()
 router.Use(gin.Logger())
 router.Use(gin.Recovery())

 api := router.Group("/api")
 {
  user := api.Group("/user")
  user.GET("", init.UserCtrl.GetAllUserData)
  user.POST("", init.UserCtrl.AddUserData)
  user.GET("/:userID", init.UserCtrl.GetUserById)
  user.PUT("/:userID", init.UserCtrl.UpdateUserData)
  user.DELETE("/:userID", init.UserCtrl.DeleteUser)
 }
 return router
}
```

### **Exercise 5: API-Endpunkte testen**
- Verwenden Sie `curl` oder Postman, um die verschiedenen Endpunkte zu testen. Zum Beispiel:
  ```bash
  curl --location --request GET 'http://localhost:8080/api/user'
  ```
- Testen Sie alle CRUD-Endpunkte: Benutzer erstellen, lesen, aktualisieren und löschen.

---

### 8. Abschluss-Tests und Rollendaten initialisieren

Bevor Sie das API vollständig testen, initialisieren Sie Rollendaten in der Datenbank, da die Benutzerdaten mit der Rollentabelle verknüpft sind.

```sql
INSERT INTO roles(id, "role", created_at, updated_at, deleted_at)
VALUES
(1, 'ADMIN', current_timestamp, null, null),
(2, 'USER', current_timestamp, null, null);
```

---

## Zusammenfassung
Mit dieser Implementierung haben wir nicht nur eine RESTful API mit Gin Gonic und PostgreSQL erstellt, sondern die Anwendung auch containerisiert und mit Docker Compose orchestriert. Dies ermöglicht es, die Anwendung einfach zu deployen und die Abhängigkeiten in isolierten Umgebungen zu verwalten.

---

### Quellen:
- [Build Scalable RESTFUL API with Golang Gin Gonic Framework](https://medium.com/@wahyubagus1910/build-scalable-restful-api-with-golang-gin-gonic-framework-43793c730d10)