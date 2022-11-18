package main

import "fmt"

func main() {
	// a := 100
	// b := 20
	// c := a + b
	// fmt.Printf("c: %v\n", c)

	// c = a - b
	// fmt.Printf("c: %v\n", c)
	// c = a * b
	// fmt.Printf("c: %v\n", c)
	// c = a / b
	// fmt.Printf("c: %v\n", c)
	// c = a % b
	// fmt.Printf("c: %v\n", c)

	//++
	// a := 100
	// a++
	// fmt.Printf("a: %v\n", a)
	// a--
	// fmt.Printf("a: %v\n", a)
	//关系运算
	// a := 10
	// b := 5
	// r := a == b //:= 赋值
	// fmt.Printf("r: %v\n", r)
	// r = a > b
	// fmt.Printf("r: %v\n", r)
	// r = a < b
	// fmt.Printf("r: %v\n", r)
	// r = a >= b
	// fmt.Printf("r: %v\n", r)
	// r = a <= b
	// fmt.Printf("r: %v\n", r)
	// r = a != b
	// fmt.Printf("r: %v\n", r)
	// //逻辑运输
	// a := true
	// b := false
	// r := a && b
	// fmt.Printf("r: %v\n", r)
	// r = a || b
	// fmt.Printf("r: %v\n", r)
	// fmt.Printf("a: %v\n", !a) //取反

	//位运算
	// a := 4
	// fmt.Printf("a: %b\n", a) //0100
	// b := 8
	// fmt.Printf("b: %b\n", b) //1000
	// r := a & b
	// fmt.Printf("r: %v\n", r) //0 两个二进位均为1时，结果位才为1 ，否则为0。参与运算的数以补码方式出现。

	// r = a | b
	// fmt.Printf("r: %v\n", r) //12 按位或运算符“|”是双目运算符。 其功能是参与运算的两数各对应的二进位相或。只要对应的二个二进位有一个为1时，结果位就为1
	// r = a ^ b
	// fmt.Printf("r: %v\n", r) //12 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。参与运算数仍以补码出现

	// r = a << 2
	// fmt.Printf("r: %v\n", r) //16 0100
	// r = a >> 2
	// fmt.Printf("r: %v\n", r) //1

	a := 100
	a = 200
	fmt.Printf("a: %v\n", a)
	a = ageA()
	fmt.Printf("a: %v\n", a)
	a += 100
	fmt.Printf("a: %v\n", a)

}
func ageA() int {

	return 100
}
