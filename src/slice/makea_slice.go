package main

import "fmt"

func main() {
	var slice1 []int = make([]int,7)

	for i := 0; i < len(slice1); i++{
		slice1[i] = 5 * i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d \n",i,slice1[i])
	}

	fmt.Printf("Slice len %d \n",len(slice1))
	fmt.Printf("Slice capacity %d \n",cap(slice1))
}
