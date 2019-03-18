package main

import "fmt"

func main() {
	nums :=[15]int{10,45,2,3,4,6,5,7,8,9,12,34,1,11,1}
	nums2 :=[15]int{10,45,2,3,4,6,5,7,8,9,12,34,1,11,1}
	fmt.Println(Bubbling(nums[:]))
	fmt.Println("=========2======")
	fmt.Println(Bubbling2(nums2[:]))

}


func Bubbling(nums []int) ([]int) {
	var temp int
	var cont int
	for i:=0;(len(nums)-1)>i ;i++  {
		for j:=i+1;(len(nums))>j ;j++  {
			if nums[i] > nums[j] {
				temp = nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
			cont++
		}
	}
	fmt.Println("cont:",cont)
	return nums
}
//bubblesort 冒泡
func Bubbling2(nums []int) ([]int) {
	var temp int
	var cont int

	for i:=0; (len(nums)-1)>i; i++ {
		for j:=0;(len(nums)-1-i)>j;j++ {
			if nums[j] > nums[j+1] {
				temp = nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
			cont++
		}
	}

	fmt.Println("cont2：",cont)
	return nums
}

func bubbsort() {

}