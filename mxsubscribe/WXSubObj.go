// Copyright (c) 2020.
// ALL Rights reserved.
// @Description WXSubObj.go
// @Author moxiao
// @Date 2020/11/22 10:19

package mxsubscribe

import (
	"errors"
	"net/http"
)

//请求参数错误
var (
	WXSUB_ERROR               = errors.New("请求失败")
	WXSUB_REQ_PARAMT_ERROR    = errors.New("缺少参数")
	WXSUB_JSON_MARSHA_ERROR   = errors.New("JSON生成错误")
	WXSUB_JSON_UNMARSHA_ERROR = errors.New("JSON解析错误")
	WXSUB_XML_MARSHA_ERROR    = errors.New("XML生成错误")
	WXSUB_XML_UNMARSHA_ERROR  = errors.New("XML解析错误")
	WXSUB_SIGN_CAL_ERROR      = errors.New("签名计算失败")
	WXSUB_SIGN_VERIFY_ERROR   = errors.New("签名验证失败")
)

//微信订阅消息
type WxSub struct {
	HttpClient         *http.Client
	HttpClientWithCert *http.Client
}

//应答结果
type sendResult struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
