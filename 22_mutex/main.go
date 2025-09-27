package main

import (
	"fmt"
	"sync"
)

// post represents a blog post with a view counter protected by a mutex.
// The mutex ensures safe concurrent access to the views field.
type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	p.mu.Lock()
	defer p.mu.Unlock()
	p.views++
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("Mutex!")
	myPost := post{
		views: 0,
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	fmt.Println(myPost.views)
}
