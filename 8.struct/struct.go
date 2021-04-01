//go 语言结构体
package main

import "fmt"

func main(){

	//结构体定义
	type st struct {
		a int
		b int
		c,d string //同一个类型的可以同一行
	}
	//申明一个变量为struct，结构体实例化
	var getStr st

	//结构体成员初始化
	getStr.a = 1
	getStr.b = 2
	getStr.c = "hello"
	getStr.d = "world"
	fmt.Println(getStr)


	fmt.Println("-------创建指针类型的结构体--------")
	//指针类型的结构体
	pointerStruct := new (st)
	fmt.Println(pointerStruct)
	//结构体成员初始化
	pointerStruct.a = 126
	pointerStruct.c = "php is the best program"
	fmt.Println(pointerStruct)

	fmt.Println("-------取结构体的地址实例化--------")
	//直接取结构体的地址，这种方法使用和new一样，不过经常用
	addressPointer := &st{}
	addressPointer.a = 100
	addressPointer.d = "address"
	fmt.Println(addressPointer)

	fmt.Println("--------使用键值对初始化结构体-------")
	st1 := st{
		a:1,
		b:1,
		c:"hello",
		d:"world",
	}
	fmt.Println(st1)

	fmt.Println("--------使用多个值的列表初始化结构体-------")
	st2 := st{
		1,
		1,
		"hello",
		"world",
	}
	fmt.Println(st2)
	fmt.Println("-------初始化匿名结构体-------")
	//不初始化
	_st := struct{
		name string
		age int
		email int
		sql []string
	}{

	}
	fmt.Println(_st)
	_st1 := struct{
		name string
		age int
		email string
		sql []string
	}{
		name:"zhangsan",
		age:20,
		email:"12@qq.com",
	}
	fmt.Println(_st1)
	fmt.Println("-------模拟构造函数------")

	fmt.Println(*newName("kitty"))
	fmt.Println(*newColor("black"))

	fmt.Println("-------父子关系的构造函数------")
	fmt.Println("-------go方法和构造器------")
	//面向过程实现方法
	bag := new(Bag)
	Insert(bag, 1001)
	fmt.Println(bag)

	//使用方法类型（接收器）
	b := new(Bag)
	b.Insert(1001)
	fmt.Println(b)

	//使用指针类型的接收器

	p := new(Program)
	p.SetName("golang")
	fmt.Println(p.GetName())

	//使用非指针类型的接收器
	//这么看，非指针类型的使用在申明结构体的时候，不为指针类型就ok了
	var pNonPointer Program
	pNonPointer.SetName("python")
	fmt.Println(pNonPointer.getNameNonPointer())

	fmt.Println("-------为任意类型添加方法------")
	//必须先声明成这种类型的，直接使用 b:= 1会报错，找不到IsZero()方法
	var int1 MyInt
	fmt.Println(int1.isZero())
	int1 = 10
	fmt.Println(int1.AddNum(21))
	fmt.Println("-------方法和函数的统一调度------")

	//声明一个函数回调
	var delegate func(int)
	//创建结构体实例
	c := new(class)
	//将回调设置为c的Do方法
	delegate = c.Do
	//调用
	delegate(100)

	//将回调设为普通函数

	delegate = funcDo
	delegate(200)

	fmt.Println("-----------结构体内嵌-----------")
	var littleTiger Tigger
	littleTiger.kind.name = "litter tigger"
	littleTiger.kind.color = "yellow"
	littleTiger.eat = "meat"
	fmt.Println(littleTiger)

	//省略字段名
	var fox Fox
	//fox.name = "I am fox"
	//fox.color = "red"
	fox.hobby = "I like stolen"
	//也可以写全
	fox.Animal.name = "another name fox"
	fox.Animal.color = "another color red"
	fmt.Println(fox)

	//结构体模拟继承
	human := new (Human)
	bird := new(Bird)
	human.walk()
	bird.walk()
	bird.fly()
	fmt.Println("-----------初始化内嵌结构体------------")
	//初始化内嵌结构体
	desk := House{
		Desk:Desk{
			price: 2.01,
			color: "blue",
		},
	}
	fmt.Println(desk)

	//初始化匿名结构体
	desk1 := House{
		Desk:Desk{
			price: 2.01,
			color: "blue",
		},
		Room: struct {
			size float32
			num  int
		}{size: 8.88, num: 20},
	}
	fmt.Println(desk1)

}
//定义一个结构体
type Desk struct {
	price float32
	color string
}
type House struct {
	Desk
	//内嵌一个匿名结构体
	Room struct{
		size float32
		num int
	}
}




//飞行
type Flying struct {

}
//给flying添加方法
func (f *Flying) fly(){
	fmt.Println("can fly")
}

//走路
type Walking struct {

}
//给walking 添加方法
func (w * Walking) walk(){
	fmt.Println(" can walk")
}

type Human struct {
	Walking
}

type Bird struct {
	Flying
	Walking
}


type Animal struct {
	name string
	color string
}

type Tigger struct {
	kind Animal
	eat string
}

type Fox struct {
	//可以省略字段名
	Animal
	hobby string
}

type class struct {

}

//给结构体添加Do方法

func (c *class) Do(v int){
	fmt.Println("call method do:",v)
}
//普通函数

func funcDo(v int){
	fmt.Println("call function do:",v)
}

type MyInt int
//给整型添加判断是否为0的方法
func (m MyInt) isZero() bool{
	return m == 0
}

//必须是同一种类型的
func(m MyInt) AddNum(num MyInt)MyInt{
	return m+num
}

type Program struct {
	Name string
}
//使用指针类型接收器来设置名称
func (p *Program) SetName(pname string){
	p.Name = pname
}
//使用指针类型的接收器来获取名称
func (p *Program) GetName() string{
	return p.Name
}

//使用非指针类型的接收器来设置名称

func (p Program) setNameNonPointer(pname string){
	p.Name = pname
}

//使用非指针类型的接收器来获取名称

func(p Program) getNameNonPointer() string{
	return p.Name
}

type Bag struct {
	items []int
}
//使用接收器
func (b *Bag) Insert(itemid int)  {
	b.items = append(b.items, itemid)
}

func Insert(b *Bag, itemid int){
	b.items = append(b.items, itemid)
}


type Cat struct {
	Color string
	Name string
}
type BlackCat struct {
	Cat  // 嵌入Cat, 类似于派生
}

func newBlackCat(color string) *BlackCat{
	cat := &BlackCat{}
	cat.Color = color
	return cat
	//不能使用下面的写法，会报错
	//return &Cat{
	//	Color:color,
	//}

}
func newName(name string) *Cat{
	return &Cat{
		Name:name,
	}
}

func newColor(color string) *Cat{
	return &Cat{
		Color:color,
	}
}
