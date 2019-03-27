package main

import (
	"fmt"
	"runtime"
	"time"
)

func Hello() {
	for i:=0;i<12;i++{
		fmt.Println("Hello world")
		time.Sleep(time.Second)
	}
}
func main() {
	go Hello()

	for i:=0;i<9;i++{
		fmt.Println("Hello Golang")
		time.Sleep(time.Second)
	}

	num := runtime.NumCPU()
	fmt.Println(num)
}
