package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	s1 := "1233445345345"
	s2 := "1231234.12341234"
	fmt.Println(comma(s1))
	fmt.Println(commaBuffer(s1))

	var t string = s2
	floindex := strings.LastIndex(s2, ".")
	if floindex > 0 {
		s2 = s2[:floindex]
		fmt.Println(commaBuffer(s2) + t[floindex:])
	} else {
		fmt.Println(commaBuffer(s2) + t[floindex:])

	}

	//fmt.Println(comma(s2)+comma(s2[floindex:])
	//fmt.Println(commaBuffer(s2)+t[floindex:])

}

//3.10
func commaBuffer(s string) string {
	var buf bytes.Buffer
	n := len(s)
	for i := n; i > 0; i-- {
		if i != n && i%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(rune(s[n-i]))
	}

	return buf.String()

}
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + comma(s[n-3:])
}
func intstostringtest() {
	var is []int = []int{1, 4, 5}
	bs := intstostring(is)
	fmt.Println(bs)
}

func intstostring(is []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range is {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func runeFileds() {
	s1 := "asdf asdf asdfwqe qwer 测试 哈哈"
	fmt.Println(strings.Fields(s1))
	fmt.Println(strings.FieldsFunc(s1, func(r rune) bool {
		if r == rune('测') {
			return true
		}
		return false
	}))
}

func parsetest() {
	s1 := "123.324"
	s2 := "123"
	s3 := "abc"
	//fmt.Println(string([]byte(s1)))
	r1, err := strconv.ParseFloat(s1, 64)
	if err != nil {

	}
	r2, err := strconv.ParseInt(s2, 0, 0)
	if err != nil {

	}
	r3, err := strconv.ParseInt(s3, 0, 0)
	if err != nil {

	}
	r4, err := strconv.ParseInt(s1, 0, 0)
	if err != nil {

	}
	fmt.Printf("%v\n", r1)
	fmt.Printf("%v\n", r2)
	fmt.Printf("%v\n", r3)
	fmt.Printf("%v\n", r4)
}

func strtest() {
	s := "hello, world"
	hello := s[:5]
	world := s[7:]
	fmt.Printf("%T,%v,%T,%v \a\a\a", hello, world, hello, world)
	s1 := `这是一个原始的字符串abde`
	s2 := "AA  eee wer123 这是一个原始的字符串abde"
	fmt.Println(utf8.RuneCountInString(s1))
	fmt.Println(utf8.RuneCount([]byte(s1)))
	for i := 0; i < len(s2); {
		r, size := utf8.DecodeRuneInString(s2[i:])
		fmt.Printf("%d\t%c IsDigit:%t;IsLetter:%t;IsLower:%t;IsUpper:%t\n",
			i, r, unicode.IsDigit(r), unicode.IsLetter(r), unicode.IsLower(r),
			unicode.IsUpper(r))
		i += size
	}
	s3 := "/adfs/dsf/asdf.go"
	fmt.Println("s3:", basename(s3))
	s4 := "asdf.go"
	fmt.Println("s4:", basename(s4))
	s5 := "asdf"
	fmt.Println("s5:", basename(s5))
	s6 := "asdf.go"
	fmt.Println("s6:", basename(s6))
	//for index,char := range s1{
	//	fmt.Printf("index:%d,char:%c\n",index,char)
	//}
	fmt.Println(findPrefix(s1, "这是1"))
	fmt.Println(contains(s2, "这是"))
	fmt.Println(findSurfix(s1, "符串1"))
	fmt.Println(contains(s1, "符串"))
	fmt.Println(contains(s1, "原始"))
	fmt.Println(len(s1))
}

func basename(s string) string {
	result := ""
	if len(s) > 0 {
		end := strings.LastIndex(s, ".")
		if end == -1 {
			end = len(s)
		}
		start := strings.LastIndex(s, "/")
		start++

		result = s[start:end]

	}
	return result
}

func findPrefix(s string, prefix string) bool {
	if len(s) >= len(prefix) && s[:len(prefix)] == prefix {
		return true
	}
	return false
}
func findSurfix(s string, surfix string) bool {
	if len(s) >= len(surfix) && s[len(s)-len(surfix):] == surfix {
		return true
	}
	return false
}

func contains(s string, substr string) bool {
	if len(s) < len(substr) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if findPrefix(s[i:], substr) {
			return true
		}
	}
	return false

}
