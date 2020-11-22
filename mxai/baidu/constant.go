// Copyright (c) 2020.
// ALL Rights reserved.
// @Description constant.go
// @Author moxiao
// @Date 2020/11/21 18:19

package baidu

const (
	ACCESS_TOKEN_URL     = "https://aip.baidubce.com/oauth/2.0/token"
	ASR_ACCESS_TOKEN_URL = "https://openapi.baidu.com/oauth/2.0/token"
	IDCARD_URL           = "https://aip.baidubce.com/rest/2.0/ocr/v1/idcard"
	ASR_SPEED_URL        = "http://vop.baidu.com/server_api"
)

const (
	MGOCR_RETURN_ERRCODE_KEY      = "errors"
	MGOCR_RETURN_ERRMSG_KEY       = "error_description"
	MGOCR_RETURN_ACCESS_TOKEN_KEY = "access_token"
	MGOCR_RETURN_WORDS_RESULT_KEY = "words_result"
)

const (
	MGASR_RETURN_RESULT_KEY  = "result"
	MGASR_RETURN_ERR_NO_KEY  = "err_no"
	MGASR_RETURN_ERR_MSG_KEY = "err_msg	"
)
