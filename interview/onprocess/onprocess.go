package main

import (
	"fmt"
	"runtime"
	"sync"
)

//因为GOMAXPROCS只有1，在主main goroutine执行完成后，才开始执行剩下的goroutine，此时第一个循环里的i早变成10了

func init() {
	fmt.Println("Current Go Version:", runtime.Version())
}
func main() {
	runtime.GOMAXPROCS(1)

	count := 10
	wg := sync.WaitGroup{}
	wg.Add(count * 2)
	for i := 0; i < count; i++ {
		go func() {
			fmt.Printf("[%d]", i)
			//time.Sleep(time.Millisecond*500)
			wg.Done()
		}()
	}
	for i := 0; i < count; i++ {
		go func(i int) {
			fmt.Printf("-%d-", i)
			wg.Done()
		}(i)
	}
	//time.Sleep(time.Second * 2)
	wg.Wait()
}
