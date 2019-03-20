package bin

import "fmt"

type Account struct {
	accountId string
	balance float64
	pwd string
}

func (ac *Account) SetAccountId(id string) {
	ids := []rune(id)
	id_len := len(ids)
	if  id_len >= 6 && 10>= id_len {
		fmt.Println("成功！")
		ac.accountId = id
	} else{
		fmt.Println("账号信息不符合要求")
		return
	}
}

func (ac *Account) SetBalance(Balance float64) {
	if Balance<20 {
		fmt.Println("存入金额太少")
		return
	}

	ac.balance = Balance
}

func (ac *Account) SetPwd(pwd string) {
	if len(pwd) != 6 {
		fmt.Println("密码必须为6位")
		return
	}
	ac.pwd = pwd
}

