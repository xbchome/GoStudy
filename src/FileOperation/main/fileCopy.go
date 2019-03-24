package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(src2 string,str string) (int64,error){
	reader,err := os.Open(str)

	if err != nil {
		fmt.Println("打开文件失败 error = ",err)
	}
	defer reader.Close()
	read := bufio.NewReader(reader)

	file,err:=os.OpenFile(src2,os.O_WRONLY|os.O_CREATE,0755)
	if err != nil {
		fmt.Println("打开文件失败 error = ",err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	return io.Copy(write,read)
}
func main() {
	CopyFile("E:/123.jpg","G:/psbCA51Q53E.jpg")
}