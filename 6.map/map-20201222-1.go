/**
按照单个文件编译，不是按照package编译
*/


package main
import "fmt"

func main(){
	fmt.Println("20201222日")

	//声明+初始化方式一
	map1 := make(map[string]string)
	map1["hello"] = "hello"
	map1["world"] = "world"

	fmt.Println(map1)
	fmt.Println(map1["hello"])

	map3 := make(map[int]int)
	map3[1] = 1
	map3[4] = 4
	fmt.Println(map3)
	fmt.Println(map3[4])

	//声明+初始化方式二
	map2:= map[string]string{
		"hello":"world",
		"stu":"golang",
	}

	fmt.Println(map2)
}
