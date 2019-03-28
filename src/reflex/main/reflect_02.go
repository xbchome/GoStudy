package main

import (
	"fmt"
	"reflect"
)

func reflect02(n interface{}) {
	rVal := reflect.ValueOf(n)
	fmt.Printf("rVal:%v rType%v \n",rVal,rVal.Kind())
	// 修改传入的值
	rVal.Elem().SetInt(3)
}

func main() {
	num := 2
	reflect02(&num)
	fmt.Println(num)
}