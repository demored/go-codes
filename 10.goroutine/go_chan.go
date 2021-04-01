package main

import (
	"fmt"
)


//go 语言通道

//发送数据
func sendData(ch chan string){
	ch <- "hello"
	ch <- "world"
	ch <- "fuck"
	ch <- "you"
}
//接收数据
func getData(ch chan string){
	for i:= range ch{
		fmt.Println(i)
	}
}

func main(){
	fmt.Println("chan use")
	//var ch chan string
	//ch = make(chan string)
	//go sendData(ch)
	//go getData(ch)

	//实验通道阻塞
	ch := make(chan int)
	go sync(ch)
	for i:=100 ; i>1 ; i--{
		ch <- i
	}

}

//实验通道阻塞
func sync(ch chan int){
	data := <- ch
	fmt.Println(data)
}

