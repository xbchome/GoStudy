package main

import "fmt"

func main() {
	fmt.Println(fbn(100))
	fmt.Println(fbn2(100))
}

func fbn(n int) ([]uint64) {
	var fb []uint64
	for i:=0;i<=n ;i++  {
		if i <= 1 {
			fb = append(fb,1)
		}else {
			fb = append(fb,fb[i-1]+fb[i-2])
		}
	}

	return fb
}

func fbn2(n int) ([]uint64) {
	var fb = make([]uint64,n)

	for i,_:=range fb {

		if i<= 1 {
			fb[i] = 1
		} else{
			fb[i] = fb[i-1] + fb[i-2]
		}
	}

	return fb
}
