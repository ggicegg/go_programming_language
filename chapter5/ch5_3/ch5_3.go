package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"regexp"
)

func main() {

	//findLinks("https://www.baidu.com")
	//log.Fatal("hello")
	//findLinks("https://apps.apple.com/cn/genre/ios/id36")
	//i := new(map[string]string)
	//i2 := new([]int)
	//
	//fmt.Println(i == nil)
	//fmt.Println(*i == nil)
	//fmt.Println(i2 == nil)
	//fmt.Println(*i2 == nil)
	//i["hello"] = "hello"
	//fmt.Printf
	var r = regexp.Regexp{}

}

func sayHello(i i1) {
	i.Demo1()
	//i.Demo2()
}

type i1 interface {
	Demo1()
	Demo1() (s string)
	Demo2()
}
type user struct {
	name string
	age  int
}

func (u user) Demo1() {
	fmt.Println(u)
}

func findLinks(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("请求失败")
	}

	var buf bytes.Buffer
	buf.ReadFrom(resp.Body)
	fmt.Printf(string(buf.Bytes()))
	node, err := html.Parse(resp.Body)
	if err != nil {

	}

	fmt.Printf(node.Data)
}

/**
测试函数返回多个值
*/
func returntest() (str1 string, str2 string, str3 string) {
	i, j, k := "a", "b", "c"
	return i, j, k
}
