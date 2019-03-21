package bin

import "fmt"

type Student struct {
	Name string
	Age int
	Sex string
	Grade float64
}

func (stu *Student) ShowInfo() {
	fmt.Printf("姓名:%v 年龄:%v 成绩：%v\n",stu.Name,stu.Age,stu.Grade)
}

