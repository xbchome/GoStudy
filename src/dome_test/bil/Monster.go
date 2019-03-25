package bil

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Monster struct{
	Name string
	Age int
	Skill string
}


func (this *Monster)Store() bool {
	str,err := json.Marshal(this)
	if err != nil {
		fmt.Printf("err:",err)
		return false
	}

	file,err :=os.OpenFile("D:/Codes/code/go/src/dome_test/bil/monster.json",os.O_CREATE|os.O_RDONLY,0555)
	if err != nil {
		fmt.Println("打开文件失败：",err)
		return false
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(string(str))
	writer.Flush()


	return true

}

func (this *Monster) ReStore(src string)(bool, Monster) {
	monstring,err := os.Open(src)
	if err != nil {
		return false,Monster{}
	}

	defer monstring.Close()

	reader := bufio.NewReader(monstring)
	var strs string
	for   {
		stt,err := reader.ReadString('\n')
		strs += stt
		if err == io.EOF {
			break
		}
	}

	var mon Monster
	json.Unmarshal([]byte(strs),&mon)
	fmt.Println(mon)
	return true,mon

}