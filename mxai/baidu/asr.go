// Copyright (c) 2020.
// ALL Rights reserved.
// @Description ars.go
// @Author moxiao
// @Date 2020/11/21 18:19

package baidu

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//获取 Access Token
func (mg *MGAi) GetAsrAccessToken() (token string, err error) {
	url := buildAsrAccessTokenURL(mg.AppKey, mg.AppSecurity)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var resultMap map[string]interface{}
	err = json.Unmarshal(body, &resultMap)
	if err != nil {
		return
	}
	if _, ok := resultMap[MGOCR_RETURN_ERRCODE_KEY]; ok {
		errmsg := fmt.Sprintf("GetAccessToken errors，errcode=%v,errmsg=%v", resultMap[MGOCR_RETURN_ERRCODE_KEY], resultMap[MGOCR_RETURN_ERRMSG_KEY])
		err = errors.New(errmsg)
		return
	}
	if _, ok := resultMap[MGOCR_RETURN_ACCESS_TOKEN_KEY]; !ok {
		err = errors.New("GetAccessToken errors,access_token does not exist")
		return
	}
	return resultMap[MGOCR_RETURN_ACCESS_TOKEN_KEY].(string), nil
}

func (mg *MGAi) Speech(speechData string, speechDataLen int64, format, token string) (content string, err error) {
	if len(speechData) == 0 || len(token) == 0 || len(format) == 0 || speechDataLen == 0 {
		err = MXAI_PARAM_ERROR
		return
	}
	requestParamMap := make(map[string]interface{})
	requestParamMap["format"] = format
	requestParamMap["rate"] = VOICE_RATE_16000
	requestParamMap["channel"] = 1
	requestParamMap["cuid"] = mg.Cuid
	requestParamMap["token"] = token
	requestParamMap["dev_pid"] = 1537
	requestParamMap["speech"] = speechData
	requestParamMap["len"] = speechDataLen

	requestParam, err := json.Marshal(requestParamMap)
	if err != nil {
		return
	}
	fmt.Println("requestParam=", string(requestParam))
	resp, err := http.Post(ASR_SPEED_URL, CONTENT_TYPE_JSON, strings.NewReader(string(requestParam)))
	defer resp.Body.Close()
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var resultMap map[string]interface{}
	err = json.Unmarshal(body, &resultMap)
	if err != nil {
		return
	}

	result, ok := resultMap[MGASR_RETURN_RESULT_KEY].(string)
	fmt.Println(result, ok)
	if !ok || len(result) == 0 {
		fmt.Println("speech resultMap=", resultMap)
		err = MGASR_ASR_FAILURE
		return
	}
	content = result
	return
}
