package main

import (
	"fmt"
	// "sumgpress.com/config"
	// "config/ini"
	// "config/json"
	// "wechat"

	"github.com/gnrliubin/config/ini"
	"github.com/gnrliubin/config/json"
)

func main() {
	// wx, newErr := wechat.NewWxTools("conf/wechat.json")
	// if newErr != nil {
	// fmt.Printf("tools error:%+v", newErr)
	// }
	// fmt.Println(wx)
	//
	type demo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	d := demo{}
	d.ID = 11
	d.Name = "github"

	j := jsc.NewJson()
	err := j.Save("conf/demo.json", &d)
	if err != nil {
		fmt.Println(err)
	}

	ifc.Ifctest()
}
