package main

import "fmt"

type Student struct {
	Name string
	Gender string
	Age int
	id int
	score float64
}

func (stu Student) Say() string {
	str := fmt.Sprintf("name:%v Gender:%v age:%v id:%v score:%v",stu.Name,stu.Gender,stu.Age,stu.id,stu.score)
	return str
}
func main() {
	stu := Student{"杨过","男",18,1,59}

	fmt.Println(stu.Say())
}
