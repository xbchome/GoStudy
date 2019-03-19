package main

import (
	"encoding/json"
	"fmt"
)

type Dog struct {
	Name string `json:"name"`
	Age int	`json:"age"`
}

func main() {
	var dog Dog = Dog{"杨康",1}
	jsonStr,_ := json.Marshal(dog)

	dog2 := map[string]string{
		"name":"xbc",
	}
	dog2Json,_ := json.Marshal(dog2)
	fmt.Println(string(dog2Json))
	fmt.Println(string(jsonStr))

}
