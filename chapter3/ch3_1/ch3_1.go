package main

import (
	"fmt"
)

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%09b\n", x&^y)

	var o = 0666
	fmt.Printf("%#d,%#[1]o,%#[1]x", o)
	s := "国家"
	c := '国'
	fmt.Printf("%q", s[0])
	fmt.Printf("%c,%[1]T", c)
}
