package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_learn/github.com/daily_study/27_mysql/errnum"
	"io/ioutil"
	"net/http"
)

// go 操作MySQL

// 1.连接
// go语言中的database/sql 包提供了保证sql或类sql数据库的泛用接口，并不提供具体的数据库驱动。使用database/sql包时必须注入（至少）一个数据库驱动
// 安装依赖：go get -u github.com/go-sql-driver/mysql

// 使用MySQL驱动
// func Open(driverName, dataSourceName string) (*DB, error)
// open打开一个driverName指定的数据库，dataSourceName指定数据源，一般至少包括数据库文件和其他连接必要的信息

//func main() {
//	// DSN : Data Source Name
//	dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
//	db, err := sql.Open("mysql", dsn)
//	if err != nil{
//		panic(err)
//	}
//	defer db.Close()	// 这行代码要写在err判断的下面
//}

// 1.1.初始化连接
// open函数可能只是验证其参数格式是否正确，实际上并不创建与数据库的连接。如果要减产数据源的名称是否真实有效，应该调用Ping方法
// 返回的DB对象可以安全地被多个goroutine并发使用，并且维护其自己的空闲连接池。所以, open函数应该仅被调用一次，很少需要关闭这个对象。

// 定义一个全局对象db
//var db *sql.DB
//
//// 定义一个初始化数据库的函数
//func initDB() (err error) {
//	// DSN: Data Source Name
//	dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
//	// 不会校验账号密码是否正确
//	// !! 这里不要使用:=, 这里是给全局变量赋值，然后在main函数中使用全局变量db
//	db, err = sql.Open("mysql", dsn)
//	if err != nil{
//		return err
//	}
//	// 尝试与数据库建立连接（校验dsn是否正确）
//	err = db.Ping()
//	if err != nil{
//		return err
//	}
//	return nil
//}
//
//func main() {
//	err := initDB() // 调用输出化数据库的函数
//	if err != nil{
//		fmt.Printf("init db failed, err:%v\n", err)
//		return
//	}
//}
// sql.DB 是表示连接的数据库对象（结构体实例），它保存了连接数据库相关的所有信息。它内部维护着一个具有零到多个底层连接的连接池，它可以安全地被多个goroutine同时使用

// 1.2. SetMaxOpenConns func (db *DB) SetMaxOpenConns(n int)
// SetMaxOpenConns 设置与数据库建立连接的最大数目。如果n大于0且小于最大闲置连接数,会将最大闲置连接数减小到匹配最大开启连接数的限制。如果n<=0,不会限制最大开启连接数，默认为0（无限制）

// 1.3. setMaxIdleConns	func (db *DB) SetMaxIdleConns(n int)
// SetMaxIdleConns设置连接池中的最大闲置连接数。如果n大于最大开启连接数，则新的最大闲置连接数会减小到匹配最大开启连接数的限制，如果n<=0，不会保留闲置连接。

// 2.CRUD
// 2.1.建库建表
/*
CREATE TABLE `user` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) DEFAULT '',
    `age` INT(11) DEFAULT '0',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
*/

// 2.2.为了方便查询，事先定义好一个结构体来存储user表的数据。
//type user struct{
//	id		int
//	age 	int
//	name   	string
//}

// 2.2.1.单行查询，db.QueryRow()执行一次查询，并期望返回最多一行结果（即Row）。QueryRow总是返回非nil的值，知道返回值得Scan方法被调用时，才会返回被延迟的错误，比如：未找到结果
//func (db *DB) QueryRow(query string, args ...interface{}) *Row
//func queryRowDemo() {
//	sqlStr := "select id, name, age from user where id=?"
//	var u user
//	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库连接不会被释放
//	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.age, &u.name)
//	if err != nil{
//		fmt.Printf("scan failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("id:%d name:%s age:%d", u.id, u.name, u.age)
//}

// 2.2.2.多行查询 多行查询db.Query()执行一次查询，返回多行结果(即Rows), 一般用于执行select命令。参数args表示query中的占位参数。
//func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
//func queryMultiRowDemo() {
//	sqlStr := "select id, age, name from user where id > ?"
//	rows, err := db.Query(sqlStr, 0)
//	if err != nil{
//		fmt.Printf("query failed, err:%v\n", err)
//		return
//	}
//	// 非常重要：关闭rows释放持有的数据库连接
//	defer rows.Close()
//
//	// 循环读取结果集中的数据
//	for rows.Next() {
//		var u user
//		err := rows.Scan(&u.id, &u.name, &u.age)
//		if err != nil{
//			fmt.Printf("scan failed, err:%v\n", err)
//			return
//		}
//		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
//	}
//}

