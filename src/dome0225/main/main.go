package main

import (
	"dome0225/fors"
	"dome0225/xbc"
	"fmt"
)
func main() {
	fmt.Println("main......")
	fors.D2()
	fors.D3(12)
	i := 1
	for ;10>i ;i++  {
		fmt.Println("你好")
	}
	xbc.X1()

	var sum = [3]int{1,3,3}

	xbc.X3(sum)

	fmt.Println(sum)




}
