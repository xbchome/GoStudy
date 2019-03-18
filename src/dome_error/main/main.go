package main

import (
	"fmt"
)

func main() {
	test(12,4)

	fmt.Println("后续代码\n")

	test(12,0)

	fmt.Println("后续代码2\n")
}

func test(num int , num2 int) {
	//
	defer func() {
		err := recover()

		if  err != nil{
			fmt.Printf("err:%v \n",err)
		}
	}()

	fmt.Printf("s:%d\n",num/num2)

}
