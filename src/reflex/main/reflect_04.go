package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (this Cal) GetSub(name string) {
	fmt.Printf("%s 完成了对GetSub的调用",name)
}


func CalReflect(num interface{}) {
		//typ := reflect.TypeOf(num)

		val := reflect.ValueOf(num)

		kd := val.Kind()

		if kd!=reflect.Struct {
			fmt.Println(" 不是结构体")
			return
		}

		valNum := val.NumField() // 获取属性个数


		// 遍历属性
	for i:=0; i<valNum;i++  {
		fmt.Printf("val[%d]:%v\n",i,val.Field(i))
	}

	modNum := val.NumMethod()
	if modNum<0 {
		fmt.Println("没有方法调用")
		return
	}

	var per []reflect.Value

	per = append(per,reflect.ValueOf("杨康"))

	res:=val.Method(0).Call(per)
	fmt.Println(res)

}
func main() {
	cat := Cal{
		Num2:12,
		Num1:34,
	}

	CalReflect(cat)
}
