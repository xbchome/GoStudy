package main

import (
	"fmt"
	"time"
)

func main() {
	d1()
	fmt.Printf("\n")
	d2(4,4)

	d3()
	fmt.Printf("\n")
	d4()

	x9()

	fmt.Println(time.Now().Unix())
}

func d1() {
	for i:=1;i<=5 ;i++  {

		for k := 0; k<i; k++{
			fmt.Printf("*")
		}

		fmt.Printf("\n")
	}
}


func d2(c int ,x int) {
	for i := 1; i<= c ; i++{
		for k:=1;x>= k; k++{
			fmt.Printf("*")
		}
		fmt.Printf("\n")

	}
}

func d3() {

	for i := 1; i <=5  ; i++  {

		for j:=1; j<= 5-i;j++  {
			fmt.Printf(" ")
		}

		for k:= i*2-1; k>=1 ; k--  {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

}


func d4() {
	var level int = 5
	for i :=1 ; level>=i ; i++  {
		// 打空格
		for k := 1; k<=level-i;k++{
			fmt.Printf(" ")
		}

		// 打*

		for j := 1; j<= i*2-1; j++{
			if j== 1 || j == i*2-1 || i==level{
				fmt.Printf("*")
			}else{
				fmt.Printf(" ")
			}

		}

		fmt.Printf("\n")
	}
}


func x9() {

	for i := 1; i<=9; i++{

		for k := 1; k<=i;k++{
			fmt.Printf("%v * %v = %v \t",k,i,i*k)
		}
		fmt.Printf("\n")

	}
}
