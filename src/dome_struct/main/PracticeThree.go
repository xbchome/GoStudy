package main

import (
	"fmt"
)

type  MethodUtils struct {
}

func (m MethodUtils) Print() {
	for i:=1;10>=i ;i++  {
		for j:=1;8>=j ;j++  {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func (Me MethodUtils) Print2(m,n int) {
	for i:=1;m>=i ;i++  {
		for j:=1;n>=j ;j++  {
			fmt.Printf("*")
		}

		fmt.Printf("\n")
	}
}
func (Me MethodUtils) IsEven(m int) {
	if m%2 == 0 {
		fmt.Println("是偶数")
	}else{
		fmt.Println("是奇数")
	}
}

func (Me MethodUtils) ArrayTwo(nums[3][3]int) {
	len1 := len(nums)
	len2 := len(nums[0])
	var temp int
	for i:=0;len1>i ;i++  {
		for j:=0;len2>j ;j++  {
			if j> i {
				temp = nums[i][j]
				nums[i][j] = nums[j][i]
				nums[j][i] = temp
			}
		}
	}

	fmt.Println(nums)
}

type Calcuator struct {
	Num1 float64
	Num2 float64
}

func (Ca *Calcuator) Add() float64 {
	return  Ca.Num1 + Ca.Num2
}

func (Ca *Calcuator) Minus() float64 {
	return  Ca.Num1 - Ca.Num2
}

func (Ca *Calcuator) Multiply() float64 {
	return  Ca.Num1 * Ca.Num2
}

func (Ca *Calcuator) Removal() float64 {
	if Ca.Num2 != 0 {
		return  float64(Ca.Num1) / float64(Ca.Num2)
	}

	return 0
}

func (Ca *Calcuator) Compute(m string) float64 {
	num :=0.0
	switch m {
	case "+":
		num = Ca.Num1 +Ca.Num2
	case "-":
		num = Ca.Num1 -Ca.Num2
	case "*":
		num = Ca.Num1 * Ca.Num2
	case "/":
		num = Ca.Num1 / Ca.Num2
	default:
		fmt.Println("操作有误")
	}
	return num
}
func main() {
	var m MethodUtils
	m.Print()

	fmt.Println("=====================")
	m.Print2(5,5)

	m.IsEven(3)

	fmt.Println("====================")
	ca :=Calcuator{4,1}

	fmt.Println("除",ca.Removal())
	fmt.Println("加",ca.Add())
	fmt.Println("-",ca.Minus())
	fmt.Println("*",ca.Multiply())

	fmt.Println(ca.Compute("/"))

	fmt.Println("===================")
	nums := [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(nums)
	m.ArrayTwo(nums)

	fmt.Println("=================")
}
