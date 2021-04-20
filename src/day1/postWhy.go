package day1

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数

	fmt.Println(r.Form)             // 在服务端打印请求参数
	fmt.Println("URL:", r.URL.Path) // 请求 URL
	fmt.Println("Scheme:", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println(k, ":", strings.Join(v, ""))
	}
	// 第一种方式
	name := r.URL.Query().Get("name")
	fmt.Printf("-name- : %v\n", name)
	fmt.Printf("---get body-----  :   %v\n", r.GetBody)
	fmt.Printf("---post argus---  :   %v\n", r.PostForm)
	fmt.Println("----------------------------------------------end----------------------------------------------")

	fmt.Fprintf(w, "你好，学院君！[method:%v]", r.Method) // 发送响应到客户端
}

func main() {
	http.HandleFunc("/", sayHelloWorld)
	err := http.ListenAndServe("127.0.0.1:9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
