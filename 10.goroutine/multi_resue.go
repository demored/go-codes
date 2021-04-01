package main

import "fmt"

//通道的多路复用

func multi_resue(ch1 chan int, ch2 chan int){
	select{
		case d:= <- ch1:
			fmt.Println(d)
		case d2:= <- ch2:
			fmt.Println(d2)
		default:
			fmt.Println("error")


	}
}
func main(){
	var ch1 chan int
	ch1 = make(chan int)
	ch2 := make(chan int)

	go multi_resue(ch1,ch2)
	 ch1 <- 1
	 ch2 <-2


}




