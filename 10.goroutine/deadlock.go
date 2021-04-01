package main

import (
	"fmt"
	"time"
)

//关于通道死锁的研究

//func getData(ch chan int){
//	data := <- ch
//	fmt.Println(data)
//}

func main(){

	//定义无缓冲通道
	ch := make(chan int)

	//go getData(ch) ，

	go func() {
		data := <- ch
		fmt.Println(data)
	}()

	//造成死锁，没有接收
	ch <- 1

	time.Sleep(time.Second)
	//close(ch)

	//看似有协程接收数据，但是上面的已经阻塞了，下面是执行不下去了
	//go func() {
	//	data := <- ch
	//	fmt.Println(data)
	//}()
	

}
