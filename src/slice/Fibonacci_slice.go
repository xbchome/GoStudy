package main

import "fmt"

func main() {
	fmt.Println(fSlice(10))
	//var arr = new([2]int)
	//arrs(arr)
	//fmt.Println(arr)
}

// 返回对应长度的斐波那契数列 切片
func fSlice(len int)([]int) {
	slice1 := make([]int,len)

	for i := 0;i<len; i++{
		if i <=1{
			slice1[i] = 1
		}else {
			slice1[i] = slice1[i-1] + slice1[i-2]
		}
	}

	return slice1
}

func  arrs(arr [2]int) {
	arr[1] = 1
	arr[0] = 3
}