package main

import "fmt"

func main() {
	// 猴子吃桃 问题
	fmt.Println(hz(10))

	fmt.Println(hz2(1))
}

func hz(level int) int {
	if level == 1 {
		return 1
	}
	return  (hz(level-1)+1)*2
}

func hz2(level int) int {
	if level ==10 {
		return 1
	}
	// 第一天的桃子等于 前一天的桃子数量+1 *2
	return (hz2(level+1) +1)*2

}


