package main

import "fmt"

func main() {
	var arr1 [5]int
	for i:=0; i<len(arr1);i++ {
		arr1[i] = i*2
	}

	for i:=0; i<len(arr1);i++{
		fmt.Printf("index:%d value:%d \n",i,arr1[i])
	}
	//go 语言遍历数组
	for i,_ := range arr1{
		fmt.Println(i)
	}
}

