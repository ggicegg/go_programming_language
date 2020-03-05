package main

import (
	"fmt"
	"math"
)

func main() {
	var f1, f2 float32 = 0.3, 0.2
	var f3, f4 float64 = 0.3, 0.2
	fmt.Println(f1 - f2)
	fmt.Println(f3 - f4)
	fmt.Printf("%g,%g", math.MaxFloat32, math.MaxFloat64)
	fmt.Println(math.IsNaN(math.NaN()))
}
