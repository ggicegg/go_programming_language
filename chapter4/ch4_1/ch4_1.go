package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

type CURRENCY int

const (
	XXX        = 0    // 0
	SSS        = 0    // 0
	AAA int    = iota // 2 当前位置建1 （3-1）
	BBB               // 3
	CCC = 1           // 1
	DDD = 2           // 2
	EEE = iota        // 6 当前位置减1（5-1）
	FFF               // 7
	GGG = 1           // 1
	HHH = iota        // 9（8-1）
	III               // 10
)

const (
	USD CURRENCY = iota
	EUR
	GBP
	RMB
)

//4.2 编写一个程序，默认打印标准输入的以SHA256哈希码，也可以通过命令行标准参 数选择SHA384或SHA512哈希算法
func SHAAdapter(t int, x []byte) {
	if x != nil && len(x) > 0 {
		switch t {
		case 512:
			result := sha512.Sum512(x)
			fmt.Println(result)
			break
		case 384:
			result := sha512.Sum384(x)
			fmt.Println(result)
			break
		case 256:
		default:
			result := sha256.Sum256(x)
			fmt.Println(result)
		}
	}
}

var symbol = [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

func main() {
	SHAAdapter(384, []byte{'1'})
	//c1 := sha256.Sum256([]byte{'x'})
	//c2 := sha256.Sum256([]byte{'X'})
	//fmt.Printf("%x %x ,%T,%T,%v %v", c1, c2, c1, c2, len(c1), len(c2))
}

func printConst() {
	fmt.Println(symbol[RMB], EUR, GBP, RMB)
	fmt.Println(XXX, SSS, AAA, BBB, CCC, DDD, EEE, FFF, GGG, HHH, III)
}
