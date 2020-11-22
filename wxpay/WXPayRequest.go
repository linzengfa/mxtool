/**********************************************
** @Des: WXPayRequest.go
** @Author: MoXiao
** @Date:   2018/10/16 9:44
** @Last Modified by:  MoXiao
** @Last Modified time: 2018/10/16 9:44
***********************************************/
package wxpay

import (
	"bytes"
	"fmt"
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
	fmt.Println("http请求结果", string(respBody))
	return respBody, err
}
