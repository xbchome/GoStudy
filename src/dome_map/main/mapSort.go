package main

import (
	"fmt"
	"sort"
)

func main() {
	var	maps = map[int]int{
		1:23,
		2:22,
		3:32,
		5:42,
		4:52,
	}

	fmt.Println(maps)

	var keys []int

	for k,_ := range maps {
		keys = append(keys,k)
	}
	fmt.Println("==================前==============")
	fmt.Println(keys)
	fmt.Println("==================后==============")
	sort.Ints(keys)
	fmt.Println(keys)
	for _,v:= range keys {
		fmt.Printf("maps[%d]:%d",v,maps[v])
	}
}
