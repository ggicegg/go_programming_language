package ch2_6

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//使用循环
func PopCount2(x uint64) int {
	var i uint64
	var sum int
	for i = 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}

//使用计算最右bit的方法
/**
将x右移一位，然后将右移产生的数据左移一位，将左移后的数据与原始x的值异或操作，如果x原始值最右bit是1，异或后的结果为1，如果最右
bit是0，异或后的结果一定为0
 */
func PopCount3(x uint64) int {
	var i uint64
	var sum int
	for i = 0; i < 63; i++ {
		t0 := x
		t1 := x >> 1
		t2 := t1 << 1
		if t0^t2 == 1 {
			sum++
		}
		x = t1
	}
	return sum
}

//
func PopCount4(x uint64) int {
	var sum int
	for i := 0; i < 63; i++ {
		t := x
		x = x & (x - 1)
		if (t == x+1) {
			sum++
		}
		u := t >> 1
		x = u
	}
	return sum
}

//func main() {
//	fmt.Println(PopCount2(5))
//}
