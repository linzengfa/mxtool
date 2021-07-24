// Copyright (c) 2020.
// ALL Rights reserved.
// @Description errors.go
// @Author moxiao
// @Date 2020/11/21 18:19

package baidu

import (
	"errors"
)

var (
	ErrIDCardInfoEmpty = errors.New("idCard info empty")

	ErrAsrFailure = errors.New("asr failure")
)
