package wechat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type JsonStruct struct{}

func NewJson() *JsonStruct {
	return &JsonStruct{}
}

func (this *JsonStruct) Load(confFile string, v interface{}) {
	data, err := ioutil.ReadFile(confFile)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &v)
	if err != nil {

	}
}

func (this *JsonStruct) Save(confFile string, v interface{}) {

}
