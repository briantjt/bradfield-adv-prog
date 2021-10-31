package main

import (
	"fmt"
)

const numTasks = 3

func main() {
	done := make(chan struct{}, numTasks)
	for i := 0; i < numTasks; i++ {
		go func(out chan <- struct{}) {
			fmt.Println("running task...")

			// Signal that task is done
			done <- struct{}{}
		}(done)
	}

	// Wait for tasks to complete
	for i := 0; i < numTasks; i++ {
		<-done
	}
	fmt.Printf("all %d tasks done!\n", numTasks)
}
