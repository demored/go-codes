package main

import (
	"fmt"
)

type User struct {
	id int
	name string
	email string
}

// strcut test

func main(){

	/************************结构体实例化***********************/
	//基本实例化
	//var user User

	//new 实例化
	//user := new(User)

	//取地址实例化
	user := &User{}

	user.id = 1
	user.name = "zhangsan"
	user.email = "demo@163.com"

	fmt.Println(user)

	/************************结构体成员初始化***********************/

	//使用键值对填充，返回的是结构体
	user1 := User{
		id:1,
		name:"zhangsan",
	}
	fmt.Println(user1)

	//使用键值对填充，返回的是指针
	user2 := &User{
		id:1,
		name:"zhangsan",
	}
	fmt.Println(user2)

	/************************初始化匿名结构体***********************/

	//不初始化成员
	user3 := &struct{
		id int
		name string
	}{}
	fmt.Println(user3)

	//初始化成员
	user4 := &struct{
		id int
		name string
	}{
		id:2,
		name:"Miaoxinren",
	}
	fmt.Println(user4)
}



