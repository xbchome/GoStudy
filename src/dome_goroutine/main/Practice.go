package main

import "fmt"

type Sums struct {
	Res int
	num int
}

func Wdata(data1 chan int) {
	for i:=1;2000>=i ;i++  {
		data1 <- i
	}

	close(data1)
}

func Rdata(data2 chan Sums,data1 chan int) {
	for v:= range data1 {
		res := add(v)

		data2<- Sums{v,res}
	}

	close(data2)
}

func add(n int) int {
	res:=0
	for i:=1;n>=i;i++{
		res+=i
	}

	return res
}
func main() {
	data1 := make(chan int,2000)
	data2 := make(chan Sums,2000)

	for i:=8;i<=8;i++{
		go Rdata(data2,data1)
	}

	go Wdata(data1)

	for v:= range data2 {
		fmt.Printf("res[%v]:%v\n",v.Res,v.num)
	}

}
