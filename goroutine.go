package main

import (
	"fmt"
	"sync"
)

func doSomething(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		fmt.Printf("%d\n", id)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go doSomething(&wg, i)
	}
	wg.Wait()
}
