package xbc

import "fmt"
func init() {
	fmt.Println("我是init")
}
func X1() {
	fmt.Println("我是x1")
}

func x2() {
	fmt.Println("我是x2")
}

func X3(sun [3]int) {
	sun[1]= 5
}