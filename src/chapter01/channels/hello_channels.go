package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(ch <-chan int) {
	for i := range ch {
		fmt.Printf("received %d\n", i)
	}
	wg.Done()
}

func main()  {
	ch := make(chan int)
	wg.Add(1)
	go printer(ch)
	for i:=0;i<10;i++ {
		ch <- i
	}
	close(ch) // if not close, will cause dead lock
	wg.Wait()
}
