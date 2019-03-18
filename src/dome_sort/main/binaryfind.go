package main

import "fmt"

// 二分 查找 binaryfind
func main() {

	nums :=[11]int{1,2,3,4,5,6,7,8,9,20,22}

	var leftIndex int = 0
	var rightIndex int = len(nums)-1
	var cinndx = (leftIndex+rightIndex) / 2
	var num int
	fmt.Println("请输入要查找的数:")
	fmt.Scanln(&num)

	if nums[cinndx] > num {

		for i:=leftIndex; i<=cinndx-1 ;i++  {
			if nums[i] == num {
				fmt.Printf("找到 %d,下标: %d",nums[i],i)
				break
			}
		}
	}

	if nums[cinndx] == num{
		fmt.Printf("找到 %d,下标: %d",nums[cinndx],cinndx)
	}

	if nums[cinndx] < num {

		for i:=cinndx+1; i<=rightIndex;i++  {
			if nums[i] == num {
				fmt.Printf("找到 %d,下标: %d",nums[i],i)
				break
			}
		}
	}

	fmt.Println("======================")
	var nums1 = [6]int{1,3,5,7,18,19}
	binaryfind(&nums1,0,len(nums1)-1,7)
	fmt.Println("===========3=========")
	fmt.Println(binaryF([]int{1,3,4,12,14,15,34,35},15))


}

func binaryfind(nums *[6]int,leftIndex,rightIndex,search int) int  {

	if leftIndex > rightIndex{
		fmt.Println("无法找到")
		return -1
	}

	Middle := (leftIndex +rightIndex)/2

	if (*nums)[Middle] > search {

		binaryfind(nums,leftIndex,Middle-1,search)
	}else if (*nums)[Middle] < search  {

		binaryfind(nums,Middle+1,rightIndex,search)
	}else {
		fmt.Println("找到下标:",Middle)
	}
	return Middle
}

func binaryFind(nums []int,leftIndex,rightIndex,search int) int {
	if leftIndex > rightIndex{
		fmt.Println("无法找到")
		return -1
	}

	Middle := (leftIndex +rightIndex)/2

	if nums[Middle] > search {

		binaryFind(nums,leftIndex,Middle-1,search)
	}else if nums[Middle] < search  {

		binaryFind(nums,Middle+1,rightIndex,search)
	}else {
		fmt.Println("找到下标:",Middle)
	}
	return Middle
}

// 二分查找 非递归
func binaryF(nums []int,search int) int {
	leftIndex := 0
	rightIndex := len(nums)-1
	var mid int
	for  leftIndex <= rightIndex {
			mid = (leftIndex+rightIndex)/2
		if leftIndex > rightIndex{
			fmt.Println("无法找到")
			return -1
		}
		if nums[mid]==search {
			return mid
		}
			if nums[mid] > search {
				rightIndex = mid -1
			} else{
				leftIndex = mid+1
			}
	}

	return -1
}
