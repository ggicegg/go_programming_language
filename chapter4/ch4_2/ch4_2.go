package main

import "fmt"

func main() {
	var s1 []string
	s1 = append(s1, "a") //append可以追加到一个nil切片上
	fmt.Println(s1)
}

func testEquals() {
	s1 := []string{"a", "b"}
	s2 := []string{"a", "b"}
	fmt.Println(equals(s1, s2))
}
func equals(s1, s2 []string) bool {
	n1 := len(s1)
	n2 := len(s2)
	if n1 != n2 {
		return false
	}
	for i := 1; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
