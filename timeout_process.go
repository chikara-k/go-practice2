package main

import (
	"fmt"
	"time"
)

func main() {
	ch := generator("Hi!")
	for i := 0; i < 10; i++ {
		// Cases are evaluated simultaneously
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("Waited too long!")
			return
		}
	}
}

func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}
