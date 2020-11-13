/*
	发送请求功能
 */
package wechat

import (
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
)

//发送请求功能
func ReqSend (method string , uri string,body io.Reader)  ([]byte,error){

	client :=&http.Client{}
	request, reqError := http.NewRequest(method, uri, body)
	if reqError != nil {
		return nil,errors.Wrap(reqError,"error to create a http request")
	}
	response ,respError :=client.Do(request)
	if respError != nil {
		return nil,errors.Wrap(respError,"send request and get response error")
	}
	responseByte,respByteError:=ioutil.ReadAll(response.Body)
	if respByteError != nil {
		return nil,errors.Wrap(respByteError,"error to read response body")
	}

	return responseByte,nil


}
