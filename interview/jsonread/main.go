/*
   @Time : 2020-03-22 13:50
   @Author : liuoneice
   @File : main
   @Software: GoLand
*/
package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	_ "github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {

	}
	bytes, e := ioutil.ReadFile(dir + "/data/data.json")
	if e != nil {
		log.Fatal(e.Error())
	}
	gjson.GetBytes(bytes, "list").ForEach(func(key, value gjson.Result) bool {
		name := value.Get("name").String()
		url := value.Get("url").String()
		fmt.Println(name, url)
		return true
	})
}
