package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
		var now = time.Now()

	fmt.Println(now)

	fmt.Printf("\n*********************\n")

	fmt.Printf("time: %v type:%T",now.Format("2006/01/02 15:04:05"),now.Format("2016-01-01 00:00:00"))
	fmt.Printf("\n*********************\n")

		start := time.Now().Unix()
		//srot()
		end := time.Now().Unix()

		fmt.Printf("耗时:%v",end-start)
	fmt.Printf("\n*********************\n")

		now1 := time.Unix(1551285314,0)
		now2 := time.Unix(1552295314,0)
		x:=now2.Sub(now1)
		fmt.Println(now2.Sub(now1))
	fmt.Printf("\n*********************\n")
		fmt.Println(x)
	fmt.Printf("\n*********************\n")
		fmt.Println("天:",int(x.Hours())/24,"\n")
		fmt.Println("天:",int(x.Hours())%24,"\n")

		fmt.Println("now1时间：",now1.Format("2006-01-02"))
	fmt.Printf("\n*********************\n")
		fmt.Println(time.Unix(time.Now().Unix(),0))
}

func srot () {
	var str string
	for  i := 0;100000>i ;i++  {
		str += "h"+ strconv.Itoa(i)
	}
}
