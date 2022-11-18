package main

import "fmt"

func main() {
	var score int
	fmt.Println("请输入考试分数")
	fmt.Scan(&score)

	if score >= 90 && score <= 100 {
		fmt.Println("奖励小汽车")
	} else if score >= 60 && score <= 89 {
		fmt.Println("奖励零食")
	} else if score <= 59 && score >= 0 {
		fmt.Println("奖励 铁拳")
	} else {
		fmt.Println("分数不对")
	}
	f2()

}

func f2() {

	gender := "woman"
	age := 20
	if gender == "man" {

		if age >= 20 {
			fmt.Println("成年男性")
		}
	} else {
		if age >= 18 {
			fmt.Println("成年女性")
		}
	}

}
