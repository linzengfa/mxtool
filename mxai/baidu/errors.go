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
	MGOCR_IDCARD_INFO_ERROR = errors.New("idCard info empty")

	MGASR_ASR_FAILURE = errors.New("asr failure")
)
