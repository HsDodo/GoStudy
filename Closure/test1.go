package main

import "fmt"

//闭包学习

func ops(base int) (func(int) int, func(int) int) {

	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	add, sub := ops(10)
	add2, sub2 := ops(100)
	add3, sub3 := ops(10)
	fmt.Println(add(1), sub(2))
	fmt.Println(add(3), sub(4))
	fmt.Println(add2(1), sub2(2))
	fmt.Println(add2(3), sub2(4))
	fmt.Println(add3(1), sub3(2))
	fmt.Println(add3(3), sub3(4))
}
