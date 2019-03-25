package bil

import "fmt"

func GetSun(n1,n2 int) int {
	return  n1+n2
}

func GetS(n1,n2 int) int{
	//var num int
	sun :=[]int{2}
	for i:=2;10000 >i ;i++  {
		if i== 2 {
			continue
		}
		for j:=2;j<=i-1 ;j++  {
			if i%j == 0 {
				break
			}else if j==i-1 {
				sun = append(sun,i)
			}
		}


	}
	fmt.Println(sun)
	return n1-n2
}