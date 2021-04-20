package main

import (
	"fmt"
	"math/rand"
)

func main() {

	guessNumber()

}

//题11：猜数字，写一个程序，随机生成一个0~100的整数n，然后用户在终端输入数字，如果和n相等，则提示用户猜对了。如果不相等，则提示用户大于或小于n。
func guessNumber() {
	//① 生成一个 0~100 之间的随机整数
	randNumber := rand.Intn(100)
	for {
		//② 获取用户在终端输入的数字
		var inputNumber int
		fmt.Scanf("%d\n", &inputNumber)
		var flag bool
		switch {
		case inputNumber == randNumber:
			flag = true
			fmt.Println("恭喜你，猜对了！数字是：", randNumber)
		case inputNumber > randNumber:
			fmt.Println("猜大了")
		case inputNumber < randNumber:
			fmt.Println("猜小了")
		}
		if flag { // 如果才对了，flag为true则跳出循环
			break
		}
	}

}
