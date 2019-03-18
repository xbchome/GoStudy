package main

import "fmt"

func main() {

	multiples9(20)
	six6()


}

// 查找9的倍数
func multiples () {
	var sum int = 0
	var num int = 0
	for i := 9; i<=100; i++{
		s := i%9
		if s == 0 {
			sum += i
			num ++
		}
	}

	fmt.Printf("sum:%d num:%d \n",sum,num)
}

func multiples9(max int) {

	var sun int = 0
	var num int = 0
	for i := 9;max>= i; i++{

		if i%9 == 0 {
			sun += i
			num ++
		}
	}

	fmt.Printf("sun : %v; num: %v \n",sun,num)

}

//func six() {
//	var k int = 6
//	for i:=0;i<=6;i++ {
//		fmt.Printf("%d + %d = %d",i,k,i+k)
//		k--
//	}
//}

func six6() {
	var n int = 10

	for i:=0; n>=i; i++{
		fmt.Printf("%v + %v = %v \n",i,n-i,i+(n-i))
	}
}



