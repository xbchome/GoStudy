package main

import (
	"dome_test/bil"
	"testing"
)

func TestGetSun(t *testing.T) {
	sun := bil.GetSun(3,5)
	if sun == 8 {
		t.Logf("测试正确")
	} else{
		t.Fatalf("error:期待的结果%v; 实际结果:%v",8,sun)
	}
}

func TestGetS(b *testing.T) {

	num := bil.GetS(5,4)
	if num != 9 {
		b.Fatalf("期待的结果%v,实际的结果%v",9,num)
	}

	b.Logf("成功")
}
