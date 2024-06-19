### Goal 1: Refactor Project Structure for CLI and Webserver

**Short Goal:** Organize the project to accommodate both a web server and a command line application.

#### Step-by-Step Refactoring

1. **Create Directory Structure:**
   - Create a `cmd` directory with subdirectories `webserver` and `cli`.

    ```sh
    mkdir -p cmd/webserver
    mkdir -p cmd/cli
    ```

2. **Move `main.go` to `webserver`:**
   - Move the existing `main.go` to `cmd/webserver`.

    ```sh
    mv main.go cmd/webserver/main.go
    ```

3. **Update Package Names:**
   - Change the package name of all non-application files to `poker`.

    ```go
    // example: file_system_store.go
    package poker
    ```

4. **Update Imports in `main.go`:**
   - Adjust the import paths in `cmd/webserver/main.go` to reflect the new structure.

    ```go
    // cmd/webserver/main.go
    package main

    import (
        "github.com/your-username/your-repo-name/poker"
        "log"
        "net/http"
        "os"
    )

    const dbFileName = "game.db.json"

    func main() {
        db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
        if err != nil {
            log.Fatalf("problem opening %s %v", dbFileName, err)
        }

        store, err := poker.NewFileSystemPlayerStore(db)
        if err != nil {
            log.Fatalf("problem creating file system player store, %v", err)
        }

        server := poker.NewPlayerServer(store)
        log.Fatal(http.ListenAndServe(":5000", server))
    }
    ```

5. **Verify Project Structure:**
   - The project should now have a structure like this:

    ```
    .
    ├── cmd
    │   ├── cli
    │   │   └── main.go
    │   └── webserver
    │       └── main.go
    ├── file_system_store.go
    ├── file_system_store_test.go
    ├── league.go
    ├── server.go
    ├── server_integration_test.go
    ├── server_test.go
    ├── tape.go
    ├── tape_test.go
    └── testing.go
    ```

### Goal 2: Implement CLI for Recording Wins

**Short Goal:** Create a CLI that records a player's win from user input.

#### Test

```go
// CLI_test.go
package poker_test

import (
    "strings"
    "testing"

    "github.com/your-username/your-repo-name/poker"
)

func TestCLI(t *testing.T) {
    t.Run("record chris win from user input", func(t *testing.T) {
        in := strings.NewReader("Chris wins\n")
        playerStore := &poker.StubPlayerStore{}

        cli := poker.NewCLI(playerStore, in)
        cli.PlayPoker()

        poker.AssertPlayerWin(t, playerStore, "Chris")
    })

    t.Run("record cleo win from user input", func(t *testing.T) {
        in := strings.NewReader("Cleo wins\n")
        playerStore := &poker.StubPlayerStore{}

        cli := poker.NewCLI(playerStore, in)
        cli.PlayPoker()

        poker.AssertPlayerWin(t, playerStore, "Cleo")
    })
}
```
**Explanation:** This test verifies that the CLI can record wins for specific players based on user input.

#### Implementation

```go
// CLI.go
package poker

import (
    "bufio"
    "io"
    "strings"
)

type CLI struct {
    playerStore PlayerStore
    in          *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
    return &CLI{
        playerStore: store,
        in:          bufio.NewScanner(in),
    }
}

func (cli *CLI) PlayPoker() {
    userInput := cli.readLine()
    cli.playerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
    return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
    cli.in.Scan()
    return cli.in.Text()
}
```
**Explanation:** This implementation reads user input, extracts the player's name, and records the win.

### Goal 3: Main Function for CLI Application

**Short Goal:** Create the main function for the CLI application.

#### Implementation

```go
// cmd/cli/main.go
package main

import (
    "fmt"
    "github.com/your-username/your-repo-name/poker"
    "log"
    "os"
)

const dbFileName = "game.db.json"

func main() {
    store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
    if err != nil {
        log.Fatal(err)
    }
    defer close()

    fmt.Println("Let's play poker")
    fmt.Println("Type {Name} wins to record a win")
    poker.NewCLI(store, os.Stdin).PlayPoker()
}
```
**Explanation:** This main function initializes the CLI application, sets up the player store, and starts reading user input to record wins.

### Goal 4: Refactor File System Player Store

**Short Goal:** Create a helper function to initialize the player store from a file.

#### Implementation

```go
// file_system_store.go
package poker

import (
    "fmt"
    "os"
)

func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
    db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
    if err != nil {
        return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
    }

    closeFunc := func() {
        db.Close()
    }

    store, err := NewFileSystemPlayerStore(db)
    if err != nil {
        return nil, nil, fmt.Errorf("problem creating file system player store, %v", err)
    }

    return store, closeFunc, nil
}
```
**Explanation:** This helper function initializes the `FileSystemPlayerStore` from a given file path and returns a function to close the file.

### Final Steps: Update Webserver to Use New Helper

**Short Goal:** Refactor the webserver's main function to use the new helper function for initializing the player store.

#### Implementation

```go
// cmd/webserver/main.go
package main

import (
    "github.com/your-username/your-repo-name/poker"
    "log"
    "net/http"
)

const dbFileName = "game.db.json"

func main() {
    store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
    if err != nil {
        log.Fatal(err)
    }
    defer close()

    server := poker.NewPlayerServer(store)
    if err := http.ListenAndServe(":5000", server); err != nil {
        log.Fatalf("could not listen on port 5000 %v", err)
    }
}
```
**Explanation:** This refactoring makes the webserver's `main.go` cleaner and reuses the helper function to initialize the player store.

### Summary

- Refactored the project structure to accommodate both a web server and a CLI application.
- Created and tested a CLI for recording player wins.
- Implemented helper functions for initializing the player store.
- Updated the webserver to use the new helper function.

This setup now allows for a more modular and maintainable project, with shared code between the CLI and webserver applications.