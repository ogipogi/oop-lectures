package concurrency

import (
	"reflect"
	"testing"
	"time"
)

type SleeperMock struct {
	Calls int
}

func (s *SleeperMock) Sleep() {
	s.Calls++
}

func TestDummyFunction(t *testing.T) {
	got := DummyFunction(2)
	want := 4

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestConcurrentDummyFunction(t *testing.T) {
	numbers := []int{1, 2, 3}
	sleeper := SleeperMock{}
	got := ConcurrentDummyFunction(numbers, &sleeper)
	want := []int{2, 4, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestParallelExecutionWithWaitGroup(t *testing.T) {
	numbers := []int{1, 2, 3}
	start := time.Now()
	sleeper := SleeperMock{}
	ConcurrentDummyFunctionWithWaitGroup(numbers, &sleeper)
	duration := time.Since(start)

	if duration >= 3*time.Second {
		t.Errorf("expected function to run in less than 3 seconds, but it took %v", duration)
	}
}

func TestParallelExecution(t *testing.T) {
	numbers := []int{1, 2, 3}
	start := time.Now()
	sleeper := SleeperMock{}
	ConcurrentDummyFunction(numbers, &sleeper)
	duration := time.Since(start)

	if duration >= 3*time.Second {
		t.Errorf("expected function to run in less than 3 seconds, but it took %v", duration)
	}
}

func TestParallelExecutionWithMock(t *testing.T) {
	numbers := []int{1, 2, 3}
	sleeper := SleeperMock{}
	ConcurrentDummyFunction(numbers, &sleeper)
	if sleeper.Calls != 3 {
		t.Errorf("expected function to run 3 times, but it ran %d times", sleeper.Calls)
	}
}
