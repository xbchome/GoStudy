package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name string `json:"name"`
	Age int `json:"monster_age"`
	Score float32 `json:"成绩"`
	Gender string
}

func (this Monster) GetSum(n1,n2 int) int {
	return n1+n2
}

func (this Monster) Set(name string,age int, score float32,gender string) {
	this.Name = name
	this.Gender = gender
	this.Age = age
	this.Score = score
}

func TestStruct(obj interface{}) {
	//	获取 reflect.Type 类型
	typ:=reflect.TypeOf(obj)
	// 获取 reflect.value 类型
	val := reflect.ValueOf(obj)

	// 获取 类别
	kd := val.Kind()

	if kd!=reflect.Struct {
		fmt.Println("kd:",kd)
		fmt.Println("类型不正确")
		return
	}

	// 获取结构体字段个数
	num:=val.NumField()
	fmt.Println("结构体字段数：",num)

	for i:= 0 ;i<num ; i++ {
		fmt.Printf("Field %d: 值为=%v\n ",i,val.Field(i))
		// 获取 结构体的tag值
		tagVal:= typ.Field(i).Tag.Get("json")
		if tagVal!= "" {
			fmt.Println(tagVal)
		}else {
			fmt.Println("tagVal null")
		}
	}


	// 获取结构体方法
	numMet:=val.NumMethod()

	fmt.Printf("有 %d 个方法 \n",numMet)
	// 方法按 ascll 码排序
	val.Method(1).Call(nil) // 调用第二个方法

	var params []reflect.Value  // 声明reflect.Value 切片
	params = append(params,reflect.ValueOf(10))
	params = append(params,reflect.ValueOf(8))

	res := val.Method(0).Call(params)

	fmt.Println("res:",res[0].Int())



}

func (this Monster) Print() {
	fmt.Println("==================")
	fmt.Println(this)
	fmt.Println("==================")
}

func main() {
	menster := Monster{
		Name:"杨康",
		Age:12,
		Gender:"女",
		Score:4.3,
	}

	TestStruct(menster)
}
