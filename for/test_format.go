package main

import "fmt"

type WebSite struct {
	name string
}

func main() {
	/* fmt.Printf("hello : %v \n", "helloooo")

	name := WebSite{name: "tom"}
	fmt.Printf("name: %v\n", name)
	fmt.Printf("name: %#v\n", name)
	fmt.Printf("name: %T\n", name)
	fmt.Printf("name: %%\n", name) //% 自己
	b := true
	fmt.Printf("b: %t\n", b) */

	i := 8
	fmt.Printf("i: %v\n", i)
	fmt.Printf("i: %b\n", i) //二进制
	i = 96
	fmt.Printf("i: %c\n", i)
	fmt.Printf("i: %x\n", i)

	//指针地址
	x := 100
	p := &x
	fmt.Printf("p: %v\n", p)
}
