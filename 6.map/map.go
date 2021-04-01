//创建map，map使用

package main
import "fmt"
func main(){
	//----------map声明+初始化--------------
	//声明+初始化方式一
	mapList := make(map[string]string)
	fmt.Println(mapList == nil)
	//给map填充数据
	mapList["hello"] = "world"
	fmt.Println(mapList, mapList["hello"])
	//申明+初始化方式二
	var map1 = map[string]string{
		"hello":"world",
		"name":"zhangsan",
	}
	fmt.Println(map1["hello"])
	fmt.Println("---------map循环----------")
	//map简单循环
	mapList2 := make(map[string]string)
	mapList2["job"] = "develop"
	mapList2["age"] = "29"
	for k,v := range mapList2{
		fmt.Println(k,v)
	}
	//只循环k
	for k := range mapList2{
		fmt.Println(k)
	}
	//只循环v
	for _,v := range mapList2{
		fmt.Println(v)
	}
	//将map中的数组循环到切片中
	var list1 []string
	for k := range mapList2{
		list1 = append(list1, k)
	}
	fmt.Println(list1)
	//删除key值
	delete(mapList2, "job")
	for k,v := range mapList2{
		fmt.Println(k,v)
	}

}
