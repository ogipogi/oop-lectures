package main

import (
	"fmt"
	"io"
)

type MessageFormatter struct {
	Prefix string
	Suffix string
}

func (mf *MessageFormatter) FormatMessage(message string) string {
	return mf.Prefix + message + mf.Suffix
}

func WriteFormattedMessages(writer io.Writer, messages []string, formatter *MessageFormatter) {
	for _, message := range messages {
		fmt.Fprintln(writer, formatter.FormatMessage(message))
	}
}
