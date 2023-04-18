package main

import "fmt"

func main() {
	var a = 10
	var b = 3

	//+ - * / % ++ --
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++
	fmt.Println(a)
	b--
	fmt.Println(b)

	var a1 = 11
	var a2 = 12
	fmt.Println(a1<=a2) //关系运算符返回的结果是布尔值
	if a1 == a2 {
		fmt.Println("a1 == a2")
	}else{
		fmt.Println("a1!= a2")
	}
}
