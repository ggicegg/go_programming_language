package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string
var err error

func init() {
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("error...%v", err)
	}
}

func main() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
			// "HELLO" (one letter per iteration)
		}
	}
}
