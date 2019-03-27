package main

import (
	"fmt"
)

func test3() {
	cho := make(chan int,4)
	cho<- 10
	cho<- 20
	fmt.Println(<-cho)
	fmt.Println(<-cho)
}

func readData(dataChan chan int,isChan chan bool) {
	for {
		v,ok:= <-dataChan
		if !ok {
			break
		}
		fmt.Println(v)
	}
	isChan <-false
	close(isChan)
}

func writeData(dataChan chan int) {

	for i:=1;50>=i ;i++  {
		dataChan <- i
	}

	close(dataChan)
}

func main() {
	dataChan := make(chan int,50)
	isChan := make(chan bool,1)

	go readData(dataChan,isChan)
	go writeData(dataChan)

	//for   {
	//	v,ok:=<-dataChan
	//	if !ok {
	//		break
	//	}
	//
	//	fmt.Println("main v:",v)
	//}

	for v:= range isChan {
		fmt.Println(v)
	}

}