package main

import (
	"fmt"
	"unicode/utf8"
)

//字符串操作

func main(){
	var a = "hello world"
	a1 := "this is golang"
	fmt.Println(a)
	fmt.Println(a1)

	//为多行字符串赋值
	a2 := `
		tody is sunday
		yesterday is cloud
	`
	fmt.Println(a2)

	//字符串的操作

	//获取字符串长度
	a3 := "golang"
	fmt.Println(len(a3))  //6

	a4 := "中华人民共和国"
	fmt.Println(len(a4))  //21
	fmt.Println(utf8.RuneCountInString(a4)) //7

	for i:=1; i<len(a4);i++{
		fmt.Printf("%c,%d\n",a4[i],a4[i])
	}

	for _,s := range a4{
		fmt.Printf("%c, %d \n", s,s)
	}

}


