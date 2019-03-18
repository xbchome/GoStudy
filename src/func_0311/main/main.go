package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	str := "xbc5好"

	fmt.Println("len:",len(str)) //长度为7  中文占3个字节

	//******

	str2 := "o123"

	n,e := strconv.Atoi(str2)

	fmt.Println("\n************\n")
	fmt.Println("n:",n,"e:",e)
	fmt.Println("\n************\n")

	str3 :="我是中国人"

	str4 := []rune(str3) // 如果要遍历带中文的字符串 需要先用 rune 转换

	for i:=0;len(str4)>i ;i++  {
		//fmt.Println("str4:",str[i],"\n")
		fmt.Printf("str%d: %c \n",i,str4[i])
	}

	fmt.Printf("\n************\n")
	str5 := []byte("xbc")

	for i:=0;len(str5)>i ;i++  {
		//fmt.Println("str4:",str[i],"\n")
		fmt.Printf("str%d: %v \n",i,str5[i])
	}

	fmt.Printf("\n************\n")

	info := strings.Contains("aaaxbcss","xbc")

	fmt.Printf("%v",info)

	fmt.Printf("\n************\n")

	fmt.Printf("num:%d",strings.Count("aaaxbcss","xbc"))


	fmt.Printf("\n************\n")

	fmt.Printf("replace %s",strings.Replace("hello go go go","go","go语言",1)) // 字符替换

	fmt.Printf("\n************\n")

	now := time.Now()
	ni,y,r:=now.Date()
	fmt.Printf("年%d 月%d 日 %d",ni,y,r)




}
