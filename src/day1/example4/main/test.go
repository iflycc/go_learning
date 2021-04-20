package main

import (
	"fmt"
	"sync"
	"time"
)

var mut sync.Mutex

func printer(str string) {
	mut.Lock()
	defer mut.Unlock()
	for _, data := range str {
		fmt.Printf("%c", data)
	}
	fmt.Println()
}
func person1() {
	printer("hello")
}
func person2() {
	printer("world")
}
func main() {
	go person1()
	person2()
	time.Sleep(time.Second)
}
