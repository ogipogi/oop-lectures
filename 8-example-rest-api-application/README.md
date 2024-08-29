# Aufbau einer skalierbaren RESTFUL API mit Golang Gin Gonic Framework

## Voraussetzungen
- Installation der Go-Programmiersprache auf Ihrem Betriebssystem
- Grundkenntnisse in Go
- Vertrautheit mit RESTFUL API und JSON

## Was und warum Gin Gonic verwenden
Gin Gonic ist ein leichtgewichtiges HTTP-Framework und eines der bekanntesten Golang-Web-Frameworks. Gin ist ein leistungsstarker HTTP-Anforderungsrouter und behauptet, 40-mal schneller als Martini, ein ähnliches Framework in Go, zu sein. Weitere Details zu den Benchmarks finden Sie hier: https://github.com/gin-gonic/gin/blob/master/BENCHMARKS.md.

## Implementierung

### Projekt und Abhängigkeiten einrichten
Starten Sie mit der Initialisierung des Projekts mit Go Module und verwenden Sie es zur Verwaltung der Abhängigkeiten. Führen Sie diesen Befehl im GOPATH-Arbeitsbereich aus:
```sh
go mod init
```

Erstellen Sie eine Makefile für die Konfiguration, die Installation von Abhängigkeiten und das Erstellen von Verzeichnissen für das Projektsetup:
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

Führen Sie dann diesen Befehl aus, um die Abhängigkeiten herunterzuladen und zur go.mod hinzuzufügen:
```sh
make init-dependency
```

Die Ordnerstruktur für dieses RESTFUL API-Projekt sieht dann so aus:
```
.
├── Makefile
├── go.mod
├── go.sum
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

### .env Datei Konfiguration
```env
PORT=8080
# Application
APPLICATION_NAME=simple-restful-api

# Database
DB_DSN="host=localhost user=root password=root dbname=gin-gonic-api port=5432"

# Logging
LOG_LEVEL=DEBUG
```

### Logger mit Golang Logrus initialisieren
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

### Verbindung zur Datenbank einrichten
```go
package config

import (
 "gorm.io/driver/postgres"
 "gorm.io/gorm"
 "log"
 "os"
)

func ConnectToDB() *gorm.DB {
 var err error
 dsn := os.Getenv("DB_DSN")

 db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
 if err != nil {
  log.Fatal("Error connecting to database. Error: ", err)
 }

 return db
}
```

### Konstanten definieren
```go
package constant

type ResponseStatus int
type Headers int
type General int

// Constant Api
const (
   Success ResponseStatus = iota + 1
   DataNotFound
   UnknownError
   InvalidRequest
   Unauthorized
)

func (r ResponseStatus) GetResponseStatus() string {
   return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED"}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
   return [...]string{"Success", "Data Not Found", "Unknown Error", "Invalid Request", "Unauthorized"}[r-1]
}
```

### Model oder DAO mit GORM und DTO definieren
```go
package dao

import (
 "gorm.io/gorm"
 "time"
)

type BaseModel struct {
 CreatedAt time.Time      `gorm:"->:false;column:created_at" json:"-"`
 UpdatedAt time.Time      `gorm:"->:false;column:updated_at" json:"-"`
 DeletedAt gorm.DeletedAt `gorm:"->:false;column:deleted_at" json:"-"`
}

type Role struct {
 ID   int    `gorm:"column:id; primary_key; not null" json:"id"`
 Role string `gorm:"column:role" json:"role"`
 BaseModel
}

type User struct {
 ID       int    `gorm:"column:id; primary_key; not null" json:"id"`
 Name     string `gorm:"column:name" json:"name"`
 Email    string `gorm:"column:email" json:"email"`
 Password string `gorm:"column:password;->:false" json:"-"`
 Status   int    `gorm:"column:status" json:"status"`
 RoleID   int    `gorm:"column:role_id;not null" json:"role_id"`
 Role     Role   `gorm:"foreignKey:RoleID;references:ID" json:"role"`
 BaseModel
}
```

### DTO definieren
```go
package dto

