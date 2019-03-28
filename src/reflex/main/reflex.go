package main

import (
	"fmt"
	"reflect"
)

func test(b interface{}) {
	// 通过反射获取类型
	varType := reflect.TypeOf(b)
	fmt.Println("varType:",varType)
	// 通过反射获取值
	varValue := reflect.ValueOf(b)
	fmt.Println("varType:",varValue)

	// 将值转为 interface{}

	varV := varValue.Interface()
	fmt.Println("varV.int",varValue.Int())
	fmt.Println(varV.(int)) // 通过类型断言 转为int

	// 获取kind

	fmt.Println("kind-",varValue.Kind())
}

func main() {
	b:= 4
	test(b)
}

