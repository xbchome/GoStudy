package main

import (
	"fmt"
)

func saveIntData(intData chan int) {
	for i:=1; i<=8000;i++  {
		intData <- i
	}
	fmt.Println("数据填充完成")
	close(intData)
}

func Prime(intData,primeData chan int,isChan chan bool) {
	for   {
		v,ok := <-intData
		if !ok {
			break
		}
		for i:=2;i<v ;i++  {
			if v%i== 0 {
				break
			}else if v==(i+1) {
				primeData <- v
			}

		}
	}

	isChan<- true

}
func main() {
	intData :=make(chan int,8000)
	primeData := make(chan int,2500)
	isChan := make(chan bool,4)

	go saveIntData(intData)
	for i:=1;i<=4 ;i++  {
		go Prime(intData,primeData,isChan)
	}
go func() {
	for i:=1;i<=4;i++ {
		fmt.Println(<-isChan)
	}
	close(primeData)
}()


	for v:=range primeData{
		fmt.Println(v)
	}
	fmt.Println("主程序退出")

}
