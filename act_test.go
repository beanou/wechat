package wechat

import (
	"fmt"
	"testing"
)

func TestAct(t *testing.T) {

	wxTools, err := NewWxTools("conf/app.ini", "1000009")
	if err != nil {
		fmt.Println(err)
	}

	wxTools.GetToken()
	// _, err = wxTools.GetUserInfo("ksdljfsdf8sdfsdf09<F9>990sdf")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wxTools)
}
