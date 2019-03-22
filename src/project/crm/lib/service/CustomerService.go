package service

import (
	"project/crm/lib/model"
)

type CustomerService struct {
	customer []model.Customer
	Id int
}

// 工厂方法
func New() *CustomerService {
	var customer []model.Customer
	customer = append(customer,model.NewCustomer(1,"管理员",18,"男","333",123))
	return &CustomerService{
		Id:1,
		customer:customer,
	}
}

func (this *CustomerService) List() string  {
	 var info string = "编号 \t 客户名 \t 性别 \t 年龄 \t 邮箱 \t 电话\n"
	for _,v:=range this.customer {
		info += v.GetInf()
	}
	return info
}

func (this *CustomerService) Add(name string,age int,gender,email string,phone int)  {
	this.Id ++
	this.customer = append(this.customer,model.NewCustomer(this.Id,name,age,gender,email,phone))

}

func(this *CustomerService) FindIndex(id int) int {
	index := -1
	for k,v := range this.customer {
		if v.Id == id {
			index = k
			break
		}
	}

	return index
}

func (this *CustomerService) Edit(index int,name string,age int,gender,email string,phone int)bool {
		this.customer[index] = model.NewCustomer(this.customer[index].Id,name,age,gender,email,phone)
		return true
}

func (this *CustomerService) Del(index int) {
	this.customer = append(this.customer[:index],this.customer[index+1:]...)
}