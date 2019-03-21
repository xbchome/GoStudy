package main

import "fmt"
type Usb interface {
	Start()
	Stop()
	//X()
}


type Pc struct {

}

func (pc Pc) Insert(usb Usb) {
	usb.Start()
	usb.Stop()
}

type xj struct {
}

func (X xj) Start() {
	fmt.Println("相机开始工作")
}

func (X xj) Stop() {
	fmt.Println("相机完成工作")
}


type sj struct {

}

func (X sj) Start() {
	fmt.Println("手机开始工作")
}

func (X sj) Stop() {
	fmt.Println("手机完成工作")
}

func main() {
	fmt.Println("=================接口===========")

	pc :=Pc{}
	sj:=sj{}
	xj:=xj{}

	pc.Insert(sj)
	pc.Insert(xj)


}

