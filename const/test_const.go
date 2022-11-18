package main

import "fmt"

func main() {

	const PI float32 = 3.14

	const PI2 = 3.1415

	fmt.Printf("PI: %v\n", PI)

	const (
		a = 100
		b = 200
	)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	const w, h = 55, 66
	fmt.Printf("w: %v\n", w)
	fmt.Printf("h: %v\n", h)

	////iota
	//const (
	//	a1 = iota
	//	a2 = iota
	//	a3 = iota
	//)
	//fmt.Printf("a1: %v\n", a1)
	//fmt.Printf("a2: %v\n", a2)
	//fmt.Printf("a3: %v\n", a3)

	////跳过
	//const (
	//	a1 = iota
	//
	//	_
	//	a2 = iota
	//)

	//fmt.Printf("a1: %v\n", a1)
	//fmt.Printf("a2: %v\n", a2)

	const (
		a1 = iota
		a2 = 100
		a3 = iota
	)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)

}
