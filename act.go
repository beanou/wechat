/*
	请求的具体实现
 */

package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/adapter/config"
	"github.com/pkg/errors"
	"time"
)

type WxTools struct {
	configFile  string		//配置文件
	cacheFile 	string		//token缓存文件
	appid		string
	state 		string
	secret 		string
	token 		*AccessToken
}


//wxapi工具类初始化方法
//
//工厂模式创建类，从配置中读取企业微信信息
func NewWxTools(configFile string) (*WxTools,error){

	//初始化类及其属性
	// BUG(liubin): a bug in code
	tools := new(WxTools)
	tools.configFile = configFile
	//读取配置
	configer, err := config.NewConfig("ini", tools.configFile)
	if (err!=nil) {
		return nil ,errors.Wrap(err,"error while reade config file")
	}
	tools.cacheFile = configer.String("wx::cache")
	tools.appid = configer.String("wx::appid")
	//tools.redirectUri = configer.String("wx::redirect_uri")
	tools.state  = configer.String("wx::state")
	tools.secret = configer.String("wx::secret")
	tools.token=new(AccessToken)

	//返货类指针
	return tools,nil
}

// 获取access_token
/**
 	查看cacheFile中的accesstoken是否过期，如果没过期则使用缓存中的数据
	如果过期则重新请求，并存入缓存文件中
*/
func (this *WxTools) GetToken() (interface{},error)  {
	//读取配置文件
	configer,err:= config.NewConfig("ini", this.cacheFile)
	if (err!=nil) {
		return nil,errors.Wrap(err,"read config file error")
	}

	//配置文件中的信息存入类属性
	this.token.Token = configer.String("token::access_token")
	this.token.Expires, _ = configer.Int64("token::expires")

	//token过期时重新获取token
	if this.token.Expires == 0 || this.token.Expires <= time.Now().Unix() {

		tokenByte,reqSendErr := ReqSend("GET",GetAccessTokenUri(this.appid,this.secret),nil)
		if reqSendErr != nil {
			//return nil,fmt.Errorf("request wx api error while getting accessToken %q",reqSendErr)
			return nil ,errors.Wrap(reqSendErr,"request wx api error while getting accessToken")
		}
		json.Unmarshal(tokenByte, &this.token)

		if this.token.Errcode > 0 {
			//请求错误处理
			return nil ,errors.New(fmt.Sprintf("get access_token error! %v\n",this.token))
		}

		//token写入缓存文件
		this.token.Expires = this.token.Expires + time.Now().Unix()
		configer.Set("token::access_token", this.token.Token)
		configer.Set("token::expires", fmt.Sprintf("%d", this.token.Expires))
		configer.SaveConfigFile("conf/token.ini")
	}

	return this.token.Token,nil
}


//使用code获取用户信息

func (this *WxTools) GetUserInfo(code string) (map[string]interface{},error) {
	//获取accesstoken
	_,getTokenErr := this.GetToken()
	if getTokenErr!=nil{
		return nil,errors.Wrap(getTokenErr,"fail to get  accesstoken\n")
	}
	//使用accesstoken和code请求接口
	userInfoByte,reqSendErr := ReqSend("GET", GetUserInfoUri(this.token.Token,code),nil)
	if reqSendErr != nil {
		return nil,errors.Wrap(reqSendErr,"request wx api error while getting userInfo")
	}

	userInfo := make(map[string]interface{})

	json.Unmarshal(userInfoByte,&userInfo)

	{
		//DP(userInfo)
	}
	//如果没有得到userid则返回错误（也可以使用userInfo["errcode"] !=0判断）
	//if(userInfo["errcode"] !=0){
	if (userInfo["UserId"] == nil){
		return nil,errors.New(fmt.Sprintf("fail to get userid or openid  |  %v",userInfo))
	}

	return userInfo,nil
}


//获取用户详细信息
func (this *WxTools) GetUserDetail (code string) (map[string]interface{},error) {

	userDetail :=make(map[string]interface{})
	//userInfo := new(UserInfo)
	userInfo := make(map[string]interface{})
	userInfo,userInfoErr := this.GetUserInfo(code)
	if userInfoErr!=nil{
		return nil,errors.Wrap(userInfoErr,"fail to get userInfo before getting userDetail\n")
	}
	userId := userInfo["UserId"]

	//在获取userinfo时已经有accesstoken被存入类属性中，这段获取是防止accesstoken恰巧过期的极端情况
	_,getTokenErr := this.GetToken()
	if getTokenErr!=nil{
		return nil,errors.Wrap(getTokenErr,"fail to get  accesstoken while getting userDetail")
	}

	userDetailByte, reqSendErr:= ReqSend("GET",GetUserDetailUri(this.token.Token,userId),nil)
	if reqSendErr != nil {
		return nil,errors.Wrap(reqSendErr,"request wx api error while getting userDetail")
	}
	json.Unmarshal(userDetailByte,&userDetail)
	{
		//DP(userDetail)
	}
	if (userDetail["userid"] == nil){
		return nil,errors.New(fmt.Sprintf("fail to get user detail  |  %v",userDetail))
	}

	return userDetail,nil

}