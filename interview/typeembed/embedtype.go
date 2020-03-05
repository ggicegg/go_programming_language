package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("TRACE:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	//log.Panicf("错误")
	log.Fatalf("错误")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA() // <==> t.People.ShowA() 所以ShowA里的p.ShowB()方法传入的是People
	i := recover()
	fmt.Printf("%T,%v", i, i)
	t.ShowB()
	fmt.Println("hello")
}
