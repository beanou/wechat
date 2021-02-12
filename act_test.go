package wechat

import (
	"fmt"
	"testing"
)

func TestAct(t *testing.T) {

	wxTools, err := NewWxTools("./conf/app.ini")
	if err != nil {
		fmt.Println(err)
	}

	wxTools.GetToken()

	fmt.Println(wxTools)
}
