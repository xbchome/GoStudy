package main

import "fmt"

type point struct {
	x,y int
}

type rect struct {
	up,down point
}

type rect2 struct {
	up,down *point
}

func main() {
	r1 := rect{point{1,2},point{3,4}}

	fmt.Printf(" r1.up.x 地址:%v\n r1.up.y 地址:%v\n r1.down.x 地址:%v\n r1.down.y 地址:%v \n",&r1.up.x,&r1.up.y,&r1.down.x,&r1.down.y)
	r2 := rect2{&point{1,2},&point{3,4}}
	fmt.Println(r2)

	fmt.Printf("%p %p",&r2.down.x,&r2.down.y)
}