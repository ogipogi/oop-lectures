package main

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(100 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("got %v want %v", got, want)
	}

}

func TestGoRoutine(t *testing.T) {
	urls := make([]string, 80)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	CheckWebsites(slowWebsiteChecker, urls)
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 20)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()
	for i := 0; i < 1; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}
