/**
	按照单个文件编译，不是按照package编译
 */
package main
import "fmt"
func test() int{
	return 2+3;
}

func main(){
	//使用匿名函数
	var f func() int
	f = test
	//f()
	a := test()
	fmt.Println(a)
	fmt.Println(f())

}