// 2.2.3.插入数据：插入、更新和删除操作都是用Exec方法。
// func (db *DB) Exec(query string, args ...interface{}) (Result, error)
// Exec执行一次命令(包括查询、删除、更新、插入等)，返回的Result是对已执行的sql命令的总结。参数args表示query中的占位参数。
// 插入数据
//func insertRowDemo(){
//	sqlStr := "insert into user(name, age) values (?, ?)"
//	ret, err := db.Exec(sqlStr, "王五", 38)
//	if err != nil{
//		fmt.Printf("insert failed, err:%v\n", err)
//		return
//	}
//	theID, err := ret.LastInsertId()	// 新插入数据的ID
//	if err != nil{
//		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("insert success, the id is %d.\n", theID)
//}

// 2.2.4.更新数据
//func updateRowDemo() {
//	sqlStr := "update user set age = ? where id = ?"
//	ret, err := db.Exec(sqlStr, 39, 3)
//	if err != nil{
//		fmt.Printf("update failed, err:%v\n", err)
//		return
//	}
//	n, err := ret.RowsAffected()  // 操作影响的行数
//	if err != nil{
//		fmt.Printf("get RowsAffected failed, err: %v\n", err)
//		return
//	}
//	fmt.Printf("update success, affected rows:%d\n", n)
//}

// 2.2.5.删除数据
//func deleteRowDemo() {
//	sqlStr := "delete from user where id = ?"
//	ret, err := db.Exec(sqlStr, 3)
//	if err != nil{
//		fmt.Printf("delete failed, err:%v\n", err)
//		return
//	}
//	n, err := ret.RowsAffected() // 操作影响到的行数
//	if err != nil{
//		fmt.Printf("get RowsAffected failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("delete success, affected rows:%d\n", n)
//}

// 3.mysql预处理
/*
普通sql语句执行过程：
	1.客户端对SQL语句进行占位符替换得到完整的SQL语句
	2.客户端发送完整SQL语句到MySQL服务端
	3.MySQL服务端执行完整的SQL语句并将结果返回给客户端

预处理进行过程：
	1.把SQL语句分成两部分，命令部分和数据部分
	2.先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理
	3.然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换
	4.MySQL服务端执行完整的语句并将结果返回给客户端

为什么要预处理？
	1.优化MySQL服务器重复执行SQL方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本
	2.避免SQL注入问题
*/

// 3.1.go实现MySQL预处理
// database/sql中使用prepare方法来实现预处理操作	func (db *DB) Prepare (query string) (*Stmt, error)
// prepare方法会先将SQL语句发送给MySQL服务端，返回一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令
// 查询操作的预处理示例代码
//func prepareQueryDemo() {
//	sqlStr := "select id, name, age from user where id > ?"
//	stmt, err := db.Prepare(sqlStr)
//	if err != nil{
//		fmt.Printf("prepare failed, err:%v\n", err)
//		return
//	}
//	defer stmt.Close()
//	rows, err := stmt.Query(0)
//	if err != nil{
//		fmt.Printf("query failed, err:%v", err)
//		return
//	}
//	defer rows.Close()
//	// 循环读取结果集中的数据
//	for rows.Next(){
//		var u user
//		err := rows.Scan(&u.id, &u.age, &u.name)
//		if err != nil{
//			fmt.Printf("sacn failed, err: %v\n", err)
//			return
//		}
//		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
//	}
//}

// 3.2.插入、更新和删除操作的预处理十分相似，这里以插入操作的预处理为例
//func prepareInsertDemo() {
//	sqlStr := "insert into user(name, age) values (?, ?)"
//	stmt, err := db.Prepare(sqlStr)
//	if err != nil{
//		fmt.Printf("prepare failed, err:%v\n", err)
//		return
//	}
//	defer stmt.Close()
//	_, err = stmt.Exec("小王子", 18)
//	if err != nil{
//		fmt.Printf("insert failed, err:%v\n", err)
//		return
//	}
//	_, err = stmt.Exec("haha", 19)
//	if err != nil{
//		fmt.Printf("insert failed, err:%v\n", err)
//		return
//	}
//	fmt.Println("insert success.")
//}

