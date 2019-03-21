package main

import (
	"dmeo_Object/bin"
	"fmt"
)

type A struct {
	name string
	age int
}

type B struct {
	A
}


func main() {
	var pup bin.Pupil = bin.Pupil{bin.Student{"杨建",10,"男",12.00}}
	pup.Exam()
	pup.Student.ShowInfo()

	var b = B{A{"x",1}}

	fmt.Println(b.age)
}
