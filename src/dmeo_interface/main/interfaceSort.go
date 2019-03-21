package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

type Person struct {
	Name string
	Age int
}

type persons []Person // 结构体切片 自定义结构

func (per persons) Len()int {
	return len(per)
}

func (per persons) Less(i,j int) bool {
	return per[i].Age < per[j].Age
}

func (per persons) Swap(i,j int) {
	temp:=per[i]
	per[i] = per[j]
	per[j] = temp
}

func main() {
	var pers persons

	for i:=0;10>= i; i++{
		//per :=
		pers = append(pers,Person{"李白"+strconv.Itoa(i),rand.Intn(100),})
	}
	fmt.Println(pers)

	sort.Sort(pers)
	fmt.Println(pers)
}