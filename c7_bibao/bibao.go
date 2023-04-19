package main

import "fmt"

func main() {
	r1 := increment()
	fmt.Println(r1)
	v1 :=r1()
	fmt.Println(v1)
	v2 :=r1()
	fmt.Println(v2)

	//正常情况下调用完函数应该已经被销毁了,但是在闭包中那个存着的变量不会被销毁
	/*
	一个外层函数中,有内层函数,在这个内层函数中会操作外层函数的局部变量
	并且外层函数的返回值就是这个内层函数
	这个内层函数和外城函数的局部变了就称为闭包函数
	局部变量的生命周期会发生改变 正常的函数变量会随诊函数的调用而创建,随着调用结束而销毁
	但是闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁
	*/
	r2 := increment()
	fmt.Println(r2)
	v3 :=r2()
	fmt.Println(v3)
}

func increment() func() int{
	i := 0
	fun := func() int { //内层函数没有执行
		i++
		return i
	}
	return fun
}
