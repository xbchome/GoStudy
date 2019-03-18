package main

import (
	"fmt"
)

func main() {
	var char [26]byte

	for i,_ := range char {
		char[i] = byte('A'+ i)
	}

	for i,v:=range char {
		fmt.Printf("char[%d]:%c\n",i,v)
	}
}

