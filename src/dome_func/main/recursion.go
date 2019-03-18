package main

import "fmt"

func main() {
	fmt.Println(f(1))
}


func f(num int) int {
	// 如果是第10天返回 1
	if num == 10 {
		return 1
	}

	return (f(num+1)*1)+2
}