package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var fh string
	fmt.Println("请输入两个数")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)

	fmt.Println("请输入符号：")

	fmt.Scanln(&fh)

	switch fh {
	case "+":
		fmt.Printf("%d + %d = %d",num1,num2,num1+num2)
	case "-":
		fmt.Printf("%d - %d = %d",num1,num2,num1-num2)
	default:
		fmt.Println("没有这个操作")

	}


}

