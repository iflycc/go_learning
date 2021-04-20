package curl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"unsafe"
)

/**
 * Http - 发送Get请求
 * @param url string 请求url地址
 */
func HttpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(fmt.Sprintf("请求`%s`失败，err ： %v", url, err))
	}

	defer func() { // 关闭请求资源
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("读取 resp.body err ： %v", err))
	}
	return string(body)
}

/**
 * Http - 发送Post请求
 * @param webUrl string 请求url地址
 * @param formatParams map[string]string) 请求参数
 */
func HttpPost(webUrl string, formatParams map[string]interface{}) string {
	argus := url.Values{} //参数容器
	for _key, _val := range formatParams {
		argus.Add(_key, _val.(string))
	}
	resp, err := http.PostForm(webUrl, argus) // 请求远程url地址

	defer func() { // 延迟关闭连接句柄
		_ = resp.Body.Close()
	}()
	if err != nil {
		panic(err) // 打印错误，抛出panic
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

func HttpPostJson(webUrl string, jsonBytes []byte) string {
	reader := bytes.NewReader(jsonBytes)
	request, err := http.NewRequest("POST", webUrl, reader)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	return *str
}
