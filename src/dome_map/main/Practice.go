package main

import "fmt"

func main() {
	var user map[string] map[string]string
	user = make(map[string]map[string]string,1)
	user["xbc"] = map[string]string{
		"nickName":"xbc001",
		"password":"7777",
	}
	fmt.Println(user)
	modifUser(user,"xbc")
	fmt.Println(user)

}

func modifUser(user map[string] map[string]string,name string) {
	if user[name] == nil {
		user[name] =map[string]string{
			"nickName":"xbc001",
			"password":"888888",
		}
	}else{
		user[name]["password"] = "88888" 
	}
}