// 3.3.SQL注入问题：  ***任何时候都不应该自己拼接SQL语句
// SQL注入的例子
//func sqlInjectDemo(name string){
//	sqlStr := fmt.Sprintf("select id, name, age, from user where name='%s'", name)
//	fmt.Printf("SQL:%s\n", sqlStr)
//	var u user
//	err := db.QueryRow(sqlStr).Scan(&u.id, &u.age, &u.name)
//	if err != nil{
//		fmt.Printf("exec failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("user:%#v\n", u)
//}

// 下面的字符串都能引起上面的方法SQL注入问题
/*
sqlInjectDemo("xxx" or 1=1#)
sqlInjectDemo("xxx" union select * from user #)
sqlInjectDemo("xxx" and (select count(*) from user) < 10#)
*/

// 4.go 实现MySQL事务
/*
什么是事务？
事务：一个最小的不可再分的工作单元；通常一个事务对应的是一个完整的业务（比如银行账户转账业务，该业务就是一个最小的工作单元，同时这个完整的业务西药执行多次的DML（
insert, update, delete）语句等共同联合完成，比如由A转账给B，这里面就需要执行两次update操作）

在MySQL中只有使用了innodb数据库引擎的数据库或表才支持事务，事务处理可以用来维护数据的完整性，保证成批的SQL语句要么全部执行，要么全部不执行

事务的ACID：
通常事务必须满足四个条件（ACID）：原子性，一致性，隔离性，持久性
原子性：一个事务中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节。事务在执行过程中发生错误，会被回滚，到书屋开始前的状态，就像这个事务从来没有执行过一样
一致性：在事务开始之前和事务结束之后，数据库的完整性没有被破坏，这表示写入的资料必须完全复合物所有的预设规则，这包含资料的准确度、串联性以及后续数据库可以自发性地完成预定的工作
隔离性：数据库允许多个事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据的不一致。事务隔离分为不同级别，包括读未提交，读提交，可重复读，串行化
持久性：事务处理结束后，对数据的修改就是永久的，即便系统故障也不会丢失
*/

// 4.1.事务的相关方法：
// 开始事务： func (db *DB) Begin() (*Tx, error)
// 提交事务：func (db *DB) Commit() error
// 回滚事务：func (db *DB) RollBack() error

// 4.2.事务示例
//func transactionDemo() {
//	tx, err := db.Begin()	// 开启事务
//	if err != nil{
//		if tx != nil{
//			tx.Rollback()	// 回滚
//		}
//		fmt.Printf("begin trans failed, err: %v\n", err)
//		return
//	}
//	sqlStr1 := "Update user set age = 20 where id = ?"
//	ret1, err := tx.Exec(sqlStr1, 2)
//	if err != nil{
//		tx.Rollback()
//		fmt.Printf("exec sql failed, err:%v\n", err)
//		return
//	}
//
//	affRow1, err := ret1.RowsAffected()
//	if err != nil{
//		tx.Rollback()	// 回滚
//		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
//		return
//	}
//
//	sqlStr2 := "Update user set age = 40 where id = ?"
//	ret2, err := tx.Exec(sqlStr2, 3)
//	if err != nil{
//		tx.Rollback()
//		fmt.Printf("exec sql failed, err:%v\n", err)
//		return
//	}
//
//	affRow2, err := ret2.RowsAffected()
//	if err != nil{
//		tx.Rollback()	// 回滚
//		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
//		return
//	}
//	fmt.Println(affRow1, affRow2)
//	if affRow1 == 1 && affRow2 == 1{
//		fmt.Println("事务提交")
//		tx.Commit()
//	} else {
//		tx.Rollback()
//		fmt.Println("事务回滚")
//	}
//	fmt.Println("exec trans success!")
//}

// 5.练习题： 结合net/http和database/sql实现一个使用MySQL存储用户信息的注册及登陆的简易web程序。

type userRegisterInfo struct {
	Name string	`json:"name"`
	Age int	`json:"age"`
	Passwd string	`json:"passwd"`
	Gander string	`json:"gander"`
	ConformPasswd string `json:"conform_passwd"`
}

type userLoginInfo struct {
	Name string `json:"name"`
	Passwd string `json:"passwd"`
}

var db *sql.DB
const dsn = "root:123456@tcp(127.0.0.1:3306)/sql_test"

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello A用户！")
}

