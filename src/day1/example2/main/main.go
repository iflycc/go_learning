package main

import "fmt"

/**
 * 水仙花数，指3位整数，如：abc = a² + b² + c²，即为水仙花数，编写打印 100~999 之间的全部水仙花数
 */

func main() {
	var cnt int
	for i := 100; i <= 999; i++ {
		if isNarc(i) {
			cnt++
			fmt.Printf("第 %d 个水仙花数为：%d\n", cnt, i)
		}
	}
}

/**
 * 判断一个函数是否是水仙花数
 */
func isNarc(n int) bool {
	if n < 100 || n > 999 {
		panic("当前整数传递不正确，请传入 100 ~ 999 之间的整数!!!\n")
	}
	// 判断是否是水仙花数
	var hundred, ten, unit int
	hundred = int(n / 100)            //百位
	ten = int((n - hundred*100) / 10) // 十位
	unit = n - hundred*100 - ten*10   // 个位
	if n == hundred*hundred*hundred+ten*ten*ten+unit*unit*unit {
		return true
	} else {
		return false
	}
}
