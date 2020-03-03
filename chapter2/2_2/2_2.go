package main

import "fmt"

func main() {
	var f1,f2 = 50.0,100.0
	fmt.Printf("f1:%f°F <==> %f°C\n",f1,fToC(f1))
	fmt.Printf("f2:%f°F <==> %f°C",f2,fToC(f2))
}

func fToC(f float64) float64{
	return (f-32) * 5/ 9
}
