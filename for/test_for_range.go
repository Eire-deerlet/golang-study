package main

import "fmt"

func main() {
	//var a = [5]int{1, 2, 3, 4, 5}
	//
	//for i, v := range a {
	//	fmt.Printf("i :%d ,v : %\n", i, v)
	//}

	//f3()
	f5()
	f6()
	f7()
}
func f3() {
	var a = [...]int{1, 2, 3}
	for i, v := range a {
		//fmt.Printf("i : %v\n ", i)
		//fmt.Printf("v : %v\n ", v)

		fmt.Printf("%v:%v", i, v)
	}
}

func f5() {
	var s = []int{1, 2, 3}
	for _, v := range s {
		println(v)
		fmt.Printf(" v : %v \n ", v)
	}
}

func f6() {

	m := make(map[string]string, 0)
	m["name"] = "tom"
	m["age"] = "20"
	m["email"] = "263@qq.com"
	for key, value := range m {
		fmt.Printf("%v :%v \n", key, value)
	}
}

func f7() {
	s := "hello world"
	for _, i2 := range s {
		fmt.Printf("i2 : %c \n", i2)
	}
}
