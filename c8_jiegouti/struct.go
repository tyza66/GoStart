package main

import "fmt"

func main() {
	//定义变量类型
	type person struct {
		name string
		age  int
	}
	var P person
	P.name = "John"
	P.age = 18
	fmt.Println(P)

	//快速方式赋值
	P = person{"Tom", 25}
	fmt.Println(P)

	//快速方式定义自定义类型的变量
	P2 := person{age:24, name:"Tom"}
	fmt.Println(P2)

	//匿名字段
	type Human struct {
		name string
		age int
		weight int
	}
	type Student struct {
		Human  // 匿名字段，那么默认Student就包含了Human的所有字段
		speciality string
	}
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	fmt.Println(mark)
	fmt.Println(mark.age) //访问里面的匿名struct的时候可以像访问自己的结构体成员一样
	fmt.Println(mark.Human.name)
}
