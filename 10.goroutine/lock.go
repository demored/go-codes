package main

//互斥锁与互斥读写锁
import (
	"fmt"
	"sync"
)
var (
	count int
	countGurad sync.Mutex

)

func getCount()int{

	countGurad.Lock()

	defer countGurad.Unlock()
	return count
}

func setCount(c int){
	countGurad.Lock()
	count = c
	countGurad.Unlock()
}


func main()  {
	setCount(1)

	fmt.Println(getCount())
}



