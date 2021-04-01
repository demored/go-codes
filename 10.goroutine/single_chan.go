package main

import "fmt"

func print(c <- chan int){

	for{
		data:= <- c
		if data == 0 {
			break
		}
		fmt.Println(data)
	}
}


//单向通道,继续使用打印的例子
func main()  {
		ch := make(chan int)

		//声明只能接收的通道
		var ch1 <- chan int = ch
		//声明只能发送的通道
		var ch2 chan <- int = ch

		//两个通道可以分开读写，有意思吧
		go print(ch1)

		for i:=1 ; i <20 ; i++{
			ch2 <- i
		}
		//总结，都可以正常输出

}