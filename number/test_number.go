package main

import (
	"fmt"
	"math"
)

func main() {
	/* var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var ui8 uint8
	var ui16 uint16
	var ui32 int32
	var ui64 int64
	fmt.Printf("i8: %v\n", i8) */

	//十进制
	var a int = 10
	fmt.Printf("%d \n", a) //10
	fmt.Printf("%b \n", a) //%b 表示二进制
	//八进制
	var b int = 077
	fmt.Printf("%o \n", b) //77

	//十六进制 以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) //小写ff
	fmt.Printf("%X \n", c) //大写 FF

	//浮点类型
	fmt.Printf("%f \n", math.Pi)   //默认
	fmt.Printf("%.2f \n", math.Pi) //两位

	//复数

}
