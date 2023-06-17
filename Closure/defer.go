package main

import "fmt"

//	func main() {
//		var whatever [5]int
//
//		for i := range whatever {
//			defer fmt.Println(i)
//		}
//
// }
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	add := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(add(i))
	}
}
