package main

import (
	"fmt"
	"strings"
)

func main() {
	/* var s string = "hello world"
	var s1 = "hello"
	s3 := "hello.go"
	fmt.Printf("s: %v\n", s)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s3: %v\n", s3)

	//反引号
	s4 := `
	go1
	g02
	go3
	`
	fmt.Printf("s4: %v\n", s4) */

	//字符串连接
	// s1 := "tom"
	// s2 := "jerry"
	// msg := s1 + s2
	// fmt.Printf("msg: %v\n", msg)

	// msg2 := fmt.Sprintf("name=%s,name2=%s", s1, s2)
	// fmt.Printf("msg2: %v\n", msg2)

	// name := "tome"
	// age := "20"
	// s := strings.Join([]string{name, age}, "++++++")
	// fmt.Printf("s: %v\n", s)

	// var buffer bytes.Buffer
	// buffer.WriteString("tom")
	// buffer.WriteString("+++++++")
	// buffer.WriteString("jerry")
	// fmt.Printf("buffer.String(): %v\n", buffer.String())

	//转义字符
	s := "hello"
	print(s + "\n")
	print(s + "\r")

	s2 := "c:\\test\\work\\"
	fmt.Printf("s2: %v\n", s2)

	//字符串切片
	s3 := "hello world"
	a := 8
	b := 10
	fmt.Printf("s3[a] : %c \n", s3[a])
	fmt.Printf("s3[a:b]: %v \n", s3[a:b]) //区间

	fmt.Printf("s3[a:]: %v\n", s3[a:])

	fmt.Printf("s3[:b]: %v\n", s3[:b])

	fmt.Printf("len(s3): %v\n", len(s3)) //长度

	fmt.Printf("strings.Split(s3,\"\"): %v\n", strings.Split(s3, " "))                 //转换数组
	fmt.Printf("strings.Contains(s3, \"hello\"): %v\n", strings.Contains(s3, "hello")) //判断是否存在

	fmt.Printf("strings.ToLower(s3): %v\n", strings.ToLower(s3))                         //小写
	fmt.Printf("strings.ToUpper(s3): %v\n", strings.ToUpper(s3))                         //大写
	fmt.Printf("strings.HasPrefix(s3, \"HELLO\"): %v\n", strings.HasPrefix(s3, "HELLO")) //打头
	fmt.Printf("strings.HasSuffix(s3, \"world\"): %v\n", strings.HasSuffix(s3, "world")) //结尾
	fmt.Printf("strings.Index(s3, \"w\"): %v\n", strings.Index(s3, "w"))                 //索引位置

	fmt.Printf("strings.LastIndex(s3, \"o\"): %v\n", strings.LastIndex(s3, "o")) //最后一个位置
}
