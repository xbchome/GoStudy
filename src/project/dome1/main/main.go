package main

import (
	"fmt"
	"project/dome1/bin"
)

type FamilyIncome struct {
	Button int // 按钮号
	look bool // 开关
	money float64
	balance float64 // 余额
	details string // 描述
	note string //记录
	orNot string // y n 是否要退出
	flag bool // 是否有收支记录

}

func (this *FamilyIncome) showMenu() {
	for this.look  {
		fmt.Println("================TEXT===============")
		fmt.Println("                1 收支明细")
		fmt.Println("                2 登记收入")
		fmt.Println("                3 登记支出")
		fmt.Println("                4 退   出")
		fmt.Print("请选择(1-4)")
		fmt.Scanln(&this.Button)
		switch this.Button {
		case 1:
			this.showInfo()
		case 2:
			this.add()
		case 3:
			this.reduce()
		case 4:
			this.exit()
		default:
			fmt.Println("输入有误!")


		}

	}
}

func (this *FamilyIncome) showInfo () {
	fmt.Println("================收支明细===============")
	if this.flag {
		fmt.Printf(this.details)
	}else {
		fmt.Println("目前还没有收支记录")
	}

}

func (this *FamilyIncome) add() {
	fmt.Print("请输入收入金额:")
	fmt.Scanln(&this.money)
	if this.money >0{
		this.balance += this.money
		fmt.Print("请输入金额来源：")

		this.flag = true
		fmt.Scanln(&this.note)
		this.details += fmt.Sprintf("收入\t %v\t %v\t %v\n",this.balance,this.money,this.note)
	}else {
		fmt.Println("收入不能小于1块钱")
	}

}

func (this *FamilyIncome) reduce() {
	fmt.Print("请输入支取金额:")
	fmt.Scanln(&this.money)
	if this.money> this.balance{
		fmt.Println("余额不足")
		return
	}

	this.flag = true
	this.balance -= this.money
	fmt.Print("请输入金额用途：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("支出\t %v\t %v\t %v\n",this.balance,this.money,this.note)
}

func (this *FamilyIncome) exit() {
	for   {
		fmt.Println("你是否要退出(y/n):")
		fmt.Scanln(&this.orNot)
		if this.orNot=="y"{
			this.look=false
			break
		}else if this.orNot == "n"{
			break
		}
	}
}

func main() {


	var fam = bin.NewFamilyIncome()
	fam.ShowMenu()

}
