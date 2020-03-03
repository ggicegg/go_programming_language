package ch2_6

import "fmt"

var a = b + c
var c = 1
var b = 2

func init() {
	fmt.Println("hello1")
}

func init()  {
	fmt.Println("hello2")
}

func init()  {
	fmt.Println("hello3")
}
//func main() {
//	i1 := 5
//	i2 := 1
//	fmt.Println(i1^i2)
//}