/**********************************************
** @Des: WXSubscribe.go
** @Author: moxiao
** @Date:   2020/7/11 15:50
** @Last Modified by:  moxiao
** @Last Modified time: 2020/7/11 15:50
***********************************************/
package mxsubscribe

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//创建微信支付对象
func New() (wx *WxSub) {
	wx = &WxSub{}
	wx.HttpClient = http.DefaultClient
	return
}

//发送订阅消息
func (wx *WxSub) SendSubscribeMessage(accessToken string, toUserOpenId string, templateId string,
	page string, data map[string]interface{}) (errcode int, errmsg string, err error) {
	if len(accessToken) == 0 || len(toUserOpenId) == 0 || len(templateId) == 0 || data == nil {
		err = WXSUB_REQ_PARAMT_ERROR
		return
	}
	requestData := map[string]interface{}{}
	requestData["touser"] = toUserOpenId
	requestData["template_id"] = templateId
	requestData["page"] = page
	requestData["data"] = data
	body, err := json.Marshal(requestData)
	if err != nil {
		err = WXSUB_JSON_MARSHA_ERROR
		return
	}
	respBody, err := requestOnce(wx.HttpClient, buildSendSubscribeMessageURL(accessToken), CONTENT_TYPE_JSON, body)
	if err != nil {
		return
	}
	fmt.Println("发送订阅消息结果", string(respBody))
	sendRespond := sendResult{}
	json.Unmarshal(respBody, &sendRespond)
	if err != nil {
		err = WXSUB_JSON_UNMARSHA_ERROR
		return
	}
	errcode = sendRespond.ErrCode
	errmsg = sendRespond.ErrMsg
	return
}

func buildSendSubscribeMessageURL(accessToken string) string {
	return fmt.Sprintf("%s?access_token=%s", WXSUBSCRIBE_SEND_URL, accessToken)
}
