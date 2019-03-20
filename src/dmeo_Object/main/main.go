package main

import (
	"dmeo_Object/bin"
	"fmt"
)

func main() {
	var p = bin.NewPerson("xbc")
	p.Setsex("男")
	p.SetAge(13)
	fmt.Println(p)

	fmt.Println("===============TEXT=============")

	var acc bin.Account

	(&acc).SetAccountId("gq12344")
	acc.SetBalance(20.50)
	acc.SetPwd("123456")

	fmt.Println(acc)
}