package main

import "fmt"

func main() {
	grade := 'B'

	switch grade {
	case 'A':
		fmt.Println("優秀")
	case 'B':
		fmt.Println("良好")
	case 'C':

		fmt.Println("一般")

	default:
		fmt.Println("差")
	}

	f2()
	f3()
	f4()
}

func f2() {
	day := 100
	switch day {

	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 7, 6:
		fmt.Println("休息日")

	default:
		fmt.Println("非法輸入")
	}

}
func f3() {
	score := 80
	switch {
	case score >= 60 && score < 80:
		fmt.Println("及格")
	case score >= 80 && score <= 100:
		fmt.Println("优秀")
	default:
		fmt.Println("其他")
	}

}

// fallthrough 穿透
func f4() {

	num := 100
	switch num {
	case 100:
		fmt.Println("100")
		fallthrough
	case 200:
		fmt.Println("200")
	case 300:
		fmt.Println("300")
	}
}
