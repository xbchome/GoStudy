package view

import (
	"fmt"
	"project/crm/bin/service"
)

type Menu struct {
	Look bool
	CustomerService *service.CustomerService
}

func (this *Menu)ShowMenu() {
	var key int
	for this.Look  {
		fmt.Println("-----------------START------------------")
		fmt.Println("               1增加客户")
		fmt.Println("               2修改客户")
		fmt.Println("               3删除客户")
		fmt.Println("               4客户列表")
		fmt.Println("               5退   出")
		fmt.Println("-----------------END-------------------")
		fmt.Print("   请选择(1-5)")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("-----------------添加开始------------------")
			this.CustomerService.Add()
			fmt.Println("-----------------添加完成------------------")
		case 2:
			fmt.Println(2)
		case 3:
			fmt.Println("请输入要删除id")
			var id uint
			fmt.Scanln(&id)
			this.CustomerService.Del(id)
		case 4:
			this.CustomerService.List()
		case 5:
			this.Exit()

		}

	}
}

func NewMenu() Menu {

	return Menu{true,service.NewCustomerService()}
}

func (this *Menu) Exit() {
	var key string
	for   {

		fmt.Print("是否退出(y/n):")
		fmt.Scanln(&key)
		if key =="y" {
			this.Look = false
			break
		}

		if key == "n" {
			break
		}
	}


}


