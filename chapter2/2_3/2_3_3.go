package main

import "fmt"

type User struct{
	name string
	age int
}

type Mananger interface {
	manager()
}

func (u *User)manager(){
	fmt.Println("Hello")
}

func main() {
	var i1 = new(int)
	i2 := new(int)
	fmt.Println(i1,&i1,*i1,i2,&i2,*i2)

	f1 := new(float32)
	s1 := new(string)
	u1 := new(User)
	fmt.Println(f1,&f1,*f1,s1,&s1,*s1,u1,&u1,*u1)
	u1.manager()

	it1 := new(interface{})
	it2 := new(interface{})
	fmt.Println(it1,&it1,*it1,it2,&it2,*it2)


	it3 := new(Mananger)

	fmt.Println(it3,&it3,*it3)

	for i:=3;i != 6;i++{
		j := *i1
		fmt.Println(j,&j)
	}

	x := 1
	y := 2
	x,y = y,x
	fmt.Println(x,y)

}