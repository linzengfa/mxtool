// Copyright (c) 2020.
// ALL Rights reserved.
// @Description common.go
// @Author moxiao
// @Date 2020/11/21 18:19

package aliyun

import "errors"

const (
	METHOD_POST             string = "POST"
	METHOD_GET              string = "GET"
)

var (
	MXAI_PARAM_ERROR        = errors.New("request param errors")
	MXAI_ASR_ERROR          = errors.New("asr errors")
	MXAI_ASR_FAILURE        = errors.New("asr failure")
	MXAI_HTTP_REQUEST_ERROR = errors.New("http request errors")
)
