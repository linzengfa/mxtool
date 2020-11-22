// Copyright (c) 2020.
// ALL Rights reserved.
// @Description mxqr.go
// @Author moxiao
// @Date 2020/11/22 10:19

package mxqr

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	DEFAULT_CONTENTTYPE_JSON = "application/json;charset=utf-8"

	QRCODE_WXACODE_CREATE_QRCODE_URL = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token="
	QRCODE_WXACODE_GET_URL           = "https://api.weixin.qq.com/wxa/getwxacode?access_token="
	QRCODE_WXACODE_GETUNLIMITED_URL  = "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token="

	QRCODE_MAX_WIDTH = 1280
)

var (
	ERROR_PARAM_EMPTY  = errors.New("请求参数不能为空")
	ERROR_HTTP_REQUEST = errors.New("获取小程序二维码错误")
	createLocalDirLock = new(sync.RWMutex)
)

//获取小程序二维码，适用于需要的码数量较少的业务场景。
// 通过该接口生成的小程序码，永久有效，有数量限制
//参数：path扫码进入的小程序页面路径，最大长度 128 字节，不能为空；对于小游戏，可以只传入 query 部分，来实现传参效果，
// 如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
//width:二维码的宽度，单位 px。最小 280px，最大 1280px
func CreateQRCode(accessToken, path string, width int, outputFileName, baseURL, basePath string) (imageURL string, err error) {
	if len(path) == 0 || len(outputFileName) == 0 || len(accessToken) == 0 || len(baseURL) == 0 || len(basePath) == 0 {
		err = ERROR_PARAM_EMPTY
		return
	}

	url := fmt.Sprintf("%s%s", QRCODE_WXACODE_CREATE_QRCODE_URL, accessToken)
	param := map[string]interface{}{}
	param["path"] = path
	param["width"] = QRCODE_MAX_WIDTH
	if width != 0 {
		param["width"] = width
	}
	data, err := json.Marshal(param)
	if err != nil {
		return
	}

	respBody, err := requestOnce(http.DefaultClient, url, DEFAULT_CONTENTTYPE_JSON, data)
	if err != nil {
		return
	}
	result := string(respBody)
	if strings.Contains(result, "errcode") {
		err = ERROR_HTTP_REQUEST
		return
	}
	//存储到本地
	imageURL, err = saveFile(outputFileName,baseURL, basePath, respBody)
	return
}

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
	return respBody, err
}

//保存文件
func saveFile(outputFileName, baseURL, basePath string, data []byte) (outputFileURL string, err error) {
	currDate := time.Now().Format("20060102")
	if !strings.HasSuffix(basePath, "/") {
		basePath = basePath + "/"
	}

	fullDir := basePath + currDate + "/"

	if !strings.HasSuffix(baseURL, "/") {
		baseURL = baseURL + "/"
	}

	outputFileURL = baseURL + currDate + "/" + outputFileName

	createLocalDirLock.Lock()
	_, err = os.Stat(fullDir)
	if err == nil {
	} else if os.IsNotExist(err) {
		err = os.MkdirAll(fullDir, os.ModePerm)
		if err != nil {
			return
		}
	} else {
		return
	}
	createLocalDirLock.Unlock()

	outputFile := fullDir + outputFileName
	fi, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	defer fi.Close()
	_, err = fi.Write(data)
	if err != nil {
		return
	}
	return
}
