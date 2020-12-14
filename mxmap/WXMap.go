// Copyright (c) 2020.
// ALL Rights reserved.
// @Description WXMap.go
// @Author moxiao
// @Date 2020/12/14 19:43

package mxmap

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//请求参数错误
var (
	WXMAP_ERROR               = errors.New("请求失败")
	WXMAP_REQ_PARAMT_ERROR    = errors.New("缺少参数")
	WXMAP_JSON_MARSHA_ERROR   = errors.New("JSON生成错误")
	WXMAP_JSON_UNMARSHA_ERROR = errors.New("JSON解析错误")
	WXMAP_XML_MARSHA_ERROR    = errors.New("XML生成错误")
	WXMAP_XML_UNMARSHA_ERROR  = errors.New("XML解析错误")
	WXMAP_SIGN_CAL_ERROR      = errors.New("签名计算失败")
	WXMAP_SIGN_VERIFY_ERROR   = errors.New("签名验证失败")
)

// 微信支付

type WxMap struct {
	key        string // API调用key
	httpClient *http.Client
}

// 创建微信支付对象
func New(key string) (wx *WxMap) {
	wx = &WxMap{}
	wx.key = key
	wx.httpClient = http.DefaultClient
	return
}

// 逆地址解析
// @param lat 纬度，必填
// @param lng 经度，必填
// @param getPoi 是否返回周边POI列表 1.返回；0不返回(默认)，非必填
// @param poiOption 用于控制POI列表
func (wx *WxMap) Geocoder(lat, lng float64, getPoi, poiOption string) (geoCoderResponse *GeoCoderResponse, err error) {
	if len(wx.key) == 0 {
		err = WXMAP_REQ_PARAMT_ERROR
		return
	}
	apiUrl := formatRequestURL(wx.key, lat, lng, getPoi, poiOption)
	httpResponse, err := wx.httpClient.Get(apiUrl)
	if err != nil {
		return
	}
	if httpResponse == nil {
		err = WXMAP_ERROR
		return
	}
	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return
	}
	if httpResponse.StatusCode != 200 {
		err = fmt.Errorf("request fail with http status code: %s, with body: %s", httpResponse.Status, body)
		return
	}
	geoCoderResponse = &GeoCoderResponse{}
	err = json.Unmarshal(body, &geoCoderResponse)

	return
}

func formatRequestURL(key string, lat, lng float64, getPoi, poiOption string) string {

	sb := bytes.Buffer{}
	param := make([]interface{}, 0)
	sb.WriteString("https://apis.map.qq.com/ws/geocoder/v1/?location=%v,%v&key=%v")
	param = append(param, lat)
	param = append(param, lng)
	param = append(param, key)

	if len(getPoi) != 0 && getPoi == "1" {
		sb.WriteString("&get_poi=%v")
		param = append(param, getPoi)
	}

	if len(poiOption) != 0 {
		sb.WriteString("&poi_options=%v")
		param = append(param, poiOption)
	}
	return fmt.Sprintf(sb.String(), param...)
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
	fmt.Println("http请求结果", string(respBody))
	return respBody, err
}
