package main

import "fmt"

//email := "aaaa.com"

var email = "qq.com"

func main() {
	// var name string = "tome"
	// var age int = 20
	// var email string = "4444"

	// var (
	// 	name  string
	// 	age   int
	// 	email string
	// )

	// name = "tom"
	// age = 20
	// email = "263"

	//类型推断
	/* 	var name = "tome"
	   	var age = true
	   	var email = "@qq.com" */

	//批量初始化

	//	var name, age, email = "tom", 26, "163.com"

	//短变量 只适用函数内部
	// name := "tom"
	// age := 15

	//匿名变量
	name, age := getNameAndAge()
	email := "gogler"

	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("email: %v\n", email)
}

func getNameAndAge() (string, int) {

	return "tom", 20
}
