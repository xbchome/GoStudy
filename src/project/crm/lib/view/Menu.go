package view

import (
	"fmt"
	//"project/crm/lib/model"
	"project/crm/lib/service"
)

type Menu struct {
	look bool
	key int
	customerService *service.CustomerService
}

func NewMenu() *Menu {
	customerService :=service.New()
	return &Menu{
		look:true,
		customerService:customerService,
	}
}

func (this *Menu) Show() {

	for this.look  {
		this.showOption()
		switch this.key {
		case 1:
			this.showList() // 显示用户列表
		case 2:
			this.add()
		case 3:
			this.edit()
		case 4:
			this.del()
		case 5:
			this.exit()
		default:
			fmt.Println("暂时没有此操作")

		}

	}
}

func (this *Menu)showOption() {
	fmt.Println("---------------MENU----------------")
	fmt.Println("               1 显示客户列表")
	fmt.Println("               2 增       加")
	fmt.Println("               3 修       改")
	fmt.Println("               4 删       除")
	fmt.Println("               5 退       出")
	fmt.Println("---------------MENU----------------")
	fmt.Print("请输入(1-5): ")
	fmt.Scanln(&this.key)
}

func (this *Menu) showList() {
	fmt.Println("---------------客户列表----------------")
	fmt.Printf(this.customerService.List())
	fmt.Printf("-----------------end------------------\n\n")
}

func (this *Menu) add() {
	fmt.Println("---------------增加----------------")
	fmt.Print("请输入姓名:")
	var name string
	fmt.Scanln(&name)
	fmt.Print("请输入性别:")
	var gender string
	fmt.Scanln(&gender)
	fmt.Print("请输入年龄:")
	var age int
	fmt.Scanln(&age)
	fmt.Print("请输入邮箱:")
	var email string
	fmt.Scanln(&email)
	fmt.Print("请输入电话:")
	var phone int
	fmt.Scanln(&phone)
	this.customerService.Add(name,age,gender,email,phone)
	fmt.Printf("---------------增加完成--------------\n\n")

}

func (this *Menu) edit() {
	fmt.Println("---------------修改----------------")
	var index int
	for   {
		fmt.Println("请输入要修改的编号(-1退出):")
		var id int
		fmt.Scanln(&id)
		if id==-1 {
			return
		}
		index = this.customerService.FindIndex(id)
		if  index == -1 {
			fmt.Println("客户不存在")
		}else {
			break
		}
	}

	fmt.Print("请输入姓名:")
	var name string
	fmt.Scanln(&name)
	fmt.Print("请输入性别:")
	var gender string
	fmt.Scanln(&gender)
	fmt.Print("请输入年龄:")
	var age int
	fmt.Scanln(&age)
	fmt.Print("请输入邮箱:")
	var email string
	fmt.Scanln(&email)
	fmt.Print("请输入电话:")
	var phone int
	fmt.Scanln(&phone)
	this.customerService.Edit(index,name,age,gender,email,phone)
	fmt.Printf("---------------增加完成--------------\n\n")
}

func (this *Menu) del() {
	fmt.Println("---------------删除----------------")
	fmt.Print("请输入要删除的编号(-1退出):")
	var id int
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	var index int
	index = this.customerService.FindIndex(id)
	if index==-1 {
		fmt.Print("没有找到\n\n")
		return
	}

	this.customerService.Del(index)
	fmt.Printf("-------------删除成功--------------\n\n")


}

func (this *Menu) exit() {
	var key string
	for   {

		fmt.Print("是否退出(y/n):")
		fmt.Scanln(&key)
		if key =="y" {
			this.look = false
			break
		}

		if key == "n" {
			break
		}
	}
}
