package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var incr int64
	gs := 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incr, 1)
			r := atomic.LoadInt64(&incr)
			fmt.Println(r)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(incr)
}
