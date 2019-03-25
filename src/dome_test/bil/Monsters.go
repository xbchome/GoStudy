package bil

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

type Monsters struct {
	Name string
	Age int
	Gender string
}

func (this *Monsters) SaveJson() error {
	str,err := json.Marshal(this)
	if err != nil {
		return err
	}
	file,err := os.OpenFile("D:/moster.json",os.O_CREATE|os.O_RDWR,0777)
	if err != nil {
		return  err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(string(str))
	writer.Flush()
	return nil
}

func (this *Monsters) GetObj(src string)Monsters {
	file ,err := os.Open(src)
	if err!= nil {
		return Monsters{}
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var jsonStr string
	for   {
		str,err :=reader.ReadString('\n')
		jsonStr += str
		if err == io.EOF {
			break
		}
	}
	var monsters Monsters
	json.Unmarshal([]byte(jsonStr),&monsters)
	return monsters
}

