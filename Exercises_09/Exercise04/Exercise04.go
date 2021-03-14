package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	total := 55
	counter := 0

	var mu sync.Mutex
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			mu.Lock()
			v := counter
			//runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)
			fmt.Println(runtime.NumGoroutine())
			mu.Unlock()
			wg.Done()
		}()

	}

	wg.Wait()
}
