package main

import "fmt"

func main() {
	var arrs = [4]int{1,2,3,4}

	slice := arrs[2:4] // 直接引用数组

	fmt.Println("slice=",slice)
	fmt.Println("len=",len(slice))
	fmt.Println("cap=",cap(slice))
	fmt.Println(arrs)
	slice[0] = 1
	fmt.Println(slice)
	fmt.Println(arrs)

	fmt.Println("====================方式二make===========")

	slice2 := make([]int,2,4) // make 创建了一个slice len 2 cap 4
	fmt.Println("slice=",slice2)
	fmt.Println("len=",len(slice2))
	fmt.Println("cap=",cap(slice2))


	fmt.Println("====================方式3===========")

	slice3 := []int{1,23,4}
	fmt.Println("slice=",slice3)
	fmt.Println("len=",len(slice3))
	fmt.Println("cap=",cap(slice3))
	fmt.Println("arrscap=",cap(arrs))

	fmt.Println("====================copy===========")
	sliceC1 := []int{1,2,34,5,55}

	sliceC2 :=make([]int,2,3)

	copy(sliceC2,sliceC1)

	fmt.Println("c2",sliceC2)

	fmt.Println("====================string===========")
	var strt string = "你好啊啊"

	var sliceS1  = []rune(strt)
	fmt.Println("sliceS1 len",len(sliceS1))

	fmt.Println("====================string===========")


}
