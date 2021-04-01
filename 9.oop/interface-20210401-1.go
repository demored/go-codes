package main

import "fmt"

type Maner struct {
	name string
	age int32
}

type Sunner interface {
	say(interface{})
}

func (m *Maner) say(parm1 interface{}){
	fmt.Println(m.name, m.age)
	fmt.Println("传入的参数", parm1)
}

func main(){
	var s Sunner

	m := &Maner{
		name: "张三",
		age:20,
	}
	s = m

	s.say(20)
	m.say(100)

}







