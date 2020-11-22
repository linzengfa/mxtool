// Copyright (c) 2020.
// ALL Rights reserved.
// @Description ocr.go
// @Author moxiao
// @Date 2020/11/21 18:19

package baidu

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type IdCardFront struct {
	UserName   string `json:"user_name"`
	Sex        string `json:"sex"`
	IdNum      string `json:"id_num"`
	Nation     string `json:"nation"`
	BirthDay   string `json:"birth_date"`
	Address    string `json:"id_address"`
	BirthPlace string `json:"birth_place"`
}

//获取 Access Token
func (mg *MGAi) GetAccessToken() (token string, err error) {
	url := buildAccessTokenURL(mg.AppKey, mg.AppSecurity)
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

//身份证（正面）识别
//image-必填，图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式
func (mg *MGAi) IdCardFront(image, accessToken string) (idCardFront *IdCardFront, err error) {
	if len(image) == 0 || len(accessToken) == 0 {
		return nil, MXAI_PARAM_ERROR
	}
	requestURL := fmt.Sprintf("%s?access_token=%s", IDCARD_URL, accessToken)
	v := url.Values{}
	v.Add("id_card_side", "front")
	v.Add("image", image)
	resp, err := http.Post(requestURL, CONTENT_TYPE_URLENCODED, strings.NewReader(v.Encode()))
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
		errmsg := fmt.Sprintf("%v-%v", resultMap[MGOCR_RETURN_ERRCODE_KEY], resultMap[MGOCR_RETURN_ERRMSG_KEY])
		err = errors.New(errmsg)
		return
	}
	if _, ok := resultMap[MGOCR_RETURN_WORDS_RESULT_KEY]; !ok {
		err = MGOCR_IDCARD_INFO_ERROR
		return
	}
	wordsResult, ok := resultMap[MGOCR_RETURN_WORDS_RESULT_KEY].(map[string]interface{})
	if !ok {
		err = MGOCR_IDCARD_INFO_ERROR
		return
	}
	return analysisIdCardFont(wordsResult)
}

func analysisIdCardFont(wordsResult map[string]interface{}) (idCardFront *IdCardFront, err error) {
	if _, ok := wordsResult["公民身份号码"]; !ok {
		err = MGOCR_IDCARD_INFO_ERROR
		return
	}
	if _, ok := wordsResult["性别"]; !ok {
		err = MGOCR_IDCARD_INFO_ERROR
		return
	}
	if _, ok := wordsResult["民族"]; !ok {
		err = MGOCR_IDCARD_INFO_ERROR
		return
	}
	if _, ok := wordsResult["住址"]; !ok {
		err = MGOCR_IDCARD_INFO_ERROR
		return
	}
	if _, ok := wordsResult["出生"]; !ok {
		err = MGOCR_IDCARD_INFO_ERROR
		return
	}
	if _, ok := wordsResult["姓名"]; !ok {
		err = MGOCR_IDCARD_INFO_ERROR
		return
	}
	idNumMap, _ := wordsResult["公民身份号码"].(map[string]interface{})
	idNum, ok := idNumMap["words"].(string)
	if !ok || len(idNum) == 0 {
		err = MGOCR_IDCARD_INFO_ERROR
		return
	}

	sexMap, _ := wordsResult["性别"].(map[string]interface{})
	sex, ok := sexMap["words"].(string)
	if !ok || len(sex) == 0 {
		err = MGOCR_IDCARD_INFO_ERROR
		return
	}
	nationMap, _ := wordsResult["民族"].(map[string]interface{})
	nation, ok := nationMap["words"].(string)

	if !ok || len(nation) == 0 {
		err = MGOCR_IDCARD_INFO_ERROR
		return
	}

	birthDayMap, _ := wordsResult["出生"].(map[string]interface{})
	birthDay, ok := birthDayMap["words"].(string)
	if !ok || len(birthDay) == 0 {
		err = MGOCR_IDCARD_INFO_ERROR
		return
	}

	addressMap, _ := wordsResult["住址"].(map[string]interface{})
	address, ok := addressMap["words"].(string)
	if !ok || len(address) == 0 {
		err = MGOCR_IDCARD_INFO_ERROR
		return
	}

	userNameMap, _ := wordsResult["姓名"].(map[string]interface{})
	userName, ok := userNameMap["words"].(string)
	if !ok || len(userName) == 0 {
		err = MGOCR_IDCARD_INFO_ERROR
		return
	}

	birthPlace := ""
	var areaCodeMap map[string]string
	if err := json.Unmarshal([]byte(AREA_CODE), &areaCodeMap); err == nil {
		area, ok := areaCodeMap[idNum[0:6]]
		if ok && len(area) != 0 {
			birthPlace = area
		}
	}

	idCardFront = &IdCardFront{}
	idCardFront.Address = address
	idCardFront.BirthDay = birthDay[0:4] + "-" + birthDay[4:6] + "-" + birthDay[6:8]
	idCardFront.IdNum = idNum
	idCardFront.Nation = nation
	idCardFront.Sex = sex
	idCardFront.UserName = userName
	idCardFront.BirthPlace = birthPlace
	return
}