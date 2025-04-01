package sync

import (
	"sync"
	"testing"
)

// IMPT LESSON -> Assignment of structs actually COPY ITS CONTENT!!!!
// thats why anything with a mutex should be passed by pointer / reference instead of val

// SYNC
// 1) sync.Mutex -> allows you to lock the state of a struct, other goroutines wil wait
// 2) sync.WaitGroup -> Add, Done, Wait will block and wait for wg to reach 0. After all goroutines Done.

func assertCounter(t testing.TB, got *Counter, want int){
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()
	
		var wg sync.WaitGroup
		wg.Add(wantedCount)
	
		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
	
		assertCounter(t, counter, wantedCount)
	})
}