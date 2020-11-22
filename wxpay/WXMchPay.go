// Copyright (c) 2020.
// ALL Rights reserved.
// @Description WXMchPay.go
// @Author moxiao
// @Date 2020/11/22 10:19

package wxpay

import (
	"fmt"
	"net/http"
)

//创建微信支付对象
func NewWxMchPay(mchAppId, mchId, partnerKey, spBillCreateIp, deviceInfo string, useSandbox bool, httpTransport *http.Transport) (wx *WxMchPay) {
	wx = &WxMchPay{}
	wx.mchAppId = mchAppId             //商户账号appid
	wx.mchId = mchId                   //商户号
	wx.partnerKey = partnerKey         //商户密钥
	wx.deviceInfo = deviceInfo         //设备号
	wx.spBillCreateIp = spBillCreateIp //终端IP
	wx.useSandbox = useSandbox         //是否是测试环境
	wx.httpClient = http.DefaultClient
	if httpTransport != nil { //带认证证书
		wx.httpClientWithCert = &http.Client{Transport: httpTransport}
	}
	return
}

//用于企业向微信用户个人付款
//partnerTradeNo,商户订单号(必填)
//openId,用户openid(必填)
//checkName,校验用户姓名选项(必填)
//reUserName,收款用户姓名(非必填，如果check_name设置为FORCE_CHECK，则必填用户真实姓名)
//amount,金额,企业付款金额，单位为分(必填)
//desc,企业付款备注,(必填)
//nonceStr,随机字符串，不长于32位(必填)
func (wx *WxMchPay) MktTransfers(partnerTradeNo string, openId string, checkName string, reUserName string, amount int,
	desc string, nonceStr string) (result *MktTransfersResult, err error) {
		fmt.Println("partnerTradeNo",partnerTradeNo)
		fmt.Println("openId",openId)
		fmt.Println("checkName",checkName)
		fmt.Println("amount",amount)
		fmt.Println("nonceStr",nonceStr)
	if len(partnerTradeNo) == 0 || len(openId) == 0 || (checkName == CHECK_NAME_FORCE_CHECK && len(reUserName) == 0) || !checkWxMchPay(wx) {
		err = WXPAY_REQ_PARAMT_ERROR
		return
	}
	fmt.Println("企业付款金额", amount)
	mktTransfersRequest := createMkttransfersRequest(wx)
	mktTransfersRequest.PartnerTradeNo = partnerTradeNo
	mktTransfersRequest.OpenId = openId
	mktTransfersRequest.CheckName = checkName
	mktTransfersRequest.ReUserName = reUserName
	mktTransfersRequest.Amount = amount
	mktTransfersRequest.Desc = desc
	mktTransfersRequest.NonceStr = nonceStr
	sign, err := signStruct(mktTransfersRequest, wx.partnerKey) //计算签名
	if err != nil {
		return
	}
	mktTransfersRequest.Sign = sign
	orderXml, err := buildXML(mktTransfersRequest)
	if err != nil {
		err = WXPAY_XML_MARSHA_ERROR
		return
	}
	apiUrl := WXMXCHPAY_TRANSFERS_URL
	if wx.useSandbox {
		apiUrl = WXMXCHPAY_SANDBOX_TRANSFERS_URL
	}

	respBody, err := requestOnce(wx.httpClientWithCert, apiUrl, K_CONTENT_TYPE_FORM, orderXml)
	if err != nil {
		return
	}
	fmt.Println("统一下单结果", string(respBody))
	//xml转结构体
	orderRespond := mktTransfersRespond{}
	err = parseXML(respBody, &orderRespond)
	if err != nil {
		err = WXPAY_XML_UNMARSHA_ERROR
		return
	}

	result = transfersRespondToResult(orderRespond)
	return
}

//创建统一下单信息
func createMkttransfersRequest(wx *WxMchPay) mktTransfersRequest {
	requestData := mktTransfersRequest{}
	requestData.MchAppId = wx.mchAppId
	requestData.MchId = wx.mchId
	requestData.DeviceInfo = wx.deviceInfo
	//requestData.NonceStr = GetRandomString(32)
	requestData.SpBillCreateIp = wx.spBillCreateIp
	return requestData
}

//检查参数是否完整
func checkWxMchPay(wx *WxMchPay) bool {
	if wx == nil {
		fmt.Println("checkWxMchPay nil")
		return false
	}
	if len(wx.mchAppId) == 0 || len(wx.mchId) == 0 ||
		len(wx.partnerKey) == 0 || len(wx.spBillCreateIp) == 0 {
		return false
	}
	return true
}
