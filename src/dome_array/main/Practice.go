package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	随机生成10个数 保存到数组  倒序打印  求平均值、最大值、最大值下标、查找里面是否存在 55
*/
func main() {
	rand.Seed(time.Now().Unix()) //
	var arrs [10]int
	var len1 = len(arrs)
	for i:=0;i<len(arrs);i++{
		arrs[i] = rand.Intn(100)+1
	}

	fmt.Println("==================arrs===========")
	fmt.Println(arrs)
	for i:=0;i<len(arrs)/2;i++{
		temp := arrs[i]
		arrs[i] = arrs[len(arrs)-1-i]
		arrs[len(arrs)-1-i] = temp
	}
	fmt.Println("==================arrs倒序后===========")
	fmt.Println(arrs)

	var sum int
	var max int = 0
	var max_ix int
	for i:=0;len1>i;i++{
		sum += arrs[i]
		if max<arrs[i] {
			max = arrs[i]
			max_ix = i
		}
	}
	fmt.Println("==================arrs 平均值===========")

	fmt.Println("平均值:",float64(sum/len1))
	fmt.Println("最大值:",max)
	fmt.Println("最大值xb:",max_ix)
	fmt.Println("==================arrs 找 55===========")
	//binaryFind(arrs,55)
	orderFind(arrs,55)
}
func orderFind(arrs[10]int,sea int) {
	var index int = -1
	for i:=0; len(arrs)>i ;i++  {
		if arrs[i] == sea {
			index = i
			break
		}
	}
	if index > -1 {
		fmt.Println("找到了")
	} else{
		fmt.Println("找不到")
	}

}

func binaryFind(arrs[10]int,sea int) {
	leftIndex := 0
	rightIndex := len(arrs)-1
	var midd int
	for  {

		if leftIndex > rightIndex {
			fmt.Println("找不到")
			break
		}
		midd = (leftIndex+rightIndex)/2
		if arrs[midd] > sea {
			rightIndex = midd-1
		}

		if arrs[midd] < sea {
			leftIndex = midd +1
		}

		if arrs[midd] == sea {
			fmt.Println("找到了：",sea)
			break
		}
	}
}
