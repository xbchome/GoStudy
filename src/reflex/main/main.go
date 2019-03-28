package main

import (
	"fmt"
	"reflect"
)
/*
反射
	typeOf
	valueOf
 */
func main() {
	fmt.Sprintf("v")
	ints := 2
	getM(ints)
	//fmt.Print(reflect.TypeOf(getM))
}


func getM(b interface{}) {
	rType := reflect.TypeOf(b)
	rvalue:= reflect.ValueOf(b)
	fmt.Println(rType)
	fmt.Println(rType.Name())
	fmt.Println(rvalue)
}
