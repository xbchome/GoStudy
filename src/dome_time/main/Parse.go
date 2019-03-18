package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	time1,_ := time.Parse("2006-01-02","2019-12-12")
	time2,_ := time.Parse("2006-01-02","2019-12-13")
	fmt.Println(int(time2.Sub(time1)))

	time3 :=time.Unix(1500124,0)

	fmt.Println("time3:",time3.Format("2006-01-02"))

	fmt.Printf("\n lastIndex %v",strings.LastIndex("xbcbb","xbc"))
	fmt.Printf("\n len %v",len("xbc"))
	

}
