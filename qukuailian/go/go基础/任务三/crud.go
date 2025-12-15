package main

import (
	"database/sql" // 导入了但没有使用
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 匿名导入，注册MySQL驱动
	"github.com/jmoiron/sqlx"
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

// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
func queryStudent(age int) {
	// 构造SQL语句
	querySQL := "SELECT * FROM students WHERE age > ?"
	// 执行SQL语句
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/mydatabase")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query(querySQL, age)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var students []Student
	for rows.Next() {
		var student Student
		err = rows.Scan(&student.id, &student.name, &student.age, &student.grade)
		if err != nil {
			panic(err)
		}
		students = append(students, student)
	}
	fmt.Println(students)
}

// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
func updateStudent(name string, grade string) {
	// 构造SQL语句
	updateSQL := "UPDATE students SET grade = ? WHERE name = ?"
	// 执行SQL语句
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/mydatabase")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmt, err := db.Prepare(updateSQL)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(grade, name)
	if err != nil {
		panic(err)
	}
	fmt.Println("Update student success!")
}

// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
func deleteStudent(age int) {
	// 构造SQL语句
	deleteSQL := "DELETE FROM students WHERE age < ?"
	// 执行SQL语句
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/mydatabase")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmt, err := db.Prepare(deleteSQL)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(age)
	if err != nil {
		panic(err)
	}
	fmt.Println("Delete student success!")
}

// 题目2：实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

type Book struct {
	Id     int
	Title  string
	Author string
	Price  float64
}

func queryBook() {
	db, err := sqlx.Connect("mysql", "root:123456@tcp(localhost:3306)/mydatabase")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var books []Book
	err = db.Select(&books, "SELECT id, title, author, price FROM books WHERE price > ?", 70)
	if err != nil {
		panic(err)
	}
	fmt.Println(books)
}

func main() {
	// insertStudent("张三1", 19, "sdds")

	// insertStudent("李四", 18, "sdds")

	// insertStudent("王五", 98, "sdds")

	// insertStudent("赵柳", 18, "sdds")

	// insertStudent("流星", 31, "sdds")

	// queryStudent(18)

	// updateStudent("张三1", "四年级")
	// deleteStudent(19)

	queryBook()
}
