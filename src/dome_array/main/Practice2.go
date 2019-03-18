package main

import "fmt"

// 在顺序数组中插入一个数 使数组顺序不变
func main() {
	var nums [10]int = [10]int{1,2,3,4,5,6,7,8,10,12}
	var nums2 [len(nums)+1] int
	var num int = 9
	lenth := len(nums)
	for i:=0; len(nums2) >i;i++{
		if i >= lenth {
			nums2[i] = num
			break
		}

		if num <= nums[i] {
			nums2[i] = num
			num = nums[i]
		}else {
			nums2[i] = nums[i]
		}

	}

	fmt.Println(nums2)
}
