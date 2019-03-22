package service

import (
	"fmt"
	"project/crm/bin/model"
)
//
type CustomerService struct {
	customers []model.Customer
	customerId uint
}

func NewCustomerService()*CustomerService {
	var customers []model.Customer
	customers=append(customers,model.NewCustomer(1,12,"xbc","男","444@qq.com","13724177651"))
	return &CustomerService{customerId:1,customers:customers}
}

func (this *CustomerService) List() {

	fmt.Printf("用户编号 \t 姓名 \t 性别 \t 年龄 \t 电话 \t\t 邮件 \n ")
	for _,v := range this.customers {
		fmt.Printf(v.ShowInfo())
	}
}

func (this *CustomerService) Add() {
		var name string
		var gender string
		var email string
		var phone string
		var age int
	fmt.Print("请输入用户名:")
	fmt.Scanln(&name)
	fmt.Print("性别：")
	fmt.Scanln(&gender)
	fmt.Print("邮件：")
	fmt.Scanln(&email)
	fmt.Print("手机：")
	fmt.Scanln(&phone)
	fmt.Print("年龄：")
	fmt.Scanln(&age)
		this.customerId ++
		this.customers = append(this.customers,model.NewCustomer(this.customerId,age,name,gender,email,phone))

}

func (this *CustomerService) Del(id uint) {
	index:=-1
	for  k,v:=range this.customers{
		if v.Id == id {
			index = k
			break
		}
	}

	if index !=-1{
		this.customers = append(this.customers[:index],this.customers[index+1:]...)
		fmt.Println("删除成功！")
		return
	}
	fmt.Println("没有找到")
}

func (this *CustomerService) findId() {

}

