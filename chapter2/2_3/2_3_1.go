package main

import "fmt"


func main() {
	var i *int
	mm := 1
	i = &mm
	fmt.Println(&i,i,*i)
	j := add(i)

	fmt.Println(&j,j,*j)

}

func add(i *int) *int{
	id := &i
	t := 1
	*i+=t
	fmt.Println(id,i,*i)
	return i
}