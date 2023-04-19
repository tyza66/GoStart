package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(getSum(3, 4))
	//定义数组
	arr :=[4]int{1,2,3,4}
	fmt.Println(arr)
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

func p1(a string){
	fmt.Println(a)
}
//有两个参数,两个返回值的函数
func swap(a,b string) (string,string){
	return b,a
}

//可变参数的函数
func getSum(nums ... int) int{
	sum:=0;
	for i:=0;i<len(nums);i++{
		sum += nums[i]
	}
	return sum
}

func update(a [4]int){

}


