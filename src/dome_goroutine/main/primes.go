package main

import (
	"fmt"
)

func writeInt(intChan chan int) {
	for i:=1;i<=8000 ; i++  {
		intChan<-i
	}
	fmt.Println("写入完成")
	close(intChan)
}

func primeNum(intChan,primes chan int,exitChan chan bool) {


	for   {
		v,ok := <-intChan
		if !ok {
			break
		}
		for i:=2;i<v ;i++  {
			if v%i== 0 {
				break
			}else if v==(i+1) {
				primes<-v
			}

		}
	}
	//close(primes)
	exitChan<- true
}
func main() {
	intChan := make(chan int,1000)
	primes :=make(chan int,2500)
	exitChan:=make(chan bool,4)

	go writeInt(intChan)
	for i:=0;i<=4 ;i++  {
		fmt.Println(i)
		go primeNum(intChan,primes,exitChan)
	}


		for i:=0;i<=4;i++{
			fmt.Println("exit:",<-exitChan)
		}

		close(primes)

	fmt.Println("主进程完成")
	fmt.Print(">>")
		go func() {}()
}