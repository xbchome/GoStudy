package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().UnixNano())

	fmt.Println(time.Now().Unix())

	fmt.Println(time.Now().UTC())

	fmt.Println(time.Now().UnixNano())
	j2()

	d()

}


func d() {
	 i := 0

	for ; i<3; i++  {

		fmt.Println(i)
	}
}

func j2() {

	for i:=1; 100>=i ;i++  {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}
}