/*
	发送请求功能
 */
package wechat

import (
	"io"
	"io/ioutil"
	"net/http"
)

//发送请求功能
func ReqSend (method string , uri string,body io.Reader)  ([]byte,error){

	client :=&http.Client{}
	request, reqError := http.NewRequest(method, uri, body)
	if reqError != nil {
		return nil,reqError
	}
	response ,respError :=client.Do(request)
	if respError != nil {
		return nil,respError
	}
	responseByte,respByteError:=ioutil.ReadAll(response.Body)
	if respByteError != nil {
		return nil,respByteError
	}

	return responseByte,nil


}
