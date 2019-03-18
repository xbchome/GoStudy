package main

import "fmt"

func main() {
	// 数组的定义

	var nullArr1 [3]int
	var nullArr2 [2]int = [2]int{1,3}
	var nullArr3 = [3]int{1,2,3}

	var nullArr4 =[...]int{1,2,3}
	fmt.Printf("1:%v\n2:%v\n 3:%v \n 4:%v",nullArr1,nullArr2,nullArr3,nullArr4)

	for v,k:= range nullArr3 {
		fmt.Printf("v:%d k:%d \n",v,k)
	}

	var i int = 0
	for ;i<len(nullArr3);i++ {
		fmt.Printf("nullArr3[%d]\n",nullArr3[i])
	}

	fmt.Println("||=================|===================||") //

	fmt.Println("||=================|===================||") //

	fmt.Printf("%c",'a'+1)


}
