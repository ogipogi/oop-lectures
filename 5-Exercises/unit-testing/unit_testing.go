package main

const (
	helloPrefix   = "Hello, "
	morningPrefix = "Good morning, "
	eveningPrefix = "Good evening, "
)

func Hello(name string, greetingType string) string {
	if name == "" {
		name = "World"
	}
	prefix := helloPrefix
	if greetingType == "morning" {
		prefix = morningPrefix
	} else if greetingType == "evening" {
		prefix = eveningPrefix
	}
	return prefix + name
}
