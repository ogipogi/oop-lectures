package main

import (
	"fmt"
	"time"
)

func doSomething(x int) {
	fmt.Println("Doing something with", x)
}

func doSomething2(y func(int) string) {
	fmt.Println("start....")
	time.Sleep(5 * time.Second)
	fmt.Println(y(42))
	fmt.Println("end....")
}

func main2() {
	x := 42
	doSomething(x)

	f := func(x int) string {
		return "my number is: " + fmt.Sprint(x)
	}
	doSomething2(f)
}
