/*
  返回企业微信api请求地址
	获取accesstoken的URL
	func GetAccessTokenUri(cropid string , secret string) string {}

	根据code获取成员信息的URL
	func GetUserInfoUri(accessToken string, code string) string {}

	根据userid获取通讯录成员详细信息
	func GetUserDetailUri (accessToken string ,userId string) string {}
 */

package wechat

import "fmt"

//获取accesstoken的URL

//返回结构
//{
//"errcode": 0,
//"errmsg": "ok",
//"access_token": "accesstoken000001",
//"expires_in": 7200
//}
func GetAccessTokenUri(cropid string , secret string) string {
	return fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s",cropid,secret)
}

//根据code获取成员信息的URL

//返回结构
//企业内部成员
//{
//"errcode": 0,
//"errmsg": "ok",
//"UserId":"USERID",
//"DeviceId":"DEVICEID"
//}
//非企业内部成员
//{
//"errcode": 0,
//"errmsg": "ok",
//"OpenId":"OPENID",
//"DeviceId":"DEVICEID"
//}
//错误信息
//{
//"errcode": 40029,
//"errmsg": "invalid code"
//}
func GetUserInfoUri(accessToken string, code string) string {
	return fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s",accessToken,code)
}

//通过userid获取成员详细信息

//返回结构
//{
//	"errcode": 0,
//	"errmsg": "ok",
//	"userid": "zhangsan",
//	"name": "李四",
//	"department": [1, 2],
//	"order": [1, 2],
//	"position": "后台工程师",
//	"mobile": "13800000000",
//	"gender": "1",
//	"email": "zhangsan@gzdev.com",
//	"is_leader_in_dept": [1, 0],
//	"avatar": "http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/0",
//	"thumb_avatar": "http://wx.qlogo.cn/mmopen/ajNVdqHZLLA3WJ6DSZUfiakYe37PKnQhBIeOQBO4czqrnZDS79FH5Wm5m4X69TBicnHFlhiafvDwklOpZeXYQQ2icg/100",
//	"telephone": "020-123456",
//	"alias": "jackzhang",
//	"address": "广州市海珠区新港中路",
//	"open_userid": "xxxxxx",
//	"main_department": 1,
//	"extattr": {
//		"attrs": [
//			{
//				"type": 0,
//				"name": "文本名称",
//				"text": {
//					"value": "文本"
//				}
//			},
//			{
//				"type": 1,
//				"name": "网页名称",
//				"web": {
//				"url": "http://www.test.com",
//					"title": "标题"
//				}
//			}
//		]
//	},
//	"status": 1,
//	"qr_code": "https://open.work.weixin.qq.com/wwopen/userQRCode?vcode=xxx",
//	"external_position": "产品经理",
//	"external_profile": {
//		"external_corp_name": "企业简称",
//		"external_attr": [
//			{
//				"type": 0,
//				"name": "文本名称",
//				"text": {
//					"value": "文本"
//				}
//			},
//			{
//				"type": 1,
//				"name": "网页名称",
//				"web": {
//					"url": "http://www.test.com",
//					"title": "标题"
//				}
//			},
//			{
//				"type": 2,
//				"name": "测试app",
//				"miniprogram": {
//					"appid": "wx8bd80126147dFAKE",
//					"pagepath": "/index",
//					"title": "my miniprogram"
//				}
//			}
//		]
//	}
//}
func GetUserDetailUri (accessToken string ,userId interface{}) string {
	return fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=%s&userid=%s",accessToken,userId)
}


