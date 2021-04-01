package main
import (
	"fmt"
	"sync"
	_ "time"

)

var wg sync.WaitGroup

func running(){
	fmt.Print("hello");
}

func hello(i int){
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}


func main(){
	fmt.Print("main process");
	//启动单个协程
	//go running();
	//time.Sleep(time.Second)

	//启动多个协程
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束

}