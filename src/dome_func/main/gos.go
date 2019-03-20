package main

import (
	"fmt"
	"strconv"
)

type Strings struct {
	Name string
}


func main() {
	var n int = 1

	s:= strconv.Itoa(n)

	fmt.Println(s)
	fmt.Printf("sT:%T nt:%T",s,n)
	fmt.Printf("\n %d",'a')
	fmt.Printf("\n A: %c",65)
	fmt.Printf("")
}

func Two() {
	fmt.Println("new ")
}