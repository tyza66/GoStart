package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	/*f, err := os.Open("filename.ext")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)*/

	f1, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f1)

	/*
		Offset字段在调用Error的时候不会被打印，但可以通过类型断言获取错误类型，然后可以打印相应的错误信息，请看下面的例子:
			if err := dec.Decode(&val); err != nil {
			    if serr, ok := err.(*json.SyntaxError); ok {
			        line, col := findLine(f, serr.Offset)
			        return fmt.Errorf("%s:%d:%d: %v", f.Name(), line, col, err)
			    }
			    return err
			}
	*/

	/*
		需要注意的是，函数返回自定义错误时，返回值推荐设置为error类型，而非自定义错误类型，特别需要注意的是不应预声明自定义错误类型的变量。例如：

		func Decode() *SyntaxError { // 错误，将可能导致上层调用者err!=nil的判断永远为true。
		    var err *SyntaxError     // 预声明错误变量
		    if 出错条件 {
		        err = &SyntaxError{}
		    }
		    return err               // 错误，err永远等于非nil，导致上层调用者err!=nil的判断始终为true
		}
	*/
	//我们判断错误的方法就是检查错误信息是否为nil
}

// 有一个内置的接口类型error 里面只有一个返回值为string的方法Error()
// 自定义错误
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
func New(text string) error {
	return &errorString{text}
}
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}

type SyntaxError struct {
	msg    string // 错误描述
	Offset int64  // 错误发生的位置
}

func (e *SyntaxError) Error() string { return e.msg }