func registerHandler(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	var answer = `{"status": "ok"}`
	userInfo := &userRegisterInfo{}
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	//r.ParseForm()
	//fmt.Println(r.PostForm)
	//fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("passwd"), r.PostForm.Get("age"), r.PostForm.Get("gander"))
	// 2. 请求类型是application/json时从r.Body读取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	err = json.Unmarshal(b, userInfo)
	if err != nil{
		fmt.Printf("解析参数时出错：%v\n", err.Error())
		return
	}
	err = Register(userInfo)
	if err != nil{
		err = errnum.New(&errnum.Er{40001, "用户注册时出错"}, err)
		answer = `{"status": "error"}`
		w.Write([]byte(answer + err.Error()))
		return
	} else {
		w.Write([]byte(answer))
		return
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	var answer = `{"status": "ok", "data": "登录成功！"}`
	userInfo := &userLoginInfo{}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	err = json.Unmarshal(b, userInfo)
	if err != nil{
		fmt.Printf("解析参数时出错：%v\n", err.Error())
		return
	}

	err = Login(userInfo)
	if err != nil{
		err = errnum.New(&errnum.Er{40002, "用户登录时出错"}, err)
		answer = `{"status": "error"}`
		w.Write([]byte(answer + err.Error()))
		return
	} else {
		w.Write([]byte(answer))
		return
	}


	w.Write([]byte(answer))
}

func Register(userInfo *userRegisterInfo) (err error){
	// 校验用户输入的两次密码是否一致
	if userInfo.Passwd != userInfo.ConformPasswd{
		err = errnum.New(&errnum.Er{50001, "请保证两次输入密码的一致性"}, nil)
		return err
	}

	// 连接数据库
	db, err = sql.Open("mysql", dsn)
	if err != nil{
		err = errnum.New(&errnum.Er{50002, "打开数据库连接失败！"}, err)
		return err
	}
	// 测试连接数据库
	err = db.Ping()
	if err != nil{
		err = errnum.New(&errnum.Er{50003, "连接数据库失败"}, err)
		return err
	}

	// 查询当前用户是否已经存在
	existsSqlStr := "select id from users where name = ?"
	stmt, err := db.Prepare(existsSqlStr)
	if err != nil{
		err = errnum.New(&errnum.Er{50004, "预处理查询SQL失败"}, err)
		return err
	}

	defer stmt.Close()
	rows, err := stmt.Query(userInfo.Name)
	if err != nil{
		err = errnum.New(&errnum.Er{50005, "查询失败"}, err)
		return err
	}

	if rows.Next(){
		err = errnum.New(&errnum.Er{50006, "注册用户已存在"}, nil)
		return err
	}

	// 如果注册用户不存在，则插入用户信息
	insertSqlStr := "insert into users (name, passwd, age, gander) values (?, ?, ?, ?)"
	stmt, err = db.Prepare(insertSqlStr)
	if err != nil{
		err = errnum.New(&errnum.Er{50007, "预处理插入SQL失败"}, err)
		return err
	}
	_, err = stmt.Exec(userInfo.Name, userInfo.Passwd, userInfo.Age, userInfo.Gander)
	if err != nil{
		err = errnum.New(&errnum.Er{50008, "插入用户信息失败"}, err)
		return err
	}

	return nil
}

func Login(userInfo * userLoginInfo) (err error){
	var u userLoginInfo
	// 创建数据库连接
	db, err = sql.Open("mysql", dsn)
	if err != nil{
		err = errnum.New(&errnum.Er{50002, "打开数据库连接失败！"}, err)
		return err
	}

	err = db.Ping()
	if err != nil{
		err = errnum.New(&errnum.Er{50003, "连接数据库失败"}, err)
		return err
	}

	// 校验查询用户是否存在；如果存在，则校验用户的密码是否正确
	searchSql := "select name, passwd from users where name = ?"

	stmt, err := db.Prepare(searchSql)
	if err != nil{
		err = errnum.New(&errnum.Er{50004, "预处理查询SQL失败"}, err)
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(userInfo.Name)
	if err != nil{
		err = errnum.New(&errnum.Er{50005, "查询失败"}, err)
		return err
	}
	if !rows.Next(){
		err = errnum.New(&errnum.Er{51001, "登录用户不存在，请先注册"}, nil)
		return err
	}
	err = rows.Scan(&u.Name, &u.Passwd)
	if err != nil{
		err = errnum.New(&errnum.Er{51002, "解析单行数据出错"}, err)
		return err
	}

	if u.Passwd != userInfo.Passwd{
		err = errnum.New(&errnum.Er{51003, "用户输入的密码有误"}, err)
		return err
	}

	return nil
}

func main() {
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	err := http.ListenAndServe(":9999", nil)
	if err != nil{
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}