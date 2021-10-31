package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("launched goroutine %d\n", i)
			wg.Done()
		}(i)
	}
	// Wait for goroutines to finish
	wg.Wait()
}
