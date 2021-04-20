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
 * @param logContent string  日志内容
 */
func (logH *LogHandler) Writer(logName, logPrefix, logContent string) {
	dt := time.Now().Format("20060102") // 当前的日期
	//pwd, err := os.Getwd() //获取当前路径的目录地址
	//if err != nil{
	//	panic(fmt.Sprintf("fetch base dir failed, err : %v", err))
	//}
	logPath := fmt.Sprintf("%s/log_files/%s", baseDir, dt) //日志目标目录
	// 判断目录是否存档
	if _, oErr := os.Stat(logPath); oErr != nil { //目标目录不存在，则自动创建
		_ = os.Mkdir(logPath, 0777)
	}
	logFileCh, err := os.OpenFile(fmt.Sprintf("%s/log_files/%s/%s", baseDir, dt, logName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644) // 读写权限
	if err != nil {
		panic(fmt.Sprintf("open lt_log file err :%v", err)) // 当打开日志文件失败时，抛出panic异常
	}
	log.SetOutput(logFileCh)             // 设置文件句柄
	log.SetPrefix("[" + logPrefix + "]") // 设置日志前缀
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	log.Println(logContent) // 写入日志，并且换行
}
