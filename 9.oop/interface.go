package _interface

import "fmt"

//定义一个接口
type chFind interface {
	FindCh() []rune
	printSelf()
}
//定义一个字符串型
type MyString string
type MyInt int
//实现接口
func (ms MyString) FindCh()[]rune{
	var chs []rune
	for _,v := range ms{
		if v == 'a' || v == 'b'{
			chs = append(chs, v)
		}
	}
	return chs
}

type Test interface {
	Tester()
}
//type myFloat float32
//
//func (t myFloat)Tester(){
//	fmt.Println("我是类型实现的接口方法")
//}

func (bb MyInt)Tester(){
	fmt.Println("我是类型实现的接口方法")
}

func findType(i interface{}) {
	//类型判断
	switch i.(type) {
	case string:
		fmt.Printf("String: %s\n", i.(string))
	case int:
		fmt.Printf("Int: %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}
type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}
func findType2(i interface{}) {
	switch v := i.(type) {
	case Describer:
		fmt.Println("-------哈哈",v)
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func main(){
	//str :=  MyString("Sabm")
	//定义一个接口类型的变量
	//var mm chFind
	//将str 复制给 mm
	//mm = str
	//直接使用接口方法
	//fmt.Println(mm.FindCh())

	//直接使用字符串的接口方法，2者结果都一样，但是接口更优雅一点
	//如果接口中还存在其他方法，使用接口调用 则必须要全部实现，否则会报错
	//但是字符串正常，不会报错
	//fmt.Println(str.FindCh())

	//var aa myFloat
	//aa = 100.01
	//aa.Tester() //通过一个类型去实现

	var t Test
	var aaa MyInt
	aaa = 100

	t = aaa
	t.Tester()
	
	//类型判断
	findType("Naveen")
	findType(77)
	findType(89.98)
	////////////类型与接口进行比较/////////////
	fmt.Println("------------型与接口进行比较--------")
	findType2("Naveen")
	p := Person{
		name: "Naveen R",
		age:  25,
	}
	findType2(p)

}