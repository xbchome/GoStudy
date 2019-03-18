package main

import (
	"fmt"
	"math/rand"
	"time"
)

var num int
func main() {
	rand.Seed(time.Now().Unix()) // 生成随机种子
	num = rand.Intn(100)
	game()

}


func game() {
	var numb int
	for i:=1;i<=10;i++  {
		fmt.Printf("请输入数字:")
		fmt.Scanln(&numb)
		
		if numb == num {
			switch i {
			case 1:
				fmt.Println("你是个天才")
				return
			case 2,3:
				fmt.Println("你很聪明")
				return
			case 4,5,6,7,8,9:
				fmt.Println("一般般")
				return
			case 10:
				fmt.Println("还行")
				return
			}
		}

		if numb > num {
			fmt.Printf("你输入的数大了 \n")
		} else{
			fmt.Printf("你输入的数小了 \n")
		}
	}

	fmt.Println("辣鸡")
}
