//列表list的操作
package main

import (
	"fmt"
	"container/list"
)
func main()  {
	//声明一个列表
	l := list.New()
	fmt.Println(l)
	//往列表中插入数据
	l.PushBack("hello")
	l.PushFront("golang")
	fmt.Println(l.Front().Value) //打印第一个
	fmt.Println(l.Front().Next().Value) //打印第二个
	//列表循环
	for i:= l.Front(); i != nil ; i= i.Next(){
		fmt.Println(i.Value)
	}
	//列表删除，删除第一个值
	l.Remove(l.Front())
	fmt.Println(l.Front().Value)

}
