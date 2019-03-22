package main

import "fmt"

type Usbs interface {
	start()
	stop()
}

type A struct {

}

type B struct {

}

func (a A) start() {
	fmt.Println("A start")
}

func (a A) stop() {
	fmt.Println("A stop")
}

func (b B) start() {
	fmt.Println("B start")
}

func (b B) stop() {
	fmt.Println("b stop")
}

func (b B) say() {
	fmt.Println("b====>")
}

func isType(nums ...interface{}) {
 for i,v := range nums {
	 switch v.(type) {
	 case int,int32:
		fmt.Printf("i:%d; v:%v type:int\n",i,v)
	 case bool:
		 fmt.Printf("i:%d; v:%v type:bool\n",i,v)
	 case string:
		 fmt.Printf("i:%d; v:%v type:string\n",i,v)
	 case float32,float64:
		 fmt.Printf("i:%d; v:%v type:float\n",i,v)
	 case nil:
		 fmt.Printf("i:%d; v:%v type:nil\n",i,v)
	 case B:
		 fmt.Printf("i:%d; v:%v type:b\n",i,v)
	 case *B:
		 fmt.Printf("i:%d; v:%v type:*b\n",i,v)


	 default:
		 fmt.Printf("i:%d; v:%v type:default\n",i,v)

	 }
 }
}
func main() {
	var usbs []Usbs

	b := B{}
	a :=A{}

	usbs = append(usbs,a)
	usbs = append(usbs,b)

	for _,v:= range usbs {
		if bv,ok:= v.(B);ok {
			bv.say()
			fmt.Println("b")
		}

		v.start()
		v.stop()
	}

	isType(1,'2',"12",[2]int{1,2},12.3,b,&b)
}
