package __2

import (
	"fmt"
	"os"
	"strings"
)

//func main() {
//	practice1_2()
//}

func practice1_2() {
	for index, value := range os.Args[1:] {
		fmt.Println(index, value)
	}
}

func funcName3() (int, error) {
	return fmt.Println(os.Args[1:])
}

func FuncName2() (int, error) {
	return fmt.Println(strings.Join(os.Args[1:], " "))
}

func FuncName() {
	var s, blank string
	for i := 1; i < len(os.Args); i++ {
		s += blank + os.Args[i]
		blank = " "
	}
	fmt.Println(s)
}
