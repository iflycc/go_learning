package mysql

import (
	"async_notify/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 定义一个mySql操作句柄，初始化单例模式
var db *sql.DB

// 定义数据表`async_notify_records`结构体
type AsyncNotifyRecords struct {
	Id            int
	RecordId      string //记录id
	RequestParams string //请求参数
	RequestDt     string //请求时间点
	ResponseMsg   string //响应信息
	ResponseDt    string //响应时间点
	RetryTimes    int    //重发次数
	ExtField1     string //备用字段1
	ExtField2     string //备用字段2
	ExtField3     string //备用字段3
	IsOk          int    //是否处理成功，0：失败，1：成功
	CreatedAt     string //创建时间
}

/**
 * db初始化操作
 */
//func init() {
//	// 连接数据库 : "root:root@tcp(127.0.0.1:3306)/test"
//	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",config.Username, config.Password,config.Ip,config.Port,config.DbName)
//	db, err := sql.Open("mysql", connectString)
//	if err != nil { // 查看是否有错误
//		panic(err)
//	}
//	// 设置数据库最大连接数
//	db.SetConnMaxLifetime(100)
//	// 设置数据库最大闲置连接数
//	db.SetMaxIdleConns(10)
//	// 验证链接是否成功
//	if err = db.Ping(); err != nil{
//		panic(fmt.Sprintf("db Ping() failed : %v\n", err))
//	}
//	dt := time.Now().Format("2006/01/02 15:04:05") // 当前的格式化时间显示
//	fmt.Printf("%s: db connection success!\n", dt)
//	fmt.Println(db)
//}

/**
 * db初始化操作
 */
func (asyncNotify *AsyncNotifyRecords) connecting() *sql.DB {
	// 连接数据库 : "root:root@tcp(127.0.0.1:3306)/test"
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Ip, config.Port, config.DbName)
	db, err := sql.Open("mysql", connectString)
	if err != nil { // 查看是否有错误
		panic(err)
	}
	// 设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	// 设置数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	// 验证链接是否成功
	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("db Ping() failed : %v\n", err))
	}
	dt := time.Now().Format("2006/01/02 15:04:05") // 当前的格式化时间显示
	fmt.Printf("%s: db connection success!\n", dt)
	return db
}

/**
 * 查询
 * 		—— Query查询
 */
func (asyncNotify *AsyncNotifyRecords) Query(querySql string) (result []map[string]string) {
	// 链接数据库
	db = asyncNotify.connecting()
	// 延迟关闭数据库
	defer func() {
		_ = db.Close()
	}()
	//试一下
	querySql = "select * from async_notify_records"
	rows, err := db.Query(querySql)
	if err != nil {
		panic(fmt.Sprintf("获取数据错误：%v\n", err))
	}

	//获取列名
	columns, _ := rows.Columns()

	//定义一个切片,长度是字段的个数,切片里面的元素类型是sql.RawBytes
	values := make([]sql.RawBytes, len(columns))
	//定义一个切片,元素类型是interface{} 接口
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		//把sql.RawBytes类型的地址存进去了
		scanArgs[i] = &values[i]
	}
	//获取字段值
	for rows.Next() {
		res := make(map[string]string)
		_ = rows.Scan(scanArgs...)
		for i, col := range values {
			res[columns[i]] = string(col)
		}
		result = append(result, res)
	}
	return
}

/**
 * Insert
 * 		—— 插入记录
 */
func (asyncNotify *AsyncNotifyRecords) Insert() int64 {
	// 链接数据库
	db = asyncNotify.connecting()
	// 延迟关闭数据库
	defer func() {
		_ = db.Close()
	}()
	stmt, err := db.Prepare("INSERT INTO async_notify_records (record_id, request_params, request_dt, response_msg, response_dt, retry_times, ext_field1, ext_field2, ext_field3, is_ok, created_at) VALUES (?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(fmt.Sprintf("insert prepare err:%s\n", err.Error()))
	}
	asyncNotify.CreatedAt = time.Now().Format("2006-01-02 15:04:05") // 格式化写入记录的时间点
	res, err := stmt.Exec(asyncNotify.RecordId, asyncNotify.RequestParams, asyncNotify.RequestDt, asyncNotify.ResponseMsg, asyncNotify.ResponseDt, asyncNotify.RetryTimes, asyncNotify.ExtField1, asyncNotify.ExtField2, asyncNotify.ExtField3, asyncNotify.IsOk, asyncNotify.CreatedAt)
	if err != nil {
		panic(err.Error())
	}
	lastId, _ := res.LastInsertId()
	//affect, _ := res.RowsAffected() // 影响的行数
	return lastId
}
