// Copyright (c) 2020.
// ALL Rights reserved.
// @Description common.go
// @Author moxiao
// @Date 2020/11/21 18:19

package aliyun

import "errors"

const (
	MethodPost string = "POST"
	MethodGet  string = "GET"
)

var (
	ErrRequestParam     = errors.New("request param errors")
	ErrAsrErrors        = errors.New("asr errors")
	ErrAsrFailure       = errors.New("asr failure")
	ErrHttpRequestError = errors.New("http request errors")
)
