/**
firstday to learn goLang
 */

//func main() {
//	fmt.Println("hello, world")
//	fmt.Print("hello \n")
//	//常量定义
//	const a = 23;
//	fmt.Print(a)
//	const b string = "I am goLang\n"
//	fmt.Print(b);
//
//	//常量枚举
//	const (
//		c = iota
//		d
//		e
//	)
//	fmt.Print(c)
//	fmt.Print(d)
//	//变量
//
//	var name string = "myname is go";
//	fmt.Print(name)
//}
package main
import (
	"fmt"
	"runtime"
	"os"
)

func main() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}


