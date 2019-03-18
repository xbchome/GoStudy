package main

import (
	"fmt"
	"strings"
)


func Suffix(hz string) func(name string) string {

	return  func(name string) string {
		if strings.HasSuffix(name,hz) {
			return name
		}
		return name+hz
	}
}
func main() {
	f :=Suffix(".jpg")

	fmt.Println("方法:",f("haha.jp"))
	fmt.Println("f2:",f("xbc.jpg"))
	fmt.Println("max：",max(3,4))
}


func max (num1 int,num2 int ) int {

	return  num1 - num2
}



