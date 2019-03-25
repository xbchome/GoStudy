package main

import (
	"encoding/json"
	"fmt"
)

type Student struct{
	Name string
	Age int
	Gender string
}

func jsonStudent() {
	 stu := Student{
	 	Name:"x",
	 	Age:12,
	 	Gender:"NV",
	 }

	 str,err := json.Marshal(stu)
	 if err!=nil{
	 	fmt.Println("转换出错")
	 }else {
	 	fmt.Println(string(str))
	 }

}

func jsonSlice () {
	var slices []Student
	slices = append(slices,Student{
		"LI",
		12,
		"NV",
	} )
	slices = append(slices,Student{
		"LI",
		12,
		"NV",
	} )

	str,err := json.Marshal(slices)

	if err != nil {
		fmt.Println("json转换失败")
	}else{
		fmt.Println("slices json",string(str))
	}
}

// 将map 转为json
func jsonMap() []byte {
	var str map[string] interface{}

	str = make(map[string]interface{})
	str["name"] = "lb"
	str["age"] = 12
	str["gender"] ="nv"

	jsonStr,err :=json.Marshal(str)
	if err != nil {
		fmt.Printf("jsonStr%v",string(jsonStr))
	} else{
		fmt.Printf("jsonStr%v\n",string(jsonStr))
	}
	return jsonStr
}

func GetList() {
	str := jsonMap()
	var maps map[string]interface{}

	json.Unmarshal(str,&maps)
	fmt.Println(maps)
}

func main() {
	jsonMap()
	fmt.Println("=================")
	jsonSlice()
	fmt.Println("=================")
	jsonStudent()
	fmt.Println("=======map=======")
	GetList()
}
