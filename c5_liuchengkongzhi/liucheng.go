package main

import "fmt"

func main() {
	//流程控制一般就有三种 顺序 选择 循环
	//单个if
	var a int = 20
	if a < 30 {
		fmt.Println("a<30")
	} else {
		fmt.Println("a不<30")
	}
	//if else if
	var score = 80
	if score < 60 {
		fmt.Println("不及格")
	} else if score >= 60 && score < 80 {
		fmt.Println("及格")
	} else if score >= 80 {
		fmt.Println("优秀")
	}

	//switch go中的switch是支持任意类型的 并且不需要break
	var score1 = 60
	switch score1 {
	case 40:
		fmt.Println("将就")
	case 60:
		fmt.Println("还行")
	default:
		fmt.Println("giao")
	}
	//不给值的时候默认值是true
	switch {
	case true:
		fmt.Println("true")
	}

	//以为go里面不需要break那样 我们想实现其他语言里面那种穿透 就可以用fallthrough
	//只能穿透一次 但是在这一次穿透中 我们如果有条件 可以使用break终止穿透
	var b = false
	switch {
	case true:
		fmt.Println("giao1")
		fallthrough
	case false:
		if b == false {
			break
		}
		fmt.Println("giao2")
	default:
		fmt.Println("giao3")
	}

	//for循环
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	//求和1到10
	var sum = 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	//go里面for循环的三个条件哦度可以省略,但是要在循环里面控制循环条件
	//可以使用break跳出循环 使用continue跳出本次循环
	for {
		break
	}
	//for遍历字符串
	str := "giaogiao"
	fmt.Println("字符长度为:", len(str)) //查看字符串的长度
	fmt.Printf("第二个字节为:%c\n", str[1])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
	fmt.Println()
}
