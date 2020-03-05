package main

import (
	"crypto/sha256"
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

var symbol = [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

func main() {
	fmt.Println(sha256.Sum224([]byte{'x'}))
	fmt.Println(sha256.Sum224([]byte{'X'}))
}

func printConst() {
	fmt.Println(symbol[RMB], EUR, GBP, RMB)
	fmt.Println(XXX, SSS, AAA, BBB, CCC, DDD, EEE, FFF, GGG, HHH, III)
}
