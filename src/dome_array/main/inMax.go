package main

import "fmt"

func main() {
	arrs := [10]int{1,2,3,4,5,6,12,9,9,10}
	var top int = 0
	var index int =0
	for i,v := range arrs {

		if top < v {
			top = v
			index=i
		}
	}

	fmt.Printf("index:%d value:%d",index,top)

	var sum int
	var p float64
	for i,v :=range arrs {
		sum +=v
		if (i+1)==len(arrs) {
			p = float64(sum)/float64(len(arrs))
		}
	}

	fmt.Printf("和:%d 平均:%0.2f",sum,p)

}
