package main

import "fmt"

//空接口的使用
func main(){

	var any interface{}
	any = 1
	fmt.Println(any)

	any = "hello"
	fmt.Println(any)

	any = false
	fmt.Println(any)

	//这样直接会报错
	//var name string = any
	var name bool = any.(bool)
	fmt.Println(name)


}




