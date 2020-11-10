package main

import (
	"fmt"
	"wechat"
)

func main (){
	wx,newErr :=wechat.NewWxTools("conf/secert.conf")
	if newErr!= nil {
		fmt.Printf("%+v",newErr)
	}
	token, err := wx.GetUserDetail("kyMu3uSG5nrReswyQphnkctHSN3pqWFtj7cS0aNnbFY")
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(token)

}