type ApiResponse[T any] struct {
 ResponseKey     string `json:"response_key"`
 ResponseMessage string `json:"response_message"`
 Data            T      `json:"data"`
}
```

### Util für API-Antworten erstellen
```go
package pkg

import (
 "gin-gonic-api/app/constant"
 "gin-gonic-api/app/domain/dto"
)

func Null() interface{} {
 return nil
}

func BuildResponse[T any](responseStatus constant.ResponseStatus, data T) dto.ApiResponse[T] {
 return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func BuildResponse_[T any](status string, message string, data T) dto.ApiResponse[T] {
 return dto.ApiResponse[T]{
  ResponseKey:     status,
  ResponseMessage: message,
  Data:            data,
 }
}
```

### Benutzerdefinierte Fehler und Fehlerbehandlung hinzufügen
```go
package pkg

import (
 "errors"
 "fmt"
 "gin-gonic-api/app/constant"
)

func PanicException_(key string, message string) {
 err := errors.New(message)
 err = fmt.Errorf("%s: %w", key, err)
 if err != nil {
  panic(err)
 }
}

func PanicException(responseKey constant.ResponseStatus) {
 PanicException_(responseKey.GetResponseStatus(), responseKey.GetResponseMessage())
}

func PanicHandler(c *gin.Context) {
 if err := recover(); err != nil {
  str := fmt.Sprint(err)
  strArr := strings.Split(str, ":")

  key := strArr[0]
  msg := strings.Trim(strArr[1], " ")

  switch key {
  case
   constant.DataNotFound.GetResponseStatus():
   c.JSON(http.StatusBadRequest, BuildResponse_(key, msg, Null()))
   c.Abort()
  case
   constant.Unauthorized.GetResponseStatus():
   c.JSON(http.StatusUnauthorized, BuildResponse_(key, msg, Null()))
   c.Abort()
  default:
   c.JSON(http.StatusInternalServerError, BuildResponse_(key, msg, Null()))
   c.Abort()
  }
 }
}
```

### Route, Controller, Service und Repository implementieren
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

### Controller definieren
```go
package controller

import (
 "gin-gonic-api/app/service"
 "github.com/gin-gonic/gin"
)

type UserController interface {
 GetAllUserData(c *gin.Context)
 AddUserData(c *gin.Context)
 GetUserById(c *gin.Context)
 UpdateUserData(c *gin.Context)
 DeleteUser(c *gin.Context)
}

type UserControllerImpl struct {
 svc service.UserService
}

func (u UserControllerImpl) GetAllUserData(c *gin.Context) {
 u.svc.GetAllUser(c)
}

func (u UserControllerImpl) AddUserData(c *gin.Context) {
 u.svc.AddUserData(c)
}

func (u UserControllerImpl) GetUserById(c *gin.Context) {
 u.svc.GetUserById(c)
}

func (u UserControllerImpl) UpdateUserData(c *gin.Context) {
 u.svc.UpdateUserData(c)
}

func (u UserControllerImpl) DeleteUser(c *gin.Context) {
 u.svc.DeleteUser(c)
}

func UserControllerInit(userService service.UserService) *UserControllerImpl {
 return &UserControllerImpl{
  svc: userService,
 }
}
```

### Service definieren
```go
package service

import (
   "gin-gonic-api/app/constant"
   "gin-gonic-api/app/domain/dao"
   "gin-gonic-api/app/pkg"
   "gin-gonic-api/app/repository"
   "github.com/gin-gonic/gin"
   log "github.com/sirupsen/logrus"
   "golang.org/x/crypto/bcrypt"
   "net/http"
   "strconv"
)

type UserService interface {
   GetAllUser(c *gin.Context)
   GetUserById(c *gin.Context)
   AddUserData(c *gin.Context)
   UpdateUserData(c *gin.Context)
   DeleteUser(c *gin.Context)
}

type UserServiceImpl struct {
   userRepository repository.UserRepository
}

func (u UserServiceImpl) UpdateUserData(c *gin.Context) {
   defer pkg.PanicHandler(c)

   log.Info("start to execute program update user data by id")
   userID, _ := strconv.Atoi(c.Param("userID"))

   var request dao.User
   if err := c.ShouldBindJSON(&request); err != nil {
      log.Error("Happened error when mapping request from FE. Error", err)
      pkg.PanicException(constant.InvalidRequest)
   }

   data, err := u.userRepository.FindUserById(userID)
   if err != nil {
      log.Error("Happened error when get data from database. Error", err)
      pkg.PanicException(constant.DataNotFound)
   }

   data.RoleID = request.RoleID
   data.Email = request.Email
   data.Name = request.Password
   data.Status = request.Status
   u.userRepository.Save(&data)

   if err != nil {
      log.Error("Happened error when updating data to database. Error", err)
      pkg.PanicException(constant.UnknownError)
   }

   c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) GetUserById(c *gin.Context) {
   defer pkg.PanicHandler(c)

   log.Info("start to execute program get user by id")
   userID, _ := strconv.Atoi(c.Param("userID"))

   data, err := u.userRepository.FindUserById(userID)
   if err != nil {
      log.Error("Happened error when get data from database. Error", err)
      pkg.PanicException(constant.DataNotFound)
   }

   c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) AddUserData(c *gin.Context) {
   defer pkg.PanicHandler(c)

   log.Info("start to execute program add data user")
   var request dao.User
   if err := c.ShouldBindJSON(&request); err != nil {
      log.Error("Happened error when mapping request from FE. Error", err)
      pkg.PanicException(constant.InvalidRequest)
   }

   hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
   request.Password = string(hash)

   data, err := u.userRepository.Save(&request)
   if err != nil {
      log.Error("Happened error when saving data to database. Error", err)
      pkg.PanicException(constant.UnknownError)
   }

   c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) GetAllUser(c *gin.Context) {
   defer pkg.PanicHandler(c)

   log.Info("start to execute get all data user")

   data, err := u.userRepository.FindAllUser()
   if err != nil {
      log.Error("Happened Error when find all user data. Error: ", err)
      pkg.PanicException(constant.UnknownError)
   }

   c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) DeleteUser(c *gin.Context) {
   defer pkg.PanicHandler(c)

   log.Info("start to execute delete data user by id")
   userID, _ := strconv.Atoi(c.Param("userID"))

   err := u.userRepository.DeleteUserById(userID)
   if err != nil {
      log.Error("Happened Error when try delete data user from DB. Error:", err)
      pkg.PanicException(constant.UnknownError)
   }

   c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func UserServiceInit(userRepository repository.UserRepository) *UserServiceImpl {
   return &UserServiceImpl{
      userRepository: userRepository,
   }
}
```

### Repository definieren
```go
package repository

import (
 "gin-gonic-api/app/domain/dao"
 log "github.com/sirupsen/logrus"
 "gorm.io/gorm"
)

type UserRepository interface {
 FindAllUser() ([]dao.User, error)
 FindUserById(id int) (dao.User, error)
 Save(user *dao.User) (dao.User, error)
 DeleteUserById(id int) error
}

type UserRepositoryImpl struct {
 db *gorm.DB
}

func (u UserRepositoryImpl) FindAllUser() ([]dao.User, error) {
 var users []dao.User

 var err = u.db.Preload("Role").Find(&users).Error
 if err != nil {
  log.Error("Got an error finding all couples. Error: ", err)
  return nil, err
 }

 return users, nil
}

func (u UserRepositoryImpl) FindUserById(id int) (dao.User, error) {
 user := dao.User{
  ID: id,
 }
 err := u.db.Preload("Role").First(&user).Error
 if err != nil {
  log.Error("Got and error when find user by id. Error: ", err)
  return dao.User{}, err
 }
 return user, nil
}

func (u UserRepositoryImpl) Save(user *dao.User) (dao.User, error) {
 var err = u.db.Save(user).Error
 if err != nil {
  log.Error("Got an error when save user. Error: ", err)
  return dao.User{}, err
 }
 return *user, nil
}

func (u UserRepositoryImpl) DeleteUserById(id int) error {
 err := u.db.Delete(&dao.User{}, id).Error
 if err != nil {
  log.Error("Got an error when delete user. Error: ", err)
  return err
 }
 return nil
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
 db.AutoMigrate(&dao.User{})
 return &UserRepositoryImpl{
  db: db,
 }
}
```

### Initialisierung der Komponenten mit Google Wire
```go
package config

import (
   "gin-gonic-api/app/controller"
   "gin-gonic-api/app/repository"
   "gin-gonic-api/app/service"
)

type Initialization struct {
   userRepo repository.UserRepository
   userSvc  service.UserService
   UserCtrl controller.UserController
   RoleRepo repository.RoleRepository
}

func NewInitialization(userRepo repository.UserRepository,
   userService service.UserService,
   userCtrl controller.UserController,
   roleRepo repository.RoleRepository) *Initialization {
   return &Initialization{
      userRepo: userRepo,
      userSvc:  userService,
      UserCtrl: userCtrl,
      RoleRepo: roleRepo,
   }
}
```

### Injector definieren
```go
// go:build wireinject
// +build wireinject

package config

import (
 "gin-gonic-api/app/controller"
 "gin-gonic-api/app/repository"
 "gin-gonic-api/app/service"
 "github.com/google/wire"
)

var db = wire.NewSet(ConnectToDB)

var userServiceSet = wire.NewSet(service.UserServiceInit,
 wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
 wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var userCtrlSet = wire.NewSet(controller.UserControllerInit,
 wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

var roleRepoSet = wire.NewSet(repository.RoleRepositoryInit,
 wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepositoryImpl)),
)

func Init() *Initialization {
 wire.Build(NewInitialization, db, userCtrlSet, userServiceSet, userRepoSet, roleRepoSet)
 return nil
}
```

### Main-Datei
```markdown
package main

import (
 "gin-gonic-api/config"
 "gin-gonic-api/router"
 "github.com/joho/godotenv"
 "os"
)

func init() {
 godotenv.Load()
 config.InitLog()
}

func main() {
 port := os.Getenv("PORT")

 init := config.Init()
 app := router.Init(init)

 app.Run(":" + port)
}
```

## Testen der API

Hinweis: Initialisieren Sie die Rollendaten in der Datenbank, da die Benutzerdaten zur Rollentabelle gehören.

Führen Sie diesen SQL-Befehl aus:
```sql
INSERT INTO roles(id, "role", created_at, updated_at, deleted_at)
VALUES
(1, 'ADMIN', current_timestamp, null, null),
(2, 'USER', current_timestamp, null, null);
```

### Testen der Endpunkte

#### Alle Benutzerdaten abrufen
```sh
curl --location --request GET 'http://localhost:8080/api/user'
```

Die Antwort sollte so aussehen:
```json
{
    "response_key": "SUCCESS",
    "response_message": "Success",
    "data": [
        {
            "id": 2,
            "name": "wayne",
            "email": "wayne@mail.ic",
            "status": 0,
            "role": {
                "id": 1,
                "role": "ADMIN"
            }
        },
        {
            "id": 3,
            "name": "wayne",
            "email": "wayne@mail.ic",
            "status": 0,
            "role": {
                "id": 1,
                "role": "ADMIN"
            }
        }
    ]
}
```

#### Benutzerdaten erstellen
```sh
curl --location --request POST 'http://localhost:8080/api/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "wayne",
    "email": "wayne@mail.id",
    "password": "plain_password",
    "role_id": 1
}'
```

Die Antwort sollte so aussehen:
```json
{
    "response_key": "SUCCESS",
    "response_message": "Success",
    "data": {
        "id": 8,
        "name": "wayne",
        "email": "wayne@mail.id",
        "status": 0,
        "role_id": 1
    }
}
```

Die restlichen Endpunkte wie Aktualisieren und Löschen können Sie selbst testen.

## Quellen
https://medium.com/@wahyubagus1910/build-scalable-restful-api-with-golang-gin-gonic-framework-43793c730d10 
```