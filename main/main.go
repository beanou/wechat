package main

import (
	"fmt"
	"wechat"
)

func main() {
	wx, newErr := wechat.NewWxTools("../conf/wechat.json")
	if newErr != nil {
		fmt.Printf("tools error:%+v", newErr)
	}
	fmt.Println(wx)

}
