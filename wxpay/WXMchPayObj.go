// Copyright (c) 2020.
// ALL Rights reserved.
// @Description WXMchPayObj.go
// @Author moxiao
// @Date 2020/11/22 10:19

package wxpay

import "net/http"

//微信企业付款
type WxMchPay struct {
	mchAppId              string //商户账号appid
	mchId              string //商户号
	partnerKey         string //商户密钥
	deviceInfo         string //设备号
	spBillCreateIp     string //终端IP
	useSandbox         bool   //是否是测试环境
	httpClient         *http.Client
	httpClientWithCert *http.Client
}

//企业付款-请求信息
type mktTransfersRequest struct {
	XMLName        struct{} `xml:"xml"`
	MchAppId       string   `xml:"mch_appid"`        //申请商户号的appid或商户号绑定的appid
	MchId          string   `xml:"mchid"`           //微信支付分配的商户号
	DeviceInfo     string   `xml:"device_info"`      //设备号
	NonceStr       string   `xml:"nonce_str"`        //随机字符串
	Sign           string   `xml:"sign"`             //签名
	PartnerTradeNo string   `xml:"partner_trade_no"` //商户订单号
	OpenId         string   `xml:"openid"`           //用户标识
	CheckName      string   `xml:"check_name"`       //校验用户姓名选项
	ReUserName     string   `xml:"re_user_name"`     //收款用户姓名
	Amount         int   `xml:"amount"`           //企业付款金额，单位为分
	Desc           string   `xml:"desc"`             //企业付款备注
	SpBillCreateIp string   `xml:"spbill_create_ip"` //终端IP
}

//企业付款-响应信息
type mktTransfersRespond struct {
	XMLName        struct{} `xml:"xml"`
	ReturnCode     string   `xml:"return_code"`      //返回状态码
	ReturnMsg      string   `xml:"return_msg"`       //返回信息
	MchAppId       string   `xml:"mch_appid"`        //商户appid
	MchId          string   `xml:"mch_id"`           //商户号
	DeviceInfo     string   `xml:"device_info"`      //设备号
	NonceStr       string   `xml:"nonce_str"`        //随机字符串
	ResultCode     string   `xml:"result_code"`      //业务结果
	ErrCode        string   `xml:"err_code"`         //错误代码
	ErrCodeDes     string   `xml:"err_code_des"`     //错误代码描述
	PartnerTradeNo string   `xml:"partner_trade_no"` //商户订单号
	PaymentNo      string   `xml:"payment_no"`       //微信付款单号
	PaymentTime    string   `xml:"payment_time"`     //付款成功时间
}

//企业付款-返回结果
type MktTransfersResult struct {
	ReturnCode     string //返回状态码
	ReturnMsg      string //返回信息
	MchAppId       string //商户appid
	MchId          string //商户号
	DeviceInfo     string //设备号
	NonceStr       string //随机字符串
	ResultCode     string //业务结果
	ErrCode        string //错误代码
	ErrCodeDes     string //错误代码描述
	PartnerTradeNo string //商户订单号
	PaymentNo      string //微信付款单号
	PaymentTime    string //付款成功时间
}

func transfersRespondToResult(respond mktTransfersRespond) *MktTransfersResult {
	result := &MktTransfersResult{}
	result.ReturnCode = respond.ReturnCode         //返回状态码
	result.ReturnMsg = respond.ReturnMsg           //返回信息
	result.MchAppId = respond.MchAppId             //商户appid
	result.MchId = respond.MchId                   //商户号
	result.DeviceInfo = respond.DeviceInfo         //设备号
	result.NonceStr = respond.NonceStr             //随机字符串
	result.ResultCode = respond.ResultCode         //业务结果
	result.ErrCode = respond.ErrCode               //错误代码
	result.ErrCodeDes = respond.ErrCodeDes         //错误代码描述
	result.PartnerTradeNo = respond.PartnerTradeNo //商户订单号
	result.PaymentNo = respond.PaymentNo           //微信付款单号
	result.PaymentTime = respond.PaymentTime       //付款成功时间
	return result
}
