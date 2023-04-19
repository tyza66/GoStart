package main

import "fmt"

func main() {
	var a = 10
	var b = 3

	//算数运算符+ - * / % ++ --
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++
	fmt.Println(a)
	b--
	fmt.Println(b)

	//关系运算符
	var a1 = 11
	var a2 = 12
	fmt.Println(a1<=a2) //关系运算符返回的结果是布尔值
	if a1 == a2 {
		fmt.Println("a1 == a2")
	}else{
		fmt.Println("a1!= a2")
	}

	//逻辑运算符
	var t = true;
	var f = false;

	if t&&f{
		fmt.Println("t&&f为真")
	}else {
		fmt.Println("t&&f为假")
	}

	if t||f{
		fmt.Println("t||f为真")
	}

	if t!=f {
		fmt.Println("t!=f为真")
	}

	fmt.Println(!f)

	//位运算符(二进制)
	//二进制里面是逢二进一
	//60 0011 1100
	//13 0000 1101
	//& 都是一得到的才是一
	//| 都是零得到的才是零
	var n1 uint= 60
	var n2 uint= 13
	fmt.Printf("%d,%b\n",n1&n2,n1&n2)
	fmt.Printf("%d,%b\n",n1|n2,n1|n2)
}
