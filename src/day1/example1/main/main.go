package main

import "fmt"

/*
 * 打印出 100~200之间的全部素数
 */
func main() {
	var cnt, startNumber, endNumber = 0, 100, 200
	for i := startNumber; i <= endNumber; i++ {
		if isPrime(i) {
			cnt++
			fmt.Printf("100~200之间，第%d个素数为：%d\n", cnt, i)
		}
	}
}

/**
 * 判断一个整数是否为素数
 */
func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	} else if n%2 == 0 || n%3 == 0 { //当一个数能被 2 or 3 整除，则不是素数
		return false
	} else { // 循环判断，当一个数，只能被1和自己整除，则为素数
		for i := 5; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
}
