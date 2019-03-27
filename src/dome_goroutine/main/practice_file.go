package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func writeDataToFile() {
	//var strSlice []int
	//
	//for i:=1;1000>=i ;i++  {
	//	strSlice = append(strSlice,i)
	//}
	//
	//jsonStr,err := json.Marshal(strSlice)
	//if err!= nil {
	//	fmt.Println("json 出错")
	//	return
	//}

	file,err := os.OpenFile("D:/sort.json",os.O_CREATE|os.O_WRONLY,0755)
	if err != nil {
		fmt.Println("打开文件流失败")
		return
	}
	defer file.Close()

	for i:=1; 1000>= i;i++{
		str:= fmt.Sprintf("%d\n",i)
		file.WriteString(str)
	}

	//writer := bufio.NewWriter(file)
	//
	//writer.WriteString(string(jsonStr))
	//writer.Flush()
}

func readDateSort(isChan chan bool) {
	file,err :=os.Open("D:/sort.json")
	if err != nil {
		fmt.Println("打开文件流出错")
	}
	defer file.Close()
	read := bufio.NewReader(file)
	var intSlice []int
	for{
		str,err:=read.ReadString('\n')
		if err ==io.EOF {
			fmt.Println("e:",err)
			break
		}
		ints,err:=strconv.Atoi(strings.Replace(str,"\n","",-1))
		if err != nil {
			fmt.Println("err:",err)
		}
		intSlice = append(intSlice,ints)
	}

	fmt.Println(intSlice)
	isChan<- true
	close(isChan)

}

func main() {
	isChan := make(chan bool,1)
	//isChan<- false
	go writeDataToFile()
	go readDateSort(isChan)


		i:= <-isChan
		fmt.Println(i)

		fmt.Println("isChan")


	//time.Sleep(time.Second)
}
