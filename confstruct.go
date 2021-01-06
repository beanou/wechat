/*
	微信功能配置信息所需要的结构体
*/
package wechat

// 企业微信ID和应用安全信息结构体
type WeConf struct {
	Appid     string `json:"appid"`  // 微信ID
	State     string `json:"state"`  // code跳转自定义码
	Secret    string `json:"secret"` // 应用安全码
	CacheFile string `json:"cache"`  // token信息存放文件
}
