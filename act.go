/*
	请求的具体实现
*/

package wechat

import (
	"encoding/json"
	"fmt"
	"time"

	"gopkg.in/ini.v1"

	"github.com/pkg/errors"
)

type WxTools struct {
	configFile string //配置文件
	tokenFile  string //token缓存文件
	appid      string
	state      string
	secret     string
	token      *AccessToken
}

//wxapi工具类初始化方法
//
//工厂模式创建类，从配置中读取企业微信信息
func NewWxTools(configFile string) (*WxTools, error) {

	// 生产微信工具类实力
	tools := new(WxTools)
	// 写入配置文件路径
	tools.configFile = configFile
	//读取配置文件
	conf, err := ini.Load(tools.configFile)
	if err != nil {
		return nil, errors.Wrap(err, "get config falid!!!")
	}

	// 将配置信息写入工具类属性中
	tools.appid = conf.Section("wechat").Key("Appid").String()
	tools.state = conf.Section("wechat").Key("State").String()
	tools.secret = conf.Section("wechat").Key("Secret").String()
	tools.tokenFile = conf.Section("wechat").Key("TokenFile").String()
	tools.token = new(AccessToken)
	//返货类指针
	return tools, nil
}

// 获取access_token
/**
 	查看cacheFile中的accesstoken是否过期，如果没过期则使用缓存中的数据
	如果过期则重新请求，并存入缓存文件中
*/
func (this *WxTools) GetToken() (interface{}, error) {
	//读取配置文件
	conf, err := ini.Load(this.tokenFile)
	if err != nil {
		return nil, errors.Wrap(err, "can't find token file")
	}

	//配置文件中的信息存入类属性
	this.token.Token = conf.Section("").Key("access_token").MustString("")
	this.token.Expires = conf.Section("").Key("expires").MustInt64(0)

	//token过期时重新获取token
	if this.token.Expires == 0 || this.token.Expires <= time.Now().Unix() {

		tokenByte, reqSendErr := ReqSend("GET", GetAccessTokenUri(this.appid, this.secret), nil)
		if reqSendErr != nil {
			//return nil,fmt.Errorf("request wx api error while getting accessToken %q",reqSendErr)
			return nil, errors.Wrap(reqSendErr, "request wx api error while getting accessToken")
		}

		json.Unmarshal(tokenByte, &this.token)

		if this.token.Errcode > 0 {
			//请求错误处理
			return nil, errors.Wrap(fmt.Errorf("%d--%v", this.token.Errcode, this.token.Errmsg), "fail to get accesstoken from wxapi")
		}

		//token写入缓存文件
		this.token.Expires = this.token.Expires + time.Now().Unix()
		// jsonconf.Save(this.tokenFile, this.token)
		conf.Section("").Key("access_token").SetValue(this.token.Token)
		conf.Section("").Key("expires").SetValue(fmt.Sprintf("%d", this.token.Expires))
		conf.SaveTo(this.tokenFile)
	}

	return this.token.Token, nil
}

//
// //使用code获取用户信息
//
func (this *WxTools) GetUserInfo(code string) (map[string]interface{}, error) {
	//获取accesstoken
	_, getTokenErr := this.GetToken()
	if getTokenErr != nil {
		return nil, errors.Wrap(getTokenErr, "fail to get  accesstoken\n")
	}
	//使用accesstoken和code请求接口
	userInfoByte, reqSendErr := ReqSend("GET", GetUserInfoUri(this.token.Token, code), nil)
	if reqSendErr != nil {
		return nil, errors.Wrap(reqSendErr, "request wx api error while getting userInfo")
	}

	userInfo := make(map[string]interface{})

	json.Unmarshal(userInfoByte, &userInfo)

	{
		//DP(userInfo)
	}
	//如果没有得到userid则返回错误（也可以使用userInfo["errcode"] !=0判断）
	//if(userInfo["errcode"] !=0){
	if userInfo["UserId"] == nil {
		return nil, errors.New(fmt.Sprintf("fail to get userid or openid  |  %v", userInfo))
	}

	return userInfo, nil
}

//获取用户详细信息
func (this *WxTools) GetUserDetail(code string) (map[string]interface{}, error) {

	userDetail := make(map[string]interface{})
	//userInfo := new(UserInfo)
	userInfo := make(map[string]interface{})
	userInfo, userInfoErr := this.GetUserInfo(code)
	if userInfoErr != nil {
		return nil, errors.Wrap(userInfoErr, "fail to get userInfo before getting userDetail\n")
	}
	userId := userInfo["UserId"]

	//在获取userinfo时已经有accesstoken被存入类属性中，这段获取是防止accesstoken恰巧过期的极端情况
	_, getTokenErr := this.GetToken()
	if getTokenErr != nil {
		return nil, errors.Wrap(getTokenErr, "fail to get  accesstoken while getting userDetail")
	}

	userDetailByte, reqSendErr := ReqSend("GET", GetUserDetailUri(this.token.Token, userId), nil)
	if reqSendErr != nil {
		return nil, errors.Wrap(reqSendErr, "request wx api error while getting userDetail")
	}
	json.Unmarshal(userDetailByte, &userDetail)
	{
		//DP(userDetail)
	}
	if userDetail["userid"] == nil {
		return nil, errors.New(fmt.Sprintf("fail to get user detail  |  %v", userDetail))
	}

	return userDetail, nil

}
