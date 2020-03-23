/*
   @Time : 2020-03-20 16:09
   @Author : liuoneice
   @File : demo1
   @Software: GoLand
*/
package main

import (
	"fmt"
	"runtime"
	"time"
)

type User struct {
	name string
	a    func()
}

func main() {
	strArray := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	strArray = append(strArray[3:5], strArray[4:6]...)
	fmt.Println(string(strArray))
}

func testMapValue() {
	//mm0 := make(map[string]int)
	//mm0["aa"] = User{
	//	name:"zhangsan",
	//	a: func() {
	//		fmt.Println("hello")
	//	},
	//}
	//mm0["aa"].name = "lisi"
	mm1 := make(map[string]string)
	mm1["jack1"] = "1231"
	mm1["jack2"] = "1232"
	mm1["jack3"] = "1233"
	fmt.Println(mm1)
	editMap(mm1)
	fmt.Println(mm1)
	mm2 := make(map[string]int)
	mm2["jack1"] = 1
	mm2["jack2"] = 2
	mm2["jack3"] = 3
	fmt.Println(mm2)
	mm2["jack3"] = 4
	fmt.Println(mm2)
	mm3 := make(map[string]float64)
	mm3["jack1"] = 1.1
	mm3["jack2"] = 1.2
	mm3["jack3"] = 1.3
	fmt.Println(mm3)
	mm3["jack3"] = 1.4
	fmt.Println(mm3)
	//mm5 := make(map[string]interface{})
	//mm5["jack1"] = 1
	//mm5["jack2"] = 1.2
	//mm5["jack3"] = User{
	//	name:"af",
	//}
	//fmt.Println(mm5)
	//User(mm5["jack3"]).name = "adfs"
	//fmt.Println(mm5)
	mm6 := make(map[string]chan int)
	mm6["jack1"] = make(chan int)
	mm6["jack2"] = make(chan int)
	mm6["jack3"] = make(chan int)
	fmt.Println(mm6)
	mm6["jack3"] = make(chan int)
	fmt.Println(mm6)
	mm7 := make(map[string][]int)
	mm7["jack1"] = []int{1}
	mm7["jack2"] = []int{2}
	mm7["jack3"] = []int{3}
	fmt.Println(mm7)
	mm7["jack3"][0] = 4
	fmt.Println(mm7)
	mm8 := make(map[string]*[3]int)
	mm8["jack1"] = &[3]int{1, 2, 3}
	mm8["jack2"] = &[3]int{4, 5, 6}
	mm8["jack3"] = &[3]int{7, 8, 9}
	fmt.Printf("%p", mm8["jack3"])
	fmt.Println(mm8)
	mm8["jack3"][2] = 10
	fmt.Printf("%p", mm8["jack3"])
	fmt.Println(mm8)
}
func editMap(mm map[string]string) {
	mm["jack3"] = "ree"
}

/**
字符串下标操作不可取，意味着golang的string是不可变的，只会产生新的str，而不会修改到原来的string
*/
func strIndextest() {
	str := "hello"
	// str[0] = "s"
	fmt.Println(str[0])
}

/**
append可以接受nil类型的切片作为参数
*/
func appendtest() {
	var i []int = make([]int, 0)
	fmt.Println(i)
	i = append(i, 1)
	i = append(i, 2)
	fmt.Println(len(i), cap(i))
}

/**
可变参数类型只能接受多个参数，不能接受该类型的数组或切片类型
*/
func testIntParams(args ...int) {
	//testIntParams(1,2,3)  正确
	//testIntParams([]int{4,5,6})  错误
	//testIntParams([3]int{4,5,6})  错误

	fmt.Println(args)
}

/**
可以使用``拼接""字符串
*/
func strCon() {
	str := `123` + "456"
	fmt.Println(str)
}

/**
启动多个goroutine，for循环比goroutine快
*/
func start_muti_groutine() {
	runtime.GOMAXPROCS(1)

	for i := 0; i < 100000; i++ {
		go func() {
			fmt.Println(i)
		}()
		//time.Sleep(time.Second*1)
	}
	time.Sleep(time.Second * 3)
}

/**
满足多个条件，随机出现异常
*/
func test_select() {
	var chan1 chan int = make(chan int, 1)
	var chan2 chan string = make(chan string, 1)
	chan1 <- 1
	chan2 <- "s"
	select {
	case i := <-chan1:
		{
			fmt.Println(i)

		}
	case s := <-chan2:
		{
			panic(s)
		}

	}
}

//func test_porint_index() {
//	var i string = "hello"
//	i[0] = 'x'
//	fmt.Println(i)
//}

//func test_nil() {
//	var i1 = nil
//	// err
//	var i2 int = nil
//	// err
//	var i3 string = nil
//	// err
//	var i4 bool = nil
//	// err
//	var i5 float64 = nil
//	// err
//	var i6 complex64 = nil
//	// err
//	var e error = nil
//	var i7 interface{} = nil
//	var i8 User = nil
//	// err
//	var i9 []int = nil
//	var i10 [1]int = nil
//	// err
//	var i11 map[string]string = nil
//	var i12 chan int = nil
//	var i13 func() = nil
//	fmt.Println(i1, i2, i3, i4, i5, i6, e, i7, i8, i9, i10, i11, i12, i13)
//}
