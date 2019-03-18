package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("快")
	start := time.Now()
	time.Sleep(time.Millisecond * 500)
	end := time.Now()
	fmt.Println("0.5秒")
	fmt.Println((float64(end.UnixNano())-float64(start.UnixNano()))/1000000000)
	//fmt.Println(float64(end)-float64(start))
}
