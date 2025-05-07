package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {

	assertCounterValue := func(t *testing.T, counter Counter, expected int) {
		t.Helper()
		actual := counter.GetValue()
		if actual != expected {
			t.Errorf("got %d instead of the expected %d", actual, expected)
		}
	}

	t.Run("incrementing the counter", func(t *testing.T) {
		counter := NewCounter()
		counter.Increment()
		counter.Increment()
		counter.Increment()
		expected := 3
		assertCounterValue(t, counter, expected)
	})

	t.Run("concurrently incrementing the counter", func(t *testing.T) {
		expected := 1000
		var wg sync.WaitGroup
		wg.Add(expected)
		counter := NewCounter()
		for i := 0; i < expected; i++ {
			go func() {
				counter.Increment()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounterValue(t, counter, expected)
	})
}
