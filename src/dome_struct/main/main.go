package main

import (
	"dome_struct/factory"
	"fmt"
)

type Dog struct {
	name string
	age int
	arrs [3] int // 数组
	per *int // 指针
	slices []int
	mapl map[int]string

}

type Cat struct {
	Name string
	Age int
	Sex string
}

func main() {
	var cat Cat

	cat.Name = "猇亭区"
	cat.Age = 1
	cat.Sex = "男"
	fmt.Println(cat)
	fmt.Printf("name:%s",cat.Name)
	fmt.Println("==================================")
	var dog Dog
	fmt.Println(dog)
	var per int = 8
	dog.arrs = [3]int{1,2,3}
	dog.slices = dog.arrs[:]
	dog.per = &per
	fmt.Println(dog)
	var stu  factory.Student = factory.Student{"xbc",12,"X"}
	fmt.Println(stu)
}
