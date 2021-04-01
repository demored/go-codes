package main

import "fmt"

//关闭通道

func main(){

	ch := make(chan int)
	go putData(ch)
	ch <- 1
	ch <- 2

	//关闭通道
	close(ch)
}

//输出通道的数据
func putData(ch chan int){
	for data:= range ch{
		fmt.Println(data)
	}
}

