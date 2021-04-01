// go语言函数使用
package main
import (
	"fmt"
)

func testFunc() int{
	println("hello word");
	//必须要有返回值
	return 0;
}

func fire() {
	fmt.Println("fire")
}
func add(a, b int)int{
	return a+b
}
func main() {
	//声明一个函数
	var f func() //要先申明
	f = fire
	f() //函数调用
	fire() //直接调用
	var c func(m,n int)int
	c = add

	fmt.Println(c(1,2))

	fmt.Println("---------匿名函数---------")

	f1 := func(a, b int){
		fmt.Println("a+b=",a+b)
	}
	f1(1,2) //调用函数

	//匿名函数直接调用
	func (m, n int){
		fmt.Println("m+n= ",m+n)
	}(2,3)

	//匿名函数作为参数
	//遍历一个数据并打印
	visit([]int{1,2,3,4},func(v int){
		fmt.Println(v)
	})
	fmt.Println("-----闭包-----")
	//闭包closure
	var acc func() int //调用函数之前一定要先声明
	acc = Accumulate(1)
	acc()
	acc()

	fmt.Println("-----可变参数-----")
	changeParam("python", "go", "php", "mysql", "node", "shell")

	fmt.Println("-----defer延迟执行语句-----")
	fmt.Println("hello world")
	defer fmt.Println("hello python") //最后执行
	fmt.Println("hello go")

	fmt.Println("-----panic宕机-----")
	fmt.Println("我要宕机了")
	panic("宕机")
}

//定义一个可变参数的函数
func changeParam(slist...string){
	for _,v := range slist{
		fmt.Println(v)
	}
}

//创建一个闭包的累加器
func Accumulate(value int) func()int{
	return func()int{
		value++
		fmt.Println(value)
		return value
	}
}

func visit(list[]int , ff func(v int)){
	for _,v := range list{
		ff(v)
	}
}