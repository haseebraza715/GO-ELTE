package main

import (
	"fmt"
	"sync"
)

func runP2() {

	ints := []int{1,2,-3,-4,5}
	var intMap = make(map[int]string)
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, i := range ints {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mutex.Lock()
			intMap[n] = fmt.Sprint(handleInt(n))
			mutex.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(intMap)


}

func handleInt( n int) int {
	return ( n +6) * 3;
}