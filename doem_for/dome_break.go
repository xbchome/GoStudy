package main

import "fmt"

func main() {
	//count :=0
	//rand.Seed(time.Now().UnixNano())
	//for {
	//
	//
	//	n := rand.Intn(100)
	//	count++
	//	fmt.Println(n)
	//	if n == 99 {
	//		break
	//	}
	//}
	//
	//fmt.Println("count:",count)

	var name string
	var password string
	var count int = 3
	for{
		fmt.Println("请输入用户名:")
		fmt.Scanln(&name)
		fmt.Println("请输入密码:")
		fmt.Scanln(&password)

		if name =="张无忌" && password == "888" && count>=1{
			fmt.Println("登录成功！")
			break
		}else
		{
			if count <=0 {
				break
			}
			count--
			fmt.Println("还有",count,"机会")
		}
	}
}
