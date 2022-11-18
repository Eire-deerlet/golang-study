package main

import "fmt"

func main() {
	f1()
	f2()
	f4()
}

func f1() {

	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}

}

func f2() {

	i := 1 //放在外面
	for i < 10 {
		fmt.Printf("i: %v\n", i)
		i++ //外面
	}

}

func f4() {
	for {

		fmt.Println("我一直在执行")
	}
}
