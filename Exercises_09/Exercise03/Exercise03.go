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

	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)
			fmt.Println(runtime.NumGoroutine())
			wg.Done()
		}()

	}

	wg.Wait()
}
