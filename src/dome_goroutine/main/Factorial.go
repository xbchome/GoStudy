package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	nums = make(map[int]uint,10)
	look sync.Mutex //fatal error: concurrent map writes 解决方案一 添加互斥锁
	endTime float64
)

func test2() {
	startTime:= time.Now().UnixNano()
	num := make(map[int]int)
	for i:=1;20>=i ;i++  {
		res := 1
		for j:=1;i>=j ;j++  {
			res*=j
		}
		num[i] = res
	}

	fmt.Println(num)
	endTimes:= time.Now().UnixNano()

	fmt.Printf("\n\n test2end: %v start: %f  = %f",endTimes,float64(startTime),(float64(endTimes)-float64(startTime))/1000000000)
}

func test(n int) {
	res :=1
	for i:=1;i<=n ;i++  {
		res*=i
	}
	//fmt.Println("hell")

	look.Lock()
	nums[n] = uint(res)
	endTime = float64(time.Now().UnixNano())
	look.Unlock()

}

// fatal error: concurrent map writes
func main() {
	startTime:= time.Now().UnixNano()
	for i:=1;i<=20 ;i++  {
		go test(i)
	}


	time.Sleep(time.Second *2)
	fmt.Printf("\n TEST2 =====\n end: %f start: %f  = %f",endTime,float64(startTime),(endTime-float64(startTime))/1000000000)
	fmt.Println(nums)

	test2()
	//for i,v:=range nums {
	//	fmt.Printf("i:%v v:%v\n",i,v)
	//}

}
