package main

import (
	"fmt"
	"io"
)

func WriteMessages(writer io.Writer, messages []string) {
	for _, message := range messages {
		fmt.Fprintln(writer, message)
	}
}
