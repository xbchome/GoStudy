package main

import "fmt"

type person struct {
	Name string
}

func (p person) eatName() {
	fmt.Println("我是:",p.Name)
}

func (p *person) setName() {
	p.Name = "xbc"
}

type rads struct {
	r float64
}

func (r rads) mianji() float64 {
	return r.r*r.r *3.14
}

func main() {
	person := person{"杨过"}
	person.eatName()
	fmt.Println("===================")
	person.setName()
	person.eatName()
	fmt.Println("====================")



}
