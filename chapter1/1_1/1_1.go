package __1

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")
	fmt.Println(os.Args)
	var a = []int{1, 2, 3, 4, 5}
	s := a[1:]
	s[1] = 11
	fmt.Println(a)
	fmt.Println(s)

}
