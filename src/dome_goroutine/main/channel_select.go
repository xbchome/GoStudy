package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int,10)
	strChan := make(chan string,10)


	go func() {

		for i:=1;i<=10;i++{
			fmt.Println("l")
			intChan <- i
			time.Sleep(time.Nanosecond *10)
		}

		close(intChan)
	}()

	for i:=1;i<=10;i++{
		strChan <- "Hello " + fmt.Sprintf("%d",i)
	}
	close(strChan)
	look := true
	for look {
		select {
		case v:= <-intChan:
		fmt.Println("v:",v)
		case v:=<-strChan:
			fmt.Println("strV:",v)
		default:
			if _,ok := <-intChan; !ok {
				look = false
			}
			//close(intChan)

		}

	}

	//for v:= range intChan {
	//	fmt.Println("for v:",v)
	//}
}

