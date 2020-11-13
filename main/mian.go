package main

import (
	"fmt"
	"wechat"
)

func main (){
	wx,newErr :=wechat.NewWxTools("conf/wechat.conf")
	if newErr!= nil {
		fmt.Printf("tools error:%+v",newErr)
	}
	token, err := wx.GetUserDetail("kyMu3uSG5nrReswyQphnkctHSN3pqWFtj7cS0aNnbFY")
	if err !=nil {
		fmt.Printf("%+v",err)
	}
	fmt.Println("token:",token)

}
