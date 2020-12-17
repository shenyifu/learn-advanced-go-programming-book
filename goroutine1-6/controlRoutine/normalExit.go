package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(cancle chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("hello")
		case <-cancle:
			return
		}
	}
}

func main() {
	cancle := make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(cancle, &wg)
	}
	time.Sleep(time.Second)
	close(cancle)
	wg.Wait()

}
