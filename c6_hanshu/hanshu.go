package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(getSum(3, 4))
	//定义数组
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	update(arr)
	fmt.Println("调用后", arr)
	//切片就是一个可以扩容的数组 就是没写大小
	s1 := []int{3, 2, 3, 4}
	update2(s1)
	fmt.Println("调用后", s1)
}

// 函数名()里面是参数类型 后面的是返回值类型,函数体
//有两个参数 有一个返回值的函数
func add(a, b int) int {
	c := a + b
	return c
}

// 无参无返回值的函数
func p() {
	fmt.Println("aaa")
}

func p1(a string) {
	fmt.Println(a)
}

//有两个参数,两个返回值的函数
func swap(a, b string) (string, string) {
	return b, a
}

//可变参数的函数
func getSum(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

//值传递传递的是变量的副本 修改了对原来的数据没有任何影响 基础数据类型,array,struct都是值传递
func update(a [4]int) {
	fmt.Println("变换前", a)
	a[0] = 100
	fmt.Println("变化后", a)
}

//引用类型传递的是地址空间 所以函数里面的变了 外边的就也变了
func update2(s2 []int) {
	fmt.Println("变换前", s2)
    s2[0] = 100
    fmt.Println("变换后", s2)
}
