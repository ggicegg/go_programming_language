package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	funcName()
}

func funcName() {
	start := time.Now()
	ch := make(chan string)
	fmt.Println(start, "开始执行")
	urls := []string{"http://www.baidu.com", "http://www.baidu.com", "http://www.bilibili.com", "http://www.google.cn"}
	for _, url := range urls {
		go fetchUrl(url, ch)
	}
	for i := len(urls); i > 0; i-- {
		s := <-ch
		fmt.Println(s)
	}
	end := time.Now()
	fmt.Println(end, "结束执行,执行了", time.Since(start).Seconds(), "秒")
}

func fetchUrl(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {

	}
	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)

	ch <- fmt.Sprintf("请求URL:%s完成,使用了:%f秒", url, time.Since(start).Seconds())
}
