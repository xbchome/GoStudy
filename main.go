package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello")
	Hello("word")
	f := x(".jpg")
	fmt.Println("x:",f("xbc.jp"))

	nums := [3]int {1,1,3}

	xx(&nums)
	fmt.Println(nums)


}

func Hello(string2 string) (x int)  {
	fmt.Println(string2)
	x=1
	return x
}


func x(hz string) (func(name string) string) {
	return func (name string) string {
		if strings.HasSuffix(name,hz) {
			return name
		}
		return name+hz
	}
}



func xx(nuns *[3]int) {
	nuns[1] = 2
}

