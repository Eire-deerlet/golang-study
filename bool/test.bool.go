package main

import "fmt"

func main() {
	var b1 bool = true
	var b2 bool = false
	var b3 = true
	var b4 = false
	b5 := true
	b6 := false

	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("b2: %v\n", b2)
	fmt.Printf("b3: %v\n", b3)
	fmt.Printf("b4: %v\n", b4)
	fmt.Printf("b5: %v\n", b5)
	fmt.Printf("b6: %v\n", b6)

	//if
	age := 18
	gender := "nan"
	if age >= 18 && gender == "nan" {
		fmt.Println("成年了")

	} else {
		fmt.Println("未成年")
	}

	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("i: %v\n", i)
	}

	//不能0和非0 表示 真假
	//i := 1
	//if i {
	//
	//}

}
