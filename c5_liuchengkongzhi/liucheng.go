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
}
