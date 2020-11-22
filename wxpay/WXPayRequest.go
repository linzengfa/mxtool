// Copyright (c) 2020.
// ALL Rights reserved.
// @Description WXPayRequest.go
// @Author moxiao
// @Date 2020/11/22 10:19

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
