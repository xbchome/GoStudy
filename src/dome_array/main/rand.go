package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arrs [5]int
	rand.Seed(time.Now().Unix())

	for i:=0; i<5;i++  {

		arrs[i] = rand.Intn(100)
	}
	fmt.Println(arrs,"\n")
	for i:=len(arrs);i>0 ;i--  {
		fmt.Printf("arrs[%d]:%d\n",i-1,arrs[i-1])
	}

	 lens :=len(arrs)
	var temp int
	 for i:=0; i<lens/2;i++{
	 	temp= arrs[lens-1-i]
	 	fmt.Printf("temp%d:%d\n",i,temp)
		 arrs[lens-1-i] = arrs[i]
		 arrs[i] = temp
	 }

	 fmt.Println("==========================\n")
	 fmt.Println(arrs)



}