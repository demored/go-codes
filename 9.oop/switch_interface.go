package main

import (
	"fmt"
)

func checkType(a interface{}){
	switch a.(type){
	case int:
		fmt.Println("a is int")
	case string:
		fmt.Println("a is string")
	case bool:
		fmt.Println("a is bool")
	}
}


func main(){

	var age int
	age = 1100
	//判断类型
	checkType(age)
}
