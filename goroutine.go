package main

import (
	"sync"
	"time"
)

func doSomething(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Second)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go doSomething(&wg)
	}
	wg.Wait()
}
