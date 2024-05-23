package main

import (
	"fmt"
	"io"
	"os"
)

func GreetWithoutDependencyInjection(name string) {
	// without Dependency Injection
	fmt.Printf("Hello, %s", name)
}

func Greet(writer io.Writer, name string) {
	// with Dependency Injection
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	GreetWithoutDependencyInjection("Chris")

	Greet(os.Stdout, "Chris")
}
