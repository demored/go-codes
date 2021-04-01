//go语言数组

package main

import "fmt"

func main(){
	var abc [3] string
	abc[0] = "I"
	abc[1] = "like"
	abc[2] = "golang"
	fmt.Println(abc)


	//申明+初始化
	var names = [3]string{"duola","checkorder","findThis"}
	fmt.Println(names)

	//申明整型
	var num = [4] int {1,2,3,4}
	fmt.Println(num)

	//遍历数组
	for k,v := range num {
		fmt.Println(k,v)
	}
	fmt.Println("-----------------")
	//动态生成一个数组
	var createArr[30] int

	for a:=0; a<30 ; a++{
		createArr[a] = a
	}
	fmt.Println(createArr)
	fmt.Println("--------以下是切片的内容---------")

	//------------------切片--------------------
	var stu = [...] string {"duola","checkorder","findThis"}
	fmt.Println(stu)
	fmt.Println(stu,stu[1:2],stu[:],stu[0:1])
	//复制一个切片
	var stu1 = []int{1,2,3}
	fmt.Println(stu1[:])
	//生成一个空的切片
	fmt.Println(stu[0:0])
	//申明一个新的切片
	var strList [] string
	//申明及初始化一个切片
	var intList = [] int{}
	fmt.Println(strList, intList)
	fmt.Println(len(strList), len(intList))
	fmt.Println(strList == nil, intList == nil) //true false

	//使用make动态构造切片
	nameList := make([]string, 2)
	ageList := make([]int, 2 , 10)
	fmt.Println(len(nameList), len(ageList))


	//使用append为切片添加元素
	var emailList []string
	emailList = append(emailList, "12345@qq.com")
	fmt.Println(emailList)
	for i:= 0; i < 20; i++ {
		emailList = append(emailList, string(i))
	}
	fmt.Println(emailList)
	//将一个切片中的元素添加到另一个切片
	foxList := []string{"fox1","fox2"}
	emailList = append(emailList, foxList...)
	fmt.Println(emailList)
	//输出切片中的元素
	fmt.Println(emailList[0])

	//切片复制
	fmt.Println("---------------")
	const elementCount = 1000
	srcData :=  make([]int, elementCount)
	for i := 0 ; i<elementCount; i++{
		srcData[i] = i
	}
	fmt.Println(srcData)
	refData := srcData
	srcData[0] = 999
	fmt.Println(srcData[0], refData[0]) //切片是引用赋值
	copyData := make([]int, elementCount)
	copy(copyData, srcData) //复制切片
	fmt.Println(len(copyData), len(srcData))
	//修改原切片，观察复制切片
	srcData[0] = 126
	fmt.Println(copyData[0], srcData[0])
	//切片删除
	srcData = append(srcData[:2], srcData[998:]...)
	fmt.Println(srcData)
	srcData = srcData[0:0]
	fmt.Println(srcData)

}

