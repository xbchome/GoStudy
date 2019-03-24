package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file,err := os.OpenFile("G:/NutCloud/code/go/src/FileOperation/main/testXBC.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND,0755)

	if err!= nil {
		fmt.Println("err:=",err)
	}

	defer file.Close()

	// 创建带缓冲区的写入对象

	writes := bufio.NewWriter(file)
	str := "床前明月光\n"
	for i:=0;5>i ;i++  {
		writes.WriteString(str)
	}
	// 将缓冲区内容写入文件
	writes.Flush()

}

// 判断文件是否存在
func PathExists(path string)(bool,error) {
	_,err := os.Stat(path)
	if err == nil { // 说明文件存在
		return true,nil
	}

	if os.IsNotExist(err) {
		return false,nil
	}

	return false,err
}
