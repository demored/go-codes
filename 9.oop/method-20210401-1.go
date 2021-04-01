/**
按照单个文件编译，不是按照package编译
*/

package main

import "fmt"

type St1 struct {
	name string
	age int
	doing  []string
}

func (s St1) getS (){
	fmt.Println("go使用方法")
	fmt.Println(s.name)
}

func main(){
	s := &St1{
		name:"张三",
		age:20,
	}

	s.getS()
}
