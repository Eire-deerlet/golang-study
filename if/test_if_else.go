package main

import "fmt"

func main() {
	//f1()
	f3()
	var name string
	var age int
	var email string

	fmt.Println("请输入name,age,email,用空格分隔")
	fmt.Scan(&name, &age, &email)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("email: %v\n", email)

}

func f3() {

	var number int
	fmt.Println("请输入一个数字")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Println("偶数")
	} else {

		fmt.Println("奇数")
	}

}
func f1() {

	a := 1
	b := 2
	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}
