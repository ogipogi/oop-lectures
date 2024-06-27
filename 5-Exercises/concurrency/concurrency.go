package concurrency

import (
	"sync"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func newSleeper() Sleeper {
	return &DefaultSleeper{}
}

func DummyFunction(n int) int {
	return n * 2
}

func ConcurrentDummyFunctionWithWaitGroup(numbers []int, sleeper Sleeper) []int {
	results := make([]int, len(numbers))
	var wg sync.WaitGroup
	for i, n := range numbers {
		wg.Add(1)
		go func(i, n int) {
			defer wg.Done()
			sleeper.Sleep()
			results[i] = DummyFunction(n)
		}(i, n)
	}
	wg.Wait()
	return results
}

/*
type keyValue struct {
	index int
	value int
}
*/

func ConcurrentDummyFunction(numbers []int, sleeper Sleeper) []int {
	results := make([]int, len(numbers))
	resultChannel := make(chan struct {
		index int
		value int
	})
	for i, n := range numbers {
		go func(i, n int) {
			result := DummyFunction(n)
			sleeper.Sleep()
			resultChannel <- struct {
				index int
				value int
			}{i, result}

			//resultChannel <- keyValue{i, result}
		}(i, n)
	}
	for range numbers {
		result := <-resultChannel
		results[result.index] = result.value
	}
	return results
}
