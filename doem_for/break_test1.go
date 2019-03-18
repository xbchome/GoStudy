package main

import "fmt"

func main (){
	var name string
	var password string
	var count int = 3
	for{
		fmt.Println("请输入用户名1:")
		fmt.Scanln(&name)
		fmt.Println("请输入密码:")
		fmt.Scanln(&password)

		if name =="张无忌" && password == "888" && count>=3{
			fmt.Println("登录成功！")
			break
		}else
		{
			count--
			fmt.Println("还有",count,"机会")
		}
	}
}
