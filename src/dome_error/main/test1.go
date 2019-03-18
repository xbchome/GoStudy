package main

import "fmt"

func main() {
	meun()
}


func meun() {
	var y int
	var ye int
	var ri int
	for {
		fmt.Println("请输入年份:\n")
		fmt.Scanln(&y)
		fmt.Println("请输入月份:\n")
		fmt.Scanln(&ye)

		if  ye>12 || ye<1{
			fmt.Println("月份输入有误")
			continue
		}
		fmt.Println("请输入日期:\n")
		fmt.Scanln(&ri)

		fmt.Printf("你输入的是%d年 %d月 %d日",y,ye,ri)

	}
}
