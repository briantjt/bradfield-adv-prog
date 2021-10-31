package idservice

import (
	"sync"
	"testing"
)

const maxAttempts = 1000

func TestMutex(t *testing.T) {
	var wg sync.WaitGroup
	service := NewIdGenMutex()
	for i := 0; i < maxAttempts; i++ {
		wg.Add(1)
		go func() {
			service.GetNext()
			wg.Done()
		}()
	}

	wg.Wait()
	finalResult := service.GetNext()
	if finalResult != maxAttempts {
		t.Fatalf("Service returned %d, wanted %d", finalResult, maxAttempts+1)
	}
}

func TestAtomic(t *testing.T) {
	var wg sync.WaitGroup
	service := NewIdGenAtomic()
	for i := 0; i < maxAttempts; i++ {
		wg.Add(1)
		go func() {
			service.GetNext()
			wg.Done()
		}()
	}

	wg.Wait()
	finalResult := service.GetNext()
	if finalResult != maxAttempts {
		t.Fatalf("Service returned %d, wanted %d", finalResult, maxAttempts+1)
	}
}

func TestChannel(t *testing.T) {
	var wg sync.WaitGroup
	service := NewIdGenChan()
	for i := 0; i < maxAttempts; i++ {
		wg.Add(1)
		go func() {
			service.GetNext()
			wg.Done()
		}()
	}

	wg.Wait()
	finalResult := service.GetNext()
	if finalResult != maxAttempts {
		t.Fatalf("Service returned %d, wanted %d", finalResult, maxAttempts+1)
	}
}
