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

	//在函数运行到最后的时候这些defer语句会逆序执行
	defer f()
	defer f2()
	fmt.Println(digui(5))
	testDefer()
	//函数本身也是一种类型dunc() 我们可以直接打印出来 加括号算调用
	//函数的类型就是定义的函数类型本身那些
	fmt.Printf("%T\n",add)
	var f5 func(int, int) int
	f5 = add
	fmt.Println(f5(1,2))
	//匿名函数
	func(){
		fmt.Println("giaogiao")
	}()
	func(a,b int) int{
		fmt.Println(a,b)
		return a+b
	}(1,2)

	//回调函数 我们可以把一个函数作为另一个函数的参数
	//将dun1函数作为fun2函数的参数 那么fun2函数就叫高阶函数 fun1函数就叫回调函数
	i := add2(1, 2, add)
	fmt.Println(i)
}

// 函数名()里面是参数类型 后面的是返回值类型,函数体
// 有两个参数 有一个返回值的函数
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

// 有两个参数,两个返回值的函数
func swap(a, b string) (string, string) {
	return b, a
}

// 可变参数的函数
func getSum(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

// 值传递传递的是变量的副本 修改了对原来的数据没有任何影响 基础数据类型,array,struct都是值传递
func update(a [4]int) {
	fmt.Println("变换前", a)
	a[0] = 100
	fmt.Println("变化后", a)
}

// 引用类型传递的是地址空间 所以函数里面的变了 外边的就也变了
func update2(s2 []int) {
	fmt.Println("变换前", s2)
	s2[0] = 100
	fmt.Println("变换后", s2)
}

func digui(n int) int {
	if n == 1 {
		return 1
	}
	return digui(n-1) + n
}

// 在go中可以使用defer来延迟一个函数的执行
func f() {
	fmt.Println("giao")
}

func f2() {
	fmt.Println("giao2")
}

func testDefer() {
	//用于关闭文件啥的很方便
	a := 10
	fmt.Println("a=", a)
	defer testDefer2(a)
	a++
	fmt.Println("a=", a)
}

func testDefer2(s int) {
	fmt.Println("函数接收到的a-", s)
}

func add2(a,b int, fun func(int,int) int) int{
	return fun(a,b)
}
