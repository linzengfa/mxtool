/**********************************************
** @Des: WXSubRequest.go
** @Author: moxiao
** @Date:   2020/7/11 15:51
** @Last Modified by:  moxiao
** @Last Modified time: 2020/7/11 15:51
***********************************************/
package mxsubscribe

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

//请求一次
func requestOnce(httpClient *http.Client, url string, contentType string, data []byte) ([]byte, error) {
	resp, err := httpClient.Post(url, contentType, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Println("http请求结果", string(respBody))
	return respBody, err
}
