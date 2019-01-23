package main

import (
	"fmt"
	"sync"
)

func countDown(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println(i)
	}
	wg.Done()
}

func main() {
	//prepare concurrency variables
	ch := make(chan int)
	var wg sync.WaitGroup

	//start count down
	wg.Add(1)
	go countDown(ch,&wg)

	//count down
	for i := 10; i>0; i-- {
		ch <- i
	}
	close(ch)
	wg.Wait()

	//salute
	fmt.Println("Hello World")
}
