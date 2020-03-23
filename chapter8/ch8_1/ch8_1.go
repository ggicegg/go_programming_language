package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, request *http.Request) {
		form := request.MultipartForm
		headers := form.File
	})
	for i := 1; i < 100; {
		i = i + rand.Intn(5)
		progress_bar(i)
	}
}

func testFib() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Println(fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, c := range `-\|/` {
			fmt.Printf("\r%c", c)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)

}

func progress_bar(progress int) {
	if progress < 0 {
		return
	}
	i := 1
	for i := 1; i <= progress; i++ {
		fmt.Printf("#")
	}
	fmt.Printf("完成%d%%", progress)
	time.Sleep(time.Second)
	for ; i > 0; i-- {
		fmt.Printf("\r")
	}
}
