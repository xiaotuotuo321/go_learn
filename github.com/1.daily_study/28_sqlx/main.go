package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// sqlx

// sqlx: 可以认为是Go语言内置database/sql的超集，在database的基础上提供了一组扩展。

// 1.安装
// go get github.com/jmoiron/sqlx

// 2.基本使用
// 2.1.连接数据库
var db *sqlx.DB
type User struct {
	Name string `json:"name"`
	Age int	`json:"age"`
}

func initDB() (err error){
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil{
		fmt.Printf("connect db failed, err:%v\n", err)
		return
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	return
}

// 2.2.查询单行数据
//func queryRowDemo(){
//	sqlStr := "select id, name, age from user where id = ?"
//	var u User
//	err := db.Get(&u, sqlStr, 2)
//	if err != nil{
//		fmt.Printf("get failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("id:%d name:%v age:%d", u.Id, u.Name, u.Age)
//}

// 2.3.查询多行数据
//func queryMultiRowDemo() {
//	sqlStr := "select id, name, age from user where id > ?"
//	var users []User
//	err := db.Select(&users, sqlStr, 0)
//	if err != nil{
//		fmt.Printf("query failed, err: %v\n", err)
//		return
//	}
//	fmt.Printf("users:%#v\n", users)
//}

// 2.4.插入、更新和删除
// sqlx中的exec方法与原生SQL中的exec使用方法基本一致
// 2.4.1.插入数据
//func insertRowDemo() {
//	sqlStr := "insert into users(name, age) values (?, ?)"
//	ret, err := db.Exec(sqlStr, "小红", 19)
//	if err != nil{
//		fmt.Printf("insert failed, err:%v\n", err)
//		return
//	}
//
//	theId, err := ret.LastInsertId()
//	if err != nil{
//		fmt.Printf("get lastinsertId failed, err: %v\n", err)
//		return
//	}
//	fmt.Printf("insert success, the id is %d.\n", theId)
//}

// 2.4.2.更新数据
//func updateRowDemo() {
//	sqlStr := "update user set age=? where id = ?"
//	ret, err := db.Exec(sqlStr, 39, 6)
//	if err != nil{
//		fmt.Printf("update failed, err:%v\n", err)
//		return
//	}
//	n, err := ret.RowsAffected()
//	if err != nil{
//		fmt.Printf("get RowsAffected failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("update success, affected rows:%d\n", n)
//}

// 2.4.3.删除数据
//func deleteRowDemo() {
//	sqlStr := "delete from user where id = ?"
//	ret, err := db.Exec(sqlStr, 6)
//	if err != nil{
//		fmt.Printf("delete failed, err:%v\n", err)
//		return
//	}
//	n, err := ret.RowsAffected() // 操作影响的行数
//	if err != nil{
//		fmt.Printf("get RowsAffected failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("delete success, affected rows:%d\n", n)
//}

// 2.5.NamedExec
// DB.NamedExec方法用来绑定SQL语句与结构体或map中的同名字段
//func insertUserDemo() (err error){
//	sqlStr := "insert into users (name, age) values (:name, :age)"
//	_, err = db.NamedExec(sqlStr,
//		map[string]interface{}{
//		"name": "小明",
//		"age": 28,
//		})
//	return
//}

// 2.6.NamedQuery
// 支持查询语句
//type user struct {
//	Name string `json:"name"`
//}
//func namedQuery(){
//	sqlStr := "select * from user where name = :name"
//	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name": "小明"})
//	if err != nil{
//		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
//		return
//	}
//	defer rows.Close()
//	for rows.Next(){
//		var u User
//		err := rows.StructScan(&u)
//		if err != nil{
//			fmt.Printf("scan failed, err:%v\n", err)
//			continue
//		}
//		fmt.Printf("user:%v\n", u)
//	}
//
//	u := user{
//		Name: "小明",
//	}
//	// 使用结构体命名查询，根据结构体字段的db tag 进行映射
//	rows, err = db.NamedQuery(sqlStr, u)
//	if err != nil{
//		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
//		return
//	}
//	defer rows.Close()
//	for rows.Next(){
//		var u user
//		err := rows.StructScan(&u)
//		if err != nil{
//			fmt.Printf("scan failed, err:%v\n", err)
//			continue
//		}
//		fmt.Printf("user:%#v\n", u)
//	}
//}

// 2.7.事务操作
// 对于事务操作，可以使用sqlx中提供的db.Beginx()和tx.Exec()方法。
//func transactionDemo2()(err error){
//	tx, err := db.Begin()	// 开启事务
//	if err != nil{
//		fmt.Printf("begin trans failed, err:%v\n", err)
//		return err
//	}
//	defer func(){
//		if p := recover(); p != nil{
//			tx.Rollback()
//			panic(p) // re-throw panic after Rollback
//		} else if err != nil{
//			fmt.Println("rollback")
//			tx.Rollback()	// err is non-nil; don't change it
//		} else {
//			err = tx.Commit() // err is nil; if Commit returns err update err
//			fmt.Println("commit")
//		}
//	}()
//
//	sqlStr1 := "update user set age = 20 where id = ?"
//
//	rs, err := tx.Exec(sqlStr1, 1)
//	if err != nil{
//		return err
//	}
//	n, err := rs.RowsAffected()
//	if err != nil{
//		return err
//	}
//	if n != 1{
//		return errors.New("exec sqlStr1 failed")
//	}
//
//	sqlStr2 := "update user set age=50 where id = ?"
//	rs, err = tx.Exec(sqlStr2, 5)
//	if err != nil{
//		return err
//	}
//	n, err = rs.RowsAffected()
//	if err != nil{
//		return err
//	}
//	if n != 1{
//		return errors.New("exec sqlStr1 failed")
//	}
//	return err
//}

// 3.sqlx.in
// 3.1.表结构
/*
CREATE TABLE `user` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) DEFAULT '',
    `age` INT(11) DEFAULT '0',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
*/

// 3.1.2.bindvars(绑定变量)
// 查询占位符？，在组装SQL语句时，应该使用占位符向数据库发送值，因为可以防止SQL注入攻击，
// 之前的误解是：占位符只是用来在SQL语句中插入值。其实不仅仅是用于参数化，不允许更改SQL语句的结构，

// 不能用来插入表名，不能用来插入列名

// 3.2.自己拼接语句实现批量插入
// BatchInsertUsers 自行构造批量插入的语句
//func BatchInsertUsers(users []*User) error {
//	valueStrings := make([]string, 0, len(users))
//	// 存放values的slice
//	valueArgs := make([]interface{}, 0, len(users) * 2)
//	// 遍历users准备相关数据
//	for _, u := range users{
//		// 此处占位符和插入值的个数对应
//		valueStrings = append(valueStrings, "(?, ?)")
//		valueArgs = append(valueArgs, u.Name)
//		valueArgs = append(valueArgs, u.Age)
//	}
//	// 自行拼接要执行的具体语句
//	stmt := fmt.Sprintf("Insert into user (name, age) values %s", strings.Join(valueStrings, ","))
//	_, err := db.Exec(stmt, valueArgs...)
//	return err
//}
//
////// 3.3.使用sqlx.In实现批量插入：前提是结构体需要实现driver.Valuer 接口：
//func (u User) Value() (driver.Value, error){
//	return []interface{}{u.Name, u.Age}, nil
//}
//
//func BatchInsertUsers2 (users []interface{}) error {
//	query, args, _ := sqlx.In(
//		"insert into user (name, age) values (?), (?), (?)",
//		users..., // 如果arg实现了driver.Valuer, sqlx.In 会通过value()来展开它
//		)
//	fmt.Println(query) // 查看生成的querystring
//	fmt.Println(args) // 查看生成的args
//	_, err := db.Exec(query, args...)
//	return err
//}
//
//// 3.4.使用NamedExec实现批量插入
////func BatchInsertUsers3 使用NamedExec实现批量插入
//func BatchInsertUsers3(users []*User) error {
//	_, err := db.NamedExec("insert into user (name, age) values (:name, :age)", users)
//	return err
//}
//
//func main() {
//	err := initDB()
//	if err != nil{
//		panic(err)
//	}
//	defer db.Close()
//	u1 := User{Name:"小明", Age: 14}
//	u2 := User{Name:"小红", Age: 18}
//	u3 := User{Name:"小黄", Age: 20}
//
//	// 方法1:
//	users := []*User{&u1, &u2, &u3}
//	err = BatchInsertUsers(users)
//	if err != nil{
//		fmt.Printf("BatchInsertUsers failed, err: %v\n", err)
//	}
//
//	// 方法2：
//	users2 := []interface{}{u1, u2, u3}
//	err = BatchInsertUsers2(users2)
//	if err != nil{
//		fmt.Printf("BatchInsertUsers2 failed, err: %v\n", err)
//	}
//
//	// 方法3：
//	users3 := []*User{&u1, &u2, &u3}
//	err = BatchInsertUsers3(users3)
//	if err != nil{
//		fmt.Printf("BatchInsertUsers3 failed, err: %v\n", err)
//	}
//}

// 3.5. sqlx.In 实现了一个in 方法，且可以按照可排序字段排序返回 in函数 和 find_in_set 函数
// select * from user where id in (1, 2, 3)
// select * from user where id in (1, 2, 3) order by find_in_set(id, '3, 2, 1')

// 3.5.1. in函数使用 querybyids 根据给定ID查询
//func QueryByIds(ids []int) (users []User, err error){
//	// 动态填充ID
//	query, args, err := sqlx.In("select name, age from user where id in ()", ids)
//	if err != nil{
//		return
//	}
//	// sqlx.in 返回`？` bindvar 的查询语句，我们使用rebind()重新绑定它
//	query = db.Rebind(query)
//
//	err = db.Select(&users, query, args...)
//	return
//}
//
//// 3.5.2. in查询和find_in_set函数
//// 查询ID在给定ID的集合的数据并维持给定ID集合的顺序
//// QueryAndOrderByIds 按照指定的ID查找数据，并按照一定的ID顺序返回
//
//func QueryAndOrderByIds(ids []int) (users []User, err error){
//	// 动态填充ID
//	strIDs := make([]string, 0, len(ids))
//	for _, id := range ids{
//		strIDs = append(strIDs, fmt.Sprintf("%d", id))
//	}
//
//	query, args, err := sqlx.In("select name, age from user where id in (?) order by find_in_set (id, ?)", ids, strings.Join(strIDs, ","))
//
//	query = db.Rebind(query)
//
//	err = db.Select(&users, query, args...)
//	return
//}


