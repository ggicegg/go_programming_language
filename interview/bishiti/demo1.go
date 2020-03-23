/*
   @Time : 2020-03-19 14:19
   @Author : liuoneice
   @File : demo1
   @Software: GoLand
*/
package bishiti

import "github.com/shopspring/decimal"

//计算结果缓存
var buf map[int64]decimal.Decimal

func init() {
	buf = make(map[int64]decimal.Decimal)
	buf[0] = decimal.NewFromInt(0)
	buf[1] = decimal.NewFromInt(1)
	buf[2] = decimal.NewFromInt(1)
}

/**
粗略的解法，重复计算量比较大
*/
func fib1(i int64) int64 {
	if i < 2 {
		return i
	}

	return fib1(i-1) + fib1(i-2)
}

/**
优化的斐波那契数列求值解法
使用了缓存来保存已经计算过的值，防止重复计算
*/

func fib2(i int64) decimal.Decimal {
	if i <= 2 {
		return buf[i]
	}
	if buf[i].String() == "0" {
		buf[i] = decimal.Sum(fib2(i-1), fib2(i-2))
	}
	return buf[i]
}
