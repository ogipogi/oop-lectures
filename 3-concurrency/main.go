package main

import (
	"fmt"
	"time"
)

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)

	start := time.Now()
	for _, url := range urls {
		go func(u string) {
			start := time.Now()
			resultsChannel <- result{u, wc(u)}
			//results[u] = wc(u)
			fmt.Println("go routine time: ", time.Since(start))
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultsChannel
		results[r.string] = r.bool
	}

	stop := time.Since(start)
	fmt.Println("total time: ", stop)

	time.Sleep(5 * time.Second)
	return results
}
