package main

import "fmt"

// 通过这样实现了继承的关系
type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human  //匿名字段Human
	school string
	loan   float32
}
type Employee struct {
	Human   //匿名字段Human
	company string
	money   float32
}

//通过下面的方式向"对象"进行方法绑定

// Human对象实现Sayhi方法
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human对象实现Sing方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

// Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee重载Human的Sayhi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}

// Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

// Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

// 如果某些对象是实现了接口中定义的全部方法 那么这个对象就实现了这个接口 也就是说任意类型都实现了interface{}接口
// 定义interface
type Men interface {
	SayHi()
	Sing(lyrics string)
}
type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}
type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

//

func main() {
	var stu Student = Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	stu.Sing("giao")
	stu.SayHi()
	var i Men
	i = &stu
	i.SayHi()
	i = &paul
	i.SayHi()
	x := make([]Men, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = &paul, &sam, &mike
	for _, value := range x {
		value.SayHi()
	}
	//空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface
	// 定义a为空接口
	var a interface{}
	//var i int = 5
	s := "Hello world"
	// a可以存储任意类型的数值
	//a = i
	a = s
	fmt.Println(a)
}
