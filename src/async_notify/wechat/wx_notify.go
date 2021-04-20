package wechat

import (
	"async_notify/curl"
	"async_notify/lt_log"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// 微信通知
const (
	corpId     string = "wx25bef61ee7353fa3"                                               // 公司id
	corpSecret string = "jjEPP-T-epRTvX-CB8UFja7tTPfJLjVhReL9YH0MduokuJ-GD1LvqAKl1K28tYuc" //公司secret
	tokenUrl   string = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"                     // 获取token的url
	notifyUrl  string = "https://qyapi.weixin.qq.com/cgi-bin/message/send"                 // 发送微信通知的url

	// 错误日志名称
	errorLogName string = "weChat_Err"
)

/**
 * 发送微信通知
 */
func Sending(receivers []string, msg string) {
	// ①获取微信官方token
	responseWxJson := curl.HttpGet(fmt.Sprintf("%s?corpid=%s&corpsecret=%s", tokenUrl, corpId, corpSecret)) //腾讯回执
	// ② 解码json字符串成map
	var responseWx map[string]interface{}
	err := json.Unmarshal([]byte(responseWxJson), &responseWx)
	// ③ 初始化乐堂日志助手结构体
	logHandler := new(lt_log.LogHandler)
	if err != nil {
		logHandler.Writer(errorLogName, "腾讯微信接口响应返回格式错误", responseWxJson) //记录错误日志
		log.Fatal("腾讯微信token接口返回格式错误", err)                               // 打印错误信息
	}
	// ④ 发送微信请求
	accessToken, ok := responseWx["access_token"] // ok为true时，表示取到对应值
	if !ok {
		logHandler.Writer(errorLogName, "腾讯微信access_token不存在", accessToken.(string)) //记录错误日志
		log.Fatal("腾讯微信token接口返回格式错误")                                               // 打印错误信息
	}
	// 微信发送的准备参数
	sendingArgus := map[string]interface{}{
		"touser":  strings.Join(receivers, "|"), // 发送人用竖线“|”连接
		"msgtype": "text",
		"agentid": "0",
		"safe":    "0",
		"text": map[string]interface{}{
			"content": msg,
		},
	}
	//resp := curl.HttpPost(fmt.Sprintf("%s?access_token=%s", NotifyUrl, accessToken), sendingArgus)
	//fmt.Println(resp)
	sendingArgusBytes, _ := json.Marshal(sendingArgus)
	fmt.Println("in 发微信完成...")
	responseOfWx := curl.HttpPostJson(fmt.Sprintf("%s?access_token=%s", notifyUrl, accessToken), sendingArgusBytes)
	fmt.Println(responseOfWx)
}

/**
 * map转换为string
 */
func mapToString(p map[string]string) string {
	fmt.Println("p : ", p)
	pBytes, _ := json.Marshal(p)

	fmt.Println("pBytes : ", pBytes)
	fmt.Println("string(pBytes) : ", string(pBytes))
	return string(pBytes)
}
