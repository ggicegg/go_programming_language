package main

import (
	"fmt"
	"net"
)

type user struct{
	name string
	age int
	//ad *int
}

func main() {
	// 	变量声明
	// 	var 变量名 变量类型 = 变量值

	var a int = 1
	var ad *int = &a
	var ad2 *int
	fmt.Println(a,&a,ad,&ad,ad2)
	var b = 2
	var c int
	d := 1
	fmt.Println(a,b,c,d)
	fmt.Printf("%T,%T,%T,%T\n",a,b,c,d)


	var e string = "hello"
	var f = "hello"
	g := "g"
	var h string
	var j string
	fmt.Println(e,f,g,h)
	fmt.Printf("%T,%T,%T,%T,%T\n",e,f,g,h,j)
	fmt.Printf("%p,%p,%p,%p,%p\n",&e,&f,&g,&h,&j)


	var af1 float32 = 1.0
	var bf1 = 2.0
	var cf1 float32
	df1 := 3.0
	fmt.Println(af1,bf1,cf1,df1)
	fmt.Printf("%T,%T,%T,%T\n",af1,bf1,cf1,df1)

	var af2 float64 = 1.0
	var bf2 = 2.0
	var cf2 float64
	df2 := 3.0
	fmt.Println(af2,bf2,cf2,df2)
	fmt.Printf("%T,%T,%T,%T\n",af2,bf2,cf2,df2)

	var i int = 9223372036854775807
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var irune rune = 2147483647
	fmt.Println(i,i8,i16,i32,i64,irune)
	fmt.Printf("%T,%T,%T,%T,%T,%T\n",i,i8,i16,i32,i64,irune)

	var ui uint = 9223372036854775808*2-1
	var ui8 uint8 = 128*2-1
	var ui16 uint16 = 32768*2-1
	var ui32 uint32 = 2147483648*2-1
	var ui64 uint64 = 0xFFFFFFFFFFFFFFFF
	var uiptr uintptr = 0xFFFFFFFFFFFFFFFF

	fmt.Println(ui,ui8,ui16,ui32,ui64,uiptr)
	fmt.Printf("%T,%T,%T,%T,%T,%T\n",ui,ui8,ui16,ui32,ui64,uiptr)

	var c64 complex64 = 3+5i
	var c128 complex128 = 3+8i
	fmt.Println(c64,c128)
	fmt.Printf("%T,%T\n",c64,c128)


	var err1 error
	err2 := &net.AddrError{

	}
	fmt.Println(err1,err2)
	fmt.Printf("%T,%T\n",err1,err2)

	var u1 user = user{"张三",12}
	var u2 user = user{name:"阿凡达说",age:56}
	fmt.Println(u1,u2)



}

