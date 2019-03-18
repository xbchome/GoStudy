package main

import "fmt"

/*
	map 是引用类型
 */
func main() {

	stu := make(map[string]map[string]string)

	stu["01"] = make(map[string]string,2)
	stu["01"]["name"] = "xbc"
	stu["01"]["age"] = "xbc"

	valeu,err := stu["01"]
	fmt.Println("=======================")
	fmt.Println(valeu)
	fmt.Println(err)
	fmt.Println("=======================")


	fmt.Println(stu)
	fmt.Println("=======================")
	var mapTest  = map[string]string{
		"xbc":"想不出",
		"xbc1":"想不出",
		"xbc2":"想不出",
	}

	fmt.Println(mapTest)

}
