package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file,err := os.OpenFile("D:/Codes/code/go/src/FileOperation/main/test.txt",os.O_RDWR|os.O_APPEND,0)

	if err!= nil {
		fmt.Printf("err:%v",err)
	}
	defer file.Close() // 关闭文件流
	writer:= bufio.NewWriter(file)

	writer.WriteString("全球\n")




	reader := bufio.NewReader(file)
	for   {
		str,err := reader.ReadString('\n')
		fmt.Println(str)
		 if err != nil {
		 	fmt.Println(err)
		 	break
		 }
	}

	writer.Flush()  // 将缓冲区中的内容写入文件



}
