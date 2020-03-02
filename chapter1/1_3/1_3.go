package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	//conts := make(map[string]int)

	var conts map[string]int = make(map[string]int)
	fileNames := os.Args[1:]
	if 0 == len(fileNames) {
		countStrings(os.Stdin, conts)
	}

	for _, fileName := range fileNames {
		//fileName 相对路径 应该是src/go_pragramming_language/chapter1/1_3/a.txt这样的
		countStringsByOpen(fileName, conts)
		//countStringsByReadFile(fileName, conts)
	}
	printConts(conts)
}

func countStringsByReadFile(fileName string, conts map[string]int) {
	bytes, i := ioutil.ReadFile(fileName)
	if i != nil {
		log.Fatal(i)
	}
	words := strings.Split(string(bytes), "\n")
	for _, word := range words {
		conts[word]++
	}
}

func countStringsByOpen(fileName string, conts map[string]int) {
	file, e := os.Open(fileName)
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()
	countStrings(file, conts)
}

func countStrings(inStream *os.File, conts map[string]int) {
	scanner := bufio.NewScanner(inStream)
	// mac中输入command+d终止输入 windows中输入ctrl+d
	for scanner.Scan() {
		conts[scanner.Text()]++
	}
	defer inStream.Close()
}

func printConts(conts map[string]int) {
	if conts == nil || len(conts) == 0 {
		return
	}
	for key, value := range conts {
		fmt.Printf("%d\t%s\n", value, key)
	}
}
