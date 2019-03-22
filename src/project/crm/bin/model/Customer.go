package model

import "fmt"

type Customer struct {
	Id uint
	Name string
	Gender string
	Age int
	Email string
	Phone string
}

func NewCustomer(id uint,age int, name ,gender, email, phone  string ) Customer {
	return Customer{
		Id:id,
		Name:name,
		Age:age,
		Gender:gender,
		Email:email,
		Phone:phone,
	}
}

func (this Customer) ShowInfo() string {
	return fmt.Sprintf("%v \t\t %v \t  %v \t %v \t %v \t %v \n ",this.Id,this.Name,this.Gender,this.Age,this.Phone,this.Email)
}
