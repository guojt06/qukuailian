package main

import (
	"database/sql" // 导入了但没有使用
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 匿名导入，注册MySQL驱动
)

// 题目1：基本CRUD操作
// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

//新建一个结构体Student,包换id,name,age,grade字段

type Student struct {
	id    int
	name  string
	age   int
	grade string
}

// 插入一条新记录
func insertStudent(name string, age int, grade string) {
	// 构造SQL语句
	insertSQL := "INSERT INTO students (name, age, grade) VALUES (?, ?, ?)"
	// 执行SQL语句
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/mydatabase")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmt, err := db.Prepare(insertSQL)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, age, grade)
	if err != nil {
		panic(err)
	}
	fmt.Println("Insert student success!")
}

func main() {
	insertStudent("张三", 12, "sdds")
}
