package main

import (
	"fmt"
	"strings"
)

func main() {
	//// 题1
	//multi()
	//
	//// 题2
	//for i := 1; i <= 1000; i++ {
	//	if isPrefectNumber(i) {
	//		fmt.Printf("完数：%d\n", i)
	//	}
	//}

	// 题3
	//isRollbackString()

	var mp map[string]int
	mp = make(map[string]int, 10)
	mp["a"] = 1

	fmt.Println(mp)
	mp = make(map[string]int, 10)
	fmt.Println(mp)
}

// 题1：编写程序，在终端输出九九乘法表
func multi() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, i*j)
		}
		fmt.Println()
	}
}

// 题2：一个数如果恰好等于它的因子之和，这个数就称为“完数”。例如6=1＋2＋3. 编程找出1000以内的所有完数。
func isPrefectNumber(n int) bool {
	var sum int
	var factor = make([]int, 0)
	for i := 1; i < n; i++ {
		if n%i != 0 {
			continue
		}
		// 找到了n的全部因子
		factor = append(factor, i)
	}

	for _, v := range factor {
		sum += v
	}
	if sum == n {
		return true
	} else {
		return false
	}
}

// 题3：输入一个字符串，判断其是否为回文。回文字符串是指从左到右读和从右到左读完全相同的字符串。
func isRollbackString() {
	fmt.Println("请输入待检测的字符串：")
	var inputStr string
	fmt.Scanf("%s\n", &inputStr) // 通过stnOs获取输入的字符串
	strLen := len(inputStr)      // 字符串长度
	if strLen == 0 {
		panic("不可输入空字符串！")
	}
	// 倒序循环排列字符
	strBytes := make([]string, strLen)
	for i := strLen - 1; i >= 0; i-- {
		strBytes = append(strBytes, string(inputStr[i]))
	}
	strJoinDesc := strings.Join(strBytes, "")
	if strJoinDesc == inputStr {
		fmt.Printf("%s 是回文！", strJoinDesc)
	} else {
		fmt.Printf("%s 不是回文...（原始字符串：%s， 旋转之后的字符串：%s）", inputStr, inputStr, strJoinDesc)
	}
}

// 题4：输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数。
func countString() {
	//unicode.IsNumber()
}

// 题5：计算两个大数相加的和，这两个大数会超过int64的表示范围.
