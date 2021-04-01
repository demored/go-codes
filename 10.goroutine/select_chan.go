package main

import (
	"fmt"
	"time"
)

//select  通道的多路复用


//往通道中写数据pump1
func pump1(ch chan int){
	for i:=1;; i++{
		ch <- i*2
	}
}
//往通道中写数据pump2
func pump2(ch chan int){
	for i:=1; ; i++{
		ch <- i+2
	}
}

//接收通道中的数据
func suck(ch1, ch2 chan int)  {
	for{
		select{
			case v:= <- ch1:
				fmt.Println("received on channel 1:",v)
			case v2:= <-ch2:
				fmt.Println("received on channel 2:", v2)
		}
	}
}

func main(){
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	//等所有的gpruntine创建完
	time.Sleep(time.Second)

}
