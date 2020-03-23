package main

import (
	"fmt"
	"go_programming_language/interview/testpackage"
	"time"
)

func main() {
	ch := make(chan bool)
	go work(ch)
	time.Sleep(time.Second * 1)
	ch <- true
	fmt.Println("shezhiwanle")
	close(ch)
	testpackage.DBF()
}

func work(ch chan bool) {
	for {
		select {
		default:
			fmt.Println("working...")
		case i, ok := <-ch:
			if !ok {
				fmt.Println(i)
				return
			}
			fmt.Println("tuichu")

		}
	}

}
