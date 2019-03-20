package bin

import "fmt"

type person struct {
	Name string
	age int
	sex string
}


func NewPerson(name string) *person {
	return &person{"xbc",0,""}
}

func (per *person) SetAge(num int) {
	if num <0 || num>120 {
		fmt.Println("年龄输入有误！")
		return
	}else {
		per.age = num
	}
}

func (per *person) GetAge() int {
	return (*per).age
}

func (per *person) Setsex(num string) {
	if num != "男" && num!="女" {
		fmt.Println("年龄输入有误！")
		return
	}else {
		per.sex = num
	}
}

func (per *person) Getsex() string {
	return (*per).sex
}

