package lt_log

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var baseDir string

/**
 * 初始化函数，确定项目文件的基础路径
 */
func init() {
	pwd, err := os.Getwd() //获取当前路径的目录地址
	if err != nil {
		panic(fmt.Sprintf("fetch base dir failed, err : %v", err))
	}
	pwd = strings.ReplaceAll(pwd, "\\", "/")
	lastEofIndex := strings.LastIndex(pwd, "/")
	baseDir = pwd[0:lastEofIndex]
}

/**
 * 日志处理结构体
 */
type LogHandler struct {
}

/**
 * 写日志
 * @param logName string  日志名称
 * @param logPrefix string  日志前缀
 * @param logId string  日志id（单次请求的唯一标识）
 * @param logContent string  日志内容
 */
func (logH *LogHandler) Writer(logName, logPrefix, logId, logContent string) {
	dt := time.Now().Format("20060102") // 当前的日期
	//pwd, err := os.Getwd() //获取当前路径的目录地址
	//if err != nil{
	//	panic(fmt.Sprintf("fetch base dir failed, err : %v", err))
	//}
	//return

	logPath := fmt.Sprintf("%s/log_files/%s", baseDir, dt) //日志目标目录
	// 判断目录是否存档
	if _, oErr := os.Stat(logPath); oErr != nil { //目标目录不存在，则自动创建
		mkErr := os.MkdirAll(logPath, 0777) // 创建多级目录
		if mkErr != nil {
			log.Fatal("mkErr : ", mkErr)
		}
	}
	logFileCh, err := os.OpenFile(logPath+"/"+logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644) // 读写权限
	if err != nil {
		panic(fmt.Sprintf("open lt_log file err :%v", err)) // 当打开日志文件失败时，抛出panic异常
	}
	log.SetOutput(logFileCh)                     // 设置文件句柄
	log.SetPrefix("[" + logPrefix + logId + "]") // 设置日志前缀
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	log.Println(logContent) // 写入日志，并且换行
}
