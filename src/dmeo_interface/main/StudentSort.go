package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Student struct {
	Name string
	Age int
}

type Students []Student

func (stu Students) Len() int {
	return len(stu)
}

func (stu Students) Less(i,j int) bool {
	return stu[i].Name < stu[j].Name
}

func (stu Students) Swap(i,j int) {
	stu[i], stu[j] = stu[j], stu[i]
}
func main() {
	var stus Students
	rand.Seed(time.Now().Unix())
	for i:=1;10>i;i++{
		stus = append(stus,Student{fmt.Sprintf("人名:%d",i+rand.Intn(50)),rand.Intn(100)+1})
	}
	fmt.Println("=====================")
	fmt.Println(stus)
	fmt.Println("=====================")
	sort.Sort(stus)
	fmt.Println(stus)


}
