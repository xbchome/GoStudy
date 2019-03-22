package model

import "fmt"

type Customer struct {
	Id int
	Name string
	Age int
	Gender string
	Email string
	Phone int
}


func  NewCustomer(id int,name string,age int,gender,email string,phone int) Customer {
	return Customer{
		Id:id,
		Name:name,
		Age:age,
		Gender:gender,
		Email:email,
		Phone:phone,
	}
}

func (this Customer) GetInf() string {
	return fmt.Sprintf("%v \t %v \t %v \t %v \t %v \t %v\n",this.Id,this.Name,this.Gender,this.Age,this.Email,this.Phone)
}
