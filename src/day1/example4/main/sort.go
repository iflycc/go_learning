package main

import "fmt"

func main() {
	//1 : 冒泡排序
	sl1 := []int{10, 1, 3, 4, 5, 8, 7, 9, 2, 6}
	moPoSort(sl1)
	fmt.Println(sl1)
	//2 : 选择排序
	sl2 := []int{10, 1, 3, 4, 5, 8, 7, 9, 2, 6}
	selectSort(sl2)
	fmt.Println(sl2)
	//3 : 插入排序
	sl3 := []int{10, 1, 3, 4, 5, 8, 7, 9, 2, 6}
	insertSort(sl3)
	fmt.Println(sl3)

}

//① 冒泡排序
func moPoSort(sl []int) {
	lengthSl := len(sl) // 获取切片的长度

	for i := lengthSl; i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if sl[j] <= sl[j+1] {
				continue
			}
			// 交换位置
			sl[j], sl[j+1] = sl[j+1], sl[j]
		}
	}
}

//② 选择排序
func selectSort(sl []int) {
	lengthSl := len(sl) // 获取切片的长度

	for i := 0; i < lengthSl; i++ {
		_min := i
		for j := i + 1; j < lengthSl; j++ {
			if sl[_min] > sl[j] {
				_min = j
			}
		}
		sl[i], sl[_min] = sl[_min], sl[i]
	}
}

//③ 插入排序
func insertSort(sl []int) {
	lengthSl := len(sl) //获取切片的长度

	for i := 0; i < lengthSl; i++ {
		for j := i; j > 0; j-- {
			if sl[j] > sl[j-1] {
				break
			}
			sl[j], sl[j-1] = sl[j-1], sl[j]
		}
	}
}
