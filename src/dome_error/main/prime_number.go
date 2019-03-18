package main

import "fmt"

func main() {
	prime(100)
}

func prime(num int) {
    var sum int
	var count int = 0
	var ifs int = 0
	for i:=2;i<=num;i++{
		ifs = 0
		for j:=2;j<=i;j++{
			if i%j==0 && j!=i {
				ifs++
			}
		}

		if ifs == 0 {
				sum +=i
				fmt.Printf("%d ",i)
				count++
		}

		if count == 5 {
			count =0
			fmt.Printf("\n")
		}
	}

	fmt.Printf("素数和:%d",sum)
}

