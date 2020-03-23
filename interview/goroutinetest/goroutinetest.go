package main

import (
	"fmt"
	"time"
)

func main() {
	var c1 = make(chan int, 1)
	go func() {
		fmt.Println("hello")
		close(c1)
		time.Sleep(time.Second * 3)

		i := <-c1
		//time.Sleep(time.Second*3)
		fmt.Println(i)
	}()
	c1 <- 1
	time.Sleep(time.Second * 4)
	fmt.Println("end")
}
