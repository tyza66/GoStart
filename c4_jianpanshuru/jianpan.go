package main

import "fmt"

func main() {
	//交互使用的是fmt包
	var a int
	var b float64

	fmt.Println("请输入连个数值 一个是整数 一个是64位单精度浮点数")
	//到这的时候程序会阻塞等待键盘输入
	fmt.Scanln(&a,&b)
	fmt.Println(a,b)
}
