package main

import (
	"fmt"
)

//定义一个接口
type DataWriter interface {
	WriteData(data interface{}) error
}

type file struct {

}

//实现结构体
func (d *file)WriteData(data interface{}) error{
	fmt.Println("WriteData:" , data)
	return nil
}

func main(){

	f := new(file)

	var dataWriter DataWriter

	dataWriter = f

	dataWriter.WriteData("interface")

}