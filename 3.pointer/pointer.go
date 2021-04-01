//指针运算

package main

import "fmt"

func init(){
	fmt.Printf("我是第一个执行的\n")
	if a:=10; a< 20 {
		fmt.Printf("a是小于20的\n")
	}else{
		fmt.Printf("a 是大于20的\n")
	}
}
func main(){
	s:= "hello world"
	var p *string = &s

	p1 := &s

	fmt.Println(p1)
	i:= 10

	var pi *int = &i

	fmt.Printf("s的地址是:%p\n", &s)
	fmt.Printf("s的值是:%s\n", s)
	fmt.Printf("p的值是:%p\n", p)
	fmt.Printf("p引用的值是:%s\n", *p)

	fmt.Printf("i的地址是:%p\n", &i)
	fmt.Printf("i的值是: %d\n", i)
	fmt.Printf("pi的值是:%p\n", pi)
	fmt.Printf("pi引用的值是:%d\n", *pi)

}