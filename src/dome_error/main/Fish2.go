package main

import (
	"fmt"
	"time"
)

func main() {
	Fish("1990-01-06 12:00:00")
}

func Fish(strtime string) {

	str :="2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time1,err := time.ParseInLocation(str,strtime,loc)
	time2,_ := time.ParseInLocation(str,"1990-01-01 00:00:00",loc) // 将日期转换为时间对象
	//fmt.Println(time1.Unix(),"\n")
	fmt.Println(err,"\n")
	time3 := time1.Sub(time2)  // 计算时间差
	//fmt.Println(time1.Sub(time2))

	sun := int(time3.Hours()/24)+1 // 总的距离天数
	fmt.Println(sun)
	switch sun%5 {
	case 0,4:
		fmt.Println("晒网")
	case 1,2,3:
		fmt.Println("打鱼")
	}

	//fmt.Printf("还有 %d 天",int(time3.Hours()/24)+1)



}
