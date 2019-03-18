package main

import "fmt"

func main() {
	var nums [4][4]int
	numsLen := len(nums)
	for i:=0;numsLen>i ;i++  {
		for j:=0;numsLen>j;j++{

			fmt.Printf("请输入第%d行第%d个数\n",i,j)
			fmt.Scanln(&nums[i][j])
		}

	}
	fmt.Println("=============================")
	fmt.Println(nums)
	for i:=0;(numsLen/2)>i;i++{
		tmep := nums[i]
		nums[i] = nums[numsLen-1-i]
		nums[numsLen-1-i] = tmep
	}
	fmt.Println("=============================")
	fmt.Println(nums)
}
