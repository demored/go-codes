package main

import "fmt"

//通道打印的例子

func print(c chan int)  {

	for {
		data := <- c

		//如果往通道中发送0则循环终止
		if data == 0{
			break
		}

		fmt.Println(data)
	}
	//告诉main循环已经结束
	c <- 0

}


func main(){

	ch := make(chan int)
	go print(ch)
	for i := 1; i< 10; i++{

		ch <- i
	}

	//通知并发的print循环结束，如果有通知则接收
	ch <- 0
	data := <- ch
	if data == 0 {
		fmt.Println("print结束了")
	}
}