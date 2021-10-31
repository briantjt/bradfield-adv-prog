package idservice

import (
	"sync"
	"testing"
)

const maxAttempts = 100000

func TestIdGeneration(t *testing.T) {
	for _, test := range []struct {
		name    string
		service IdService
	}{
		{
			name: "Mutex", service: NewIdGenMutex(),
		},
		{
			name: "Atomic", service: NewIdGenAtomic(),
		},
		{
			name: "Channel", service: NewIdGenChan(),
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			var wg sync.WaitGroup
			for i := 0; i < maxAttempts; i++ {
				wg.Add(1)
				go func() {
					test.service.GetNext()
					wg.Done()
				}()
			}

			wg.Wait()
			finalResult := test.service.GetNext()
			if finalResult != maxAttempts {
				t.Fatalf("Service returned %d, wanted %d", finalResult, maxAttempts)
			}
		})
	}
}

func BenchmarkIdGeneration(b *testing.B) {
	for _, bench := range []struct {
		name    string
		service IdService
	}{
		{
			name: "Mutex", service: NewIdGenMutex(),
		},
		{
			name: "Atomic", service: NewIdGenAtomic(),
		},
		{
			name: "Channel", service: NewIdGenChan(),
		},
	} {
		b.Run(bench.name, func(b *testing.B) {
			var wg sync.WaitGroup
			for i := 0; i < b.N; i++ {
				wg.Add(1)
				go func() {
					bench.service.GetNext()
					wg.Done()
				}()
			}

			wg.Wait()
		})
	}
}
