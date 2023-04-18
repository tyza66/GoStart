package main

import "fmt"

func main() {
	//变量名name 变量类型string
	var name string = "giao"
	//定义不同类型的变量
	var age int = 18
	//同时定义多个变量
	var (
		//字符串默认是空串
		name1 string
		//int默认是0
		age1  int
		addr1 string
	)
	//给变量赋值
	name = "tyza66"
	//语法糖 直接声明并且赋值 并且创建变量的同时会自动识别我们赋予变量的值的类型
	name2 := "hello"
	age2 := 19
	fmt.Println(name, age)
	fmt.Println(name1, age1, addr1)
	fmt.Println(name2, age2)
	//使%T查看变量类型
	fmt.Printf("%T,%T", name, age)

	var num int = 11
	//%d是数值 %p是内存地址
	fmt.Printf("num:%d,内存地址:%p", num, &num)

	//改变变量的值不会改变地址
	num = 200
	fmt.Printf("num:%d,内存地址:%p", num, &num)

	fmt.Println()

	//go语言的变量交换不需要使用临时变量
	var a = 100
	var b = 200
	b, a = a, b
	fmt.Println(a, b)

	println()

	//go提供了一个匿名变量 任何赋值给它的值都将被抛弃 可以极大增强代码的灵活性
	a, _ = test() //像这种情况 这个函数返回的值 其中一个对我们没有用 我们就给废弃掉
	_, b = test()
	fmt.Println(a, b)

	//变量的作用域十分重要 因为定义变量但是没使用就会报错
	//定义在main函数外面的变量就是全局变量 在整个go文件中任何的函数都可以调用
	//定义在函数体内的变量就是局部变量 可以和全局变量重名 在函数体中优先调用的是局部变量

	//常量 常量使用const定义 但是值不可以被改变
	const URL string = "www.giao.com"
	const a1, b1, c1 = 1, 23, 3
	fmt.Println(a1, b1, c1)

	//有一个特殊的常量iota 它会自动进行累加 如果下面的没有给值 也是可以默认的走iota 如果不定义就会继续上一个
	const (
		a2 = iota
		b2 = iota
		c2 = iota
		d
		e = "tyza66"
		f
	)
	fmt.Println(a2, b2, c2, d, e, f)

	//bool默认值为false
	var isFlag bool
	fmt.Println(isFlag)
	//%t是值
	fmt.Printf("%T,%t", isFlag, isFlag)
	fmt.Println()

	//数值类型 无符号的前面加个uint8就是byte 他俩就是别名的关系
	var age4 uint = 18
	fmt.Printf("%T,%d", age4, age4)
	fmt.Println()
	var money float64 = 3.14
	fmt.Printf("%T,%f", money, money)
	fmt.Println()

	//定义float会造成精度损失 计算的时候尽量不要用float进行计算 尽量用64的
	var num1 float32 = -123.0000901
	var num2 float64 = -123.0000901
	fmt.Println("num1 =", num1, "num2 =", num2)

	
}

func test() (int, int) {
	return 100, 200
}

//建议遵循驼峰命名
