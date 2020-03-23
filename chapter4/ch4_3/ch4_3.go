package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	var u0 user
	fmt.Println(u0)
	var mm = make(map[user]string)
	var mm1 = make(map[user]string)
	u1 := user{
		name: "张三",
		age:  12,
	}
	u2 := user{
		name: "张三",
		age:  12,
	}
	mm[u1] = "123"
	mm[u2] = "1234"
	fmt.Println(mm)
	u3 := user{
		name: "张三",
		age:  12,
	}
	mm[u1] = "123"
	mm1[u2] = "123"
	fmt.Println(mm1 == mm)
	if value, ok := mm[u3]; ok {
		fmt.Println(value)
		fmt.Printf("%T %[1]v\n", mm[u3])

	}
}
