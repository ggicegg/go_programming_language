package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	url := "http://www.baidu.com"
	resp, err := http.Get(url)
	defer resp.Body.Close()

	fmt.Println(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("状态码:", resp.Status)
	fmt.Println("Header:", resp.Header)
	fmt.Println("Proto:", resp.Proto)
	fmt.Println("StatusCode:", resp.StatusCode)
	fmt.Println("ContentLength:", resp.ContentLength)
	//written, err := io.Copy(os.Stdout, resp.Body)
	//if err != nil{
	//	os.Exit(1)
	//}
	//fmt.Println(written)
	//resposeBodyText, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(resposeBodyText))
}
