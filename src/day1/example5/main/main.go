package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	//
	jsonBytes, _ := json.Marshal(r.PostForm)
	jsn := string(jsonBytes)
	fmt.Printf("jsn : %s", jsn)

	fmt.Fprintf(w, "你好，学院君！[method:%v]", r.Method) // 发送响应到客户端
}

func main() {

	fmt.Println(strconv.FormatFloat(0.2225521447, 'f', -1, 32))
	fmt.Println(strconv.FormatFloat(0.2225521447, 'g', -1, 32))
	fmt.Println(strconv.FormatFloat(0.2225521447, 'f', 2, 64))
	fmt.Println(strconv.FormatFloat(0.2225521447, 'g', 2, 64))
	return

	http.HandleFunc("/", sayHelloWorld)
	err := http.ListenAndServe("127.0.0.1:9092", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
