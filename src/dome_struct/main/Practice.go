package main

import "fmt"

type dog struct {
	name string
	age int
	color string
}
// 结构体 初始化的方式
func main() {
	// 1
	fmt.Println("===============一===========")

	var dog1 dog
	dog1.name = "杨康"
	dog1.age = 12
	dog1.color = "黑"
	fmt.Println(dog1)
	// 2
	fmt.Println("===============二===========")
	dog2 := dog{name:"yk",age:1,color:"黑"}
	fmt.Println(dog2)
	// 3 new
	fmt.Println("===============3===========")
	dog3 := new(dog)
	(*dog3).name = "xx"
	fmt.Println(dog3)

	// 4
	fmt.Println("===============四===========")
	var dog4 *dog = &dog{}
	var dog5 *dog = dog4
	dog5.name = "大黄"
	dog5.color = "黄"
	fmt.Println(dog4)

}
