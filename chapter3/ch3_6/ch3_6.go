package main

import "fmt"

const (
	i = 1
	s = len("helo")
	c = 'c'
	d = 12.234
)

type weekday int

const (
	monday weekday = iota + 1
	tuesday
	wednesday
	thursday = iota
	friday
	saturday
	sunday
)

const (
	_ float64 = 8 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func init() {
	//fmt.Println(i)
	//fmt.Println(s)
	//fmt.Println(c)
	//fmt.Println(d)
	//
	//fmt.Println(monday)
	//fmt.Println(tuesday)
	//fmt.Println(wednesday)
	//fmt.Println(thursday)
	//fmt.Println(friday)
	//fmt.Println(saturday)
	//fmt.Println(sunday)

	fmt.Println(KB)
	//fmt.Println(MB)
	//fmt.Println(GB)
	//fmt.Println(TB)
	//fmt.Println(PB)
	//fmt.Println(EB)
	//fmt.Println(ZB)
	//fmt.Println(YB)

	fmt.Println(YB / ZB)
}

func main() {
	fmt.Printf("%T\t%[1]v\n", 0)
	fmt.Printf("%T\t%[1]v\n", 0.0)
	fmt.Printf("%T\t%[1]v\n", 0i)
	fmt.Printf("%T\t%[1]v\n", '\000')
}
