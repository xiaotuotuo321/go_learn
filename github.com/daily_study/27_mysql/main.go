package main

import (
	"database/sql"
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
var db *sql.DB
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
type user struct{
	id		int
	age 	int
	name   	string
}

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

*/
