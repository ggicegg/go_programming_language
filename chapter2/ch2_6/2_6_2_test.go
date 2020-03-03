package ch2_6_test

import (
	"fmt"
	"go_programming_language/chapter2/ch2_6"
	"testing"
)
var table = [10]int{0,1,1,2,1,2,2,3,1,2}

func TestPopCount1(t *testing.T) {
	// 5:两个1  6：两个1 7：三个1
	for i:=0;i<10;i++{
		temp := ch2_6.PopCount1(uint64(i))
		if table[i] != temp{
			t.Error(fmt.Sprintf("%v 计算位数为%v,计算错误",i,temp))
		}
	}

}

func TestPopCount2(t *testing.T) {
	// 5:两个1  6：两个1 7：三个1
	for i:=0;i<10;i++{
		temp := ch2_6.PopCount2(uint64(i))
		if table[i] != temp{
			t.Error(fmt.Sprintf("%v 计算位数为%v,/ch计算错误",i,temp))
		}
	}

}

func TestPopCount3(t *testing.T) {
	// 5:两个1  6：两个1 7：三个1
	for i:=0;i<10;i++{
		temp := ch2_6.PopCount3(uint64(i))
		if table[i] != temp{
			t.Error(fmt.Sprintf("%v 计算位数为%v,计算错误",i,temp))
		}
	}

}

func TestPopCount4(t *testing.T) {
	// 5:两个1  6：两个1 7：三个1
	for i:=0;i<10;i++{
		temp := ch2_6.PopCount4(uint64(i))
		if table[i] != temp{
			t.Error(fmt.Sprintf("%v 计算位数为%v,计算错误",i,temp))
		}
	}

}

func BenchmarkPopCount1(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ch2_6.PopCount1(uint64(i))
	}
}

func BenchmarkPopCount2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ch2_6.PopCount2(uint64(i))
	}
}


func BenchmarkPopCount3(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ch2_6.PopCount3(uint64(i))
	}
}


func BenchmarkPopCount4(t *testing.B) {

	for i := 0; i < t.N; i++ {
		ch2_6.PopCount4(uint64(i))
	}
}




