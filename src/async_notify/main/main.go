package main

import (
	"async_notify/lt_log"
	"async_notify/mysql"
	"async_notify/wechat"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var recordId string             // 定义一个全局变量，记录id，全局使用，作为本次请求的唯一记录id
var db mysql.AsyncNotifyRecords // mysql全局变量，记录通知表

func init() {
	recordId = strconv.Itoa(int(time.Now().UnixNano()))
	db = mysql.AsyncNotifyRecords{
		RecordId:      recordId,
		RequestParams: "", //请求参数
		RequestDt:     "", //请求时间点
		ResponseMsg:   "", //响应信息
		ResponseDt:    "", //响应时间点
		RetryTimes:    0,  //重发次数
		ExtField1:     "", //备用字段1
		ExtField2:     "", //备用字段2
		ExtField3:     "", //备用字段3
		IsOk:          1,  //是否处理成功，0：失败，1：成功
		CreatedAt:     "", //创建时间
	}
}

// 程序执行入口
func main() {
	http.HandleFunc("/", ltHttpServer)
	err := http.ListenAndServe("127.0.0.1:9091", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}

	// 接收参数
	//asyncDb := mysql.AsyncNotifyRecords{RecordId: "123", RequestDt: "2012-01-01 01:01:01", ResponseDt: "2012-01-01 02:02:02"}
	//asyncDb.Insert()
}

/**
 * Http请求:
 */
func ltHttpServer(writer http.ResponseWriter, request *http.Request) {
	startTime := time.Now()
	startDt := startTime.Format("2006-01-02 15:04:05") //开始执行的时间点
	//var clientParams map[string]string
	// 根据请求方法处理
	var requestParams url.Values //定义一个接收全部请求参数的变量
	var requestParamsJson string //json化后的请求参数
	switch request.Method {
	case "GET":
		query := request.URL.Query() // 接收Get请求
		requestParams = query
		requestParamsBytes, _ := json.Marshal(requestParams)
		requestParamsJson = string(requestParamsBytes) // 将字节切片转换为字符串（json）
		db.RequestParams = requestParamsJson
		db.IsOk, _ = strconv.Atoi(requestParams.Get("is_ok"))
	/**
	 * Post请求格式：
	 *	-- 请求头  ：Content-Type: application/x-www-form-urlencoded
	 *	-- 参数格式： key1=value1&key2=value2...
	 */
	case "POST":
		_ = request.ParseForm() // request.Form: 表示一个保存全部参数的 map[key1:[value1]...]
		requestParams = request.Form
		requestParamsBytes, _ := json.Marshal(requestParams)
		requestParamsJson = string(requestParamsBytes) // 将字节切片转换为字符串（json）
		db.RequestParams = requestParamsJson
		db.IsOk, _ = strconv.Atoi(request.FormValue("is_ok"))
	}
	//① 记录请求参数
	chWriterRequestLog := make(chan string) // 写请求日志的通道
	go func(chan string) {
		logHandler := new(lt_log.LogHandler)
		logHandler.Writer("请求&响应", "请求: ", requestParamsJson)
		chWriterRequestLog <- recordId + " : 写入请求参数日志成功"
	}(chWriterRequestLog)
	// 读取通道，释放阻塞
	fmt.Printf("request_log : %s | %s\n", <-chWriterRequestLog, time.Now().Format("2006/01/02 15:04:05"))

	// ②处理结果，下发通知
	chHandler := make(chan string) // 处理参数的通道
	chRetryTimes := make(chan int) //重试次数
	go func() {
		var retryTimeCounts int
		for {
			if retryTimeCounts >= 3 { //最多重试三次
				break
			}
			// 处理逻辑，判断是否重试
			// .....
			retryTimeCounts++
		}
		chHandler <- fmt.Sprintf("%v : 处理主逻辑完成", recordId)
		chRetryTimes <- retryTimeCounts //重试次数写入通道
		// 发送方关闭通道，否则阻塞进程
		close(chHandler)
		close(chRetryTimes)
		// 发送微信通知
		wechat.Sending([]string{"changchao"}, fmt.Sprintf("[go] 处理请求成功，重试次数：%d\n", retryTimeCounts))
	}()
	fmt.Printf("handler_log : %s |重试次数：%v | %s\n", <-chHandler, db.RetryTimes, time.Now().Format("2006/01/02 15:04:05"))
	db.RetryTimes = <-chRetryTimes

	//③ 响应客户端
	if db.IsOk == 1 {
		db.ResponseMsg = "ok"
	} else {
		db.ResponseMsg = "failed"
	}
	responseMap := make(map[string]interface{})
	responseMap["status"] = 1
	responseMap["data"] = map[string]interface{}{"name": "cc"}
	responseMap["msg"] = db.ResponseMsg
	responseMapBytes, _ := json.Marshal(responseMap)
	responseMapJson := string(responseMapBytes)

	//④ 记录响应结果
	chWriterResponseLog := make(chan string) // 写请求日志的通道
	chResponseMsg := make(chan string)       //响应消息的通道
	chResponseMsg <- db.ResponseMsg
	go func(chan string, chan string) {
		logHandler := new(lt_log.LogHandler)
		logHandler.Writer("请求&响应", "响应: ", <-chResponseMsg)
		chWriterResponseLog <- recordId + " : 写入请求参数日志成功"

	}(chWriterResponseLog, chResponseMsg)
	fmt.Printf("response_log : %s | %s\n", db.ResponseMsg, time.Now().Format("2006/01/02 15:04:05")) //读取通道，释放阻塞

	//⑤ 记录请求参数到mysql
	endTime := time.Now()
	endDt := endTime.Format("2006-01-02 15:04:05")    // 结束的时间点
	consumeDuring := endTime.Sub(startTime).Seconds() //一共执行花费的时长：精确值
	db.RequestDt = startDt
	db.ResponseDt = endDt
	db.ExtField1 = strconv.FormatFloat(consumeDuring, 'f', -1, 32)
	db.Insert()

	writer.Header().Set("Content-Type", "application/json") //设置请求头：json格式返回
	_, _ = fmt.Fprintf(writer, responseMapJson)             //发送到客户端响应
}
