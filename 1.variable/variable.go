//变量、短变量
//变量声明
//变量初始化
//多个变量声明

package main
import "fmt"

var (
	a int = 10
	b string
	c [] float32
	d func() bool
	e struct{
		x int
	  }

)
var email string = "zhangsan@163.com"

var name string
var a1 int
//a1 = 10

var a3 bool//bool类型

func main(){
	//必须在函数内部
	a2 := 111
	fmt.Println(a2)

	fmt.Println(a1)
	name := "hello"

	fmt.Print(name,"\n")
	var abc bool = true

	fmt.Print(abc)
	fmt.Println(a3)
	fmt.Println(a)
	
	fmt.Println()
}