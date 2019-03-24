package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file,err := os.Open("G:/NutCloud/code/go/src/FileOperation/main/test.txt")

	if err != nil {
		fmt.Println("文件读取失败",err)
	}

	defer file.Close()  // 关闭文件流

	reader := bufio.NewReader(file) // 创建带缓冲区的文件读取

	for   {
		str,err := reader.ReadString('\n')
		if err ==io.EOF {
			break
		}

		fmt.Print(str)
	}

	// 一次性读取文件  适用于小文件
	bytes,err := ioutil.ReadFile("G:/NutCloud/code/go/src/FileOperation/main/test.txt")

	if err != nil {
		fmt.Println("err=",err)
	}

	fmt.Println(string(bytes))
}
