package main

const (
	i = 1 << iota //iota = 0
	j = 3 << iota //	iota = 1
	k             //iota = 2
	l             //iota = 3
	g = iota      //iota = 4
	p = 4 << iota //iota = 5
)

//
//func main() {
//	fmt.Println("i=", i)
//	fmt.Println("j=", j)
//	fmt.Println("k=", k)
//	fmt.Println("l=", l)
//	fmt.Println("4<<5", 4<<5)
//	var a int = 21
//	var b int = 10
//	var c int
//
//	c = a + b
//	fmt.Printf("第一行 - c 的值为 %d\n", c)
//	c = a - b
//	fmt.Printf("第二行 - c 的值为 %d\n", c)
//	c = a * b
//	fmt.Printf("第三行 - c 的值为 %d\n", c)
//	c = a / b
//	fmt.Printf("第四行 - c 的值为 %d\n", c)
//	c = a % b
//	fmt.Printf("第五行 - c 的值为 %d\n", c)
//	a++
//	fmt.Printf("第六行 - a 的值为 %d\n", a)
//	a = 21 // 为了方便测试，a 这里重新赋值为 21
//	a--
//	fmt.Printf("第七行 - a 的值为 %d\n", a)
//
//}
