package main

import "fmt"

func main() {
	var nums [3][4] int = [3][4]int{{2,2,2,2},{2,2,2,2},{2,2,2,2}}

	for i:=0;len(nums)>i ;i++  {
		fmt.Println(nums[i])
	}

	for is,vs:= range nums {

		for i,_:= range vs {
			if is == 0 || i == 0 || is == len(nums)-1 || i==len(vs)-1{
				nums[is][i] = 0
			}

		}
	}
	fmt.Println("================")
	for i:=0;len(nums)>i ;i++  {
		fmt.Println(nums[i])
	}
}
