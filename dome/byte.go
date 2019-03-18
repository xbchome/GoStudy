package main

import (
	"fmt"
)

func main() {

	var c1 int = '被'
	var c2 int = '呗'

	fmt.Printf("c1 = %c \n c1=%d",c1,c1)
	fmt.Printf("c2 = %c \n c2=%d",c2,c2)

	i  := 3 + 'a'
	fmt.Println("i = %d \n a=%d",i,'a')

	var s  = [3]int {1,2,3}
	sort(s)
	switch v := 12;{
	case v == 1:
		fmt.Println("一")
	case v == 12:
		fmt.Println("fmt")
		fmt.Println("fmt")
	default:
		fmt.Println("没有")
	}

	n := 1
	switch n{
	case 1:
		fmt.Println("110")
	default:
		fmt.Println("121")
	case 2:
		fmt.Println("Hello World") //
	}
}

func sort( nums [3]int ) {
	fmt.Println(nums)
}

func qy() {
	fmt.Println("恶")
}




