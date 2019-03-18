package main

import "fmt"

/*
  map slice 练习
 */
func main() {
	var mapSlices []map[string]string
	 mapSlices  = make([]map[string]string,3)

	 mapSlices[0] = make(map[string]string)

	 mapSlices[0]["name"] = "xbc"
	 mapSlices[0]["age"] = "123"
	 newMpa := map[string]string{
	 	"name":"xnbc",
	 	"age":"123",
	}
	 mapSlices = append(mapSlices,newMpa)

	 fmt.Println(mapSlices)
	 fmt.Println("=======================")
	 mapSlices[2] = map[string]string{
	 	"name":"nmw",
	 	"age":"123",
	 }

	 fmt.Println(mapSlices)
	 fmt.Println("=================练习二=============")
	 // 创建map slice

	 var newMapSlice []map[string]string

	 newMapSlice = make([]map[string]string,1)

	 newMapSlice[0] = map[string]string{
		"name":"nmw",
		"age":"123",
	}

	 map2 := map[string]string{
		 "name":"nmw",
		 "age":"123",
	 }

	 fmt.Println(map2)

	 fmt.Println(newMapSlice)

}
