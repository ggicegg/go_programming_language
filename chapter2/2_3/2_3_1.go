package main

import "fmt"

func main() {
	var i *int
	mm := 1
	i = &mm
	fmt.Println("before", &i, i, *i)
	j := add(i)

	fmt.Println("after", &j, j, *j)

}

func add(i *int) *int {
	id := &i
	t := 1
	*i += t
	fmt.Println("add", id, i, *i)
	return i
}
