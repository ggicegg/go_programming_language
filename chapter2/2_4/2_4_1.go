package main

import (
	"fmt"
)

func main() {

	//var i0 int = 1
	//var i1 int = 1
	//var i2 int8 = 1
	//var i3 int16 = 1
	//var i4 int32 = 1
	//var i5 int64 = 1
	//fmt.Println(i0==i1)
	//fmt.Println(i1==i1)
	//fmt.Println(i1==i2)
	//fmt.Println(i1==i3)
	//fmt.Println(i1==i4)
	//fmt.Println(i1==i5)
	i := 1.0
	j := i / 3
	k := j * 3
	fmt.Println(i,j,k)

	l := 1.01
	m := 1.02
	fmt.Println(l+m)

	o := 0.3
	p := 0.2

	fmt.Println(o-p)
	fmt.Println((float64(0.3) - float64(0.2)))
	fmt.Println(float64(0.3) - float64(0.2))
	fmt.Println(float64(0.3) - 0.2)
	fmt.Println(0.3 - float64(0.2))
	fmt.Println(0.3 - 0.2)
	fmt.Println(float32(0.3) - float32(0.2))
	fmt.Println(0.3 - float32(0.2))
	fmt.Println(float32(0.3) - 0.2)

	k1 := 3.234
	fmt.Println(fmt.Sprintf("%g C",k1))

}

func funcName2() {
	var i interface{}
	i = 3
	value, ok := i.(float32)
	fmt.Println(value, ok, i)
}

func funcName1() {
	fmt.Println(gcd(15, 10))
	fmt.Println(gcd(27, 18))
	fmt.Println(gcd(15, 10))
	fmt.Println(gcd(27, 18))
	// 0 1 1 2 3
	fmt.Print(fibnacci(1))
	fmt.Print(fibnacci(2))
	fmt.Print(fibnacci(3))
	fmt.Print(fibnacci(4))
	fmt.Print(fibnacci(5))
	fmt.Print(fibnacci(6))
	fmt.Print(fibnacci(7))
	fmt.Println()
	fmt.Print(fibnacci2(1))
	fmt.Print(fibnacci2(2))
	fmt.Print(fibnacci2(3))
	fmt.Print(fibnacci2(4))
	fmt.Print(fibnacci2(5))
	fmt.Print(fibnacci2(6))
	fmt.Print(fibnacci2(7))
	fmt.Println()
	fmt.Print(fibnacci_recursive(1))
	fmt.Print(fibnacci_recursive(2))
	fmt.Print(fibnacci_recursive(3))
	fmt.Print(fibnacci_recursive(4))
	fmt.Print(fibnacci_recursive(5))
	fmt.Print(fibnacci_recursive(6))
	fmt.Print(fibnacci_recursive(7))
}

func fibnacci(i int) int {
	x, y := 0, 1

	for j := 1; j < i; j++ {
		t := y
		y = x + y
		x = t
	}
	return x
}

func fibnacci2(i int) int {
	x, y := 0, 1
	if i == 1{
		return 0
	}
	for j := 1; j < i; j++ {
		x,y=y,x+y
		//t := y
		//y = x + y
		//x = t
	}
	return x
}

func fibnacci_recursive(i int)int{
	if i == 1{
		return 0
	}
	if i == 2{
		return 1
	}

	return fibnacci_recursive(i-1)+fibnacci_recursive(i-2)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

func gcd2(x, y int) int {
	for y != 0 {
		t := y
		y = x % y
		x = t
	}

	return x
}
