/**
按照单个文件编译，不是按照package编译
*/

package main

import "fmt"

func main(){
	//申明字符串切片
	var a1 []string
	fmt.Printf("%T\n", a1)

	//给切片添加数据
	a1 = append(a1, "hello")
	a1 = append(a1, "world")
	fmt.Println(a1)

	//以数组的方式生成切片
	var a2  = [...]int{1,2,3,4,5}
	fmt.Println(a2)
	var a3 []int = a2[0:2]
	fmt.Println(a3)

	//切片与数组的区别
	var a4 = []int{1,2,4}
	fmt.Println(a4)
	fmt.Printf("%T\n" , a4)

	var a5 = [...]int{1,2,4}
	fmt.Printf("%T\n", a5)
	var a6 [3]int
	fmt.Printf("%T\n", a6)

	//以make方式创建切片
	a7 := make([]int , 10)
	fmt.Printf("%T\n", a7)
	


}
