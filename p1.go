package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)


func runP1() {
	var receiveCount uint64
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= 11; i++ {

			c <- i
			fmt.Println(" g1 Sent: ", i)
		}
		close(c)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for r:= range c {
			fmt.Println(" g2 Received: ", r)
			atomic.AddUint64(&receiveCount, 1)
		}

	}()

	wg.Wait()

	fmt.Println("Total Received: ", receiveCount)






}
