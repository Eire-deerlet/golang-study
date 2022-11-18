package main

import (
	"fmt"
)

func main() {

	// fmt.Printf("aaa")
	// fmt.Printf("bbb")
	// a := 100
	// if a > 20 {
	// 	fmt.Printf("a")

	// } else {

	// 	fmt.Printf("bb")
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("i: %v\n", i)
	// }

	// x := [...]int{1, 2, 3}
	// for _, v := range x {
	// 	fmt.Printf("v: %v\n", v)

	// }

	a := 1
	b := 2
	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
	flag := true
	if flag {
		fmt.Printf("flag: %v\n", flag)
	}
	//初始化变量可以声明在布尔表达式里面，注意作用域
	if age := 20; age > 18 {
		fmt.Println("成年了")
		fmt.Printf("age: %v\n", age)
	} else {
		fmt.Println("未成年")
		fmt.Printf("age: %v\n", age)
	}

	// //不能 0 非0 表示真假
	// var i = 1
	// if i {
	// 	fmt.Println("ok")
	// }

}
