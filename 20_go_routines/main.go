package main

import (
	"fmt"
	"sync"
	// "time"
)

// task simulates a concurrent task.
// It prints the task id and signals completion to the WaitGroup.
func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this task as done when function exits.
	fmt.Println("Running task is :", id)
}

func main() {
	var wg sync.WaitGroup // WaitGroup to wait for all goroutines to finish.

	// Launch 11 goroutines, each running task(i).
	for i := 0; i <= 10; i++ {
		wg.Add(1) // Increment WaitGroup counter.
		go task(i, &wg) // Start task as a goroutine.
	}

	// Wait for all goroutines to finish.
	// This is the correct way to wait instead of using Sleep.
	wg.Wait()
	// time.Sleep(1 * time.Second)
}