package main

import "fmt"

//带缓冲通道

func main(){
	fmt.Println("带缓冲的通道")
	ch := make(chan int,3)
	fmt.Println(len(ch))
	ch <- 1
	ch <- 2
	fmt.Println(len(ch))

	//创建一个无缓冲的通道
	var ch2 chan int
	ch2 = make(chan int)
	fmt.Println(ch2)
	go func() {
		ch2 <- 1
		ch2 <- 2

	}()
	//data := <- ch2
	for i := range ch2{
		fmt.Println(i)
	}

}
