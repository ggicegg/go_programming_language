package main

import "fmt"

// Package tempconv performs Celsius and Fahrenheit temperature computations.
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度
const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius)ToString() (ss string){
	return fmt.Sprintf("%g C",c)
}
func main() {
	fmt.Println(Celsius(34.234).ToString())
	i := 3
	fmt.Printf("%p,%p",&i,show(i))
}


func show(i int)(l int){
	defer func() {
		i = i + 5
	}()

	return i
}