package main

import (
	"fmt"
)

func main() {
	var nums = [8]int {45,56,65,78,88,77,56,78}
	var min int = nums[0]
	var max int = nums[0]
	var sum int
	var average float64

	for i:=0;len(nums)>i;i++ {
		// 最大和最小值
		if max < nums[i] {
			max = nums[i]
		}

		if min > nums[i] {
			min = nums[i]
		}
		fmt.Println("=============")
		fmt.Println(min)

		sum += nums[i]
	}


	sum = sum - max -min

	average = float64(sum)/float64(6)

	fmt.Printf("min:%d max:%d sum:%d average:%0.2f",min,max,sum,average)

	//var Fraction [8]int
	//var abs float64
	//var min_v float64 = 0
	//var min_i int = 0
	//var max_v float64 =0
	//var max_i int =0
	//for i:=0;len(nums)>i;i++ {
	//	abs = math.Abs(average - float64(nums[i]))
	//	if i== 0 {
	//
	//	}
	//	if  abs < min_v{
	//		min_i = i // 绝对值相差最小
	//		min_v = abs
	//	}
	//
	//	if abs >
	//}

}
