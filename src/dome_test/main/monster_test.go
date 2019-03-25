package main

import (
	"dome_test/bil"
	"testing"
)
/*
 测试 json写入和读取是否有一致
 */
func TestGetObj (t *testing.T){
	var mon bil.Monsters
	//var mon2 bil.Monsters
	mon = bil.Monsters{
		Name:"xbc",
		Age:1,
		Gender:"nan",
	}
	mon.SaveJson()
	mon2:=mon.GetObj("D:/moster.json")
	if mon2.Name == mon.Name {
		t.Logf("一致")
	} else{
		t.Fatalf("期望的：%v; 实际的:%v",mon.Name,mon2.Name)
	}

}
