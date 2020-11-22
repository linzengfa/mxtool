// Copyright (c) 2020.
// ALL Rights reserved.
// @Description WXPayObj.go
// @Author moxiao
// @Date 2020/11/22 10:19

package wxpay

import "net/http"

//微信支付
type WxPay struct {
	AppId              string //小程序appId
	SubAppId           string //子商户公众账号ID,否
	MchId              string //商户号
	SubMchId           string //子商户号，微信支付分配的子商户号，开发者模式下必填
	PartnerKey         string //商户密钥
	DeviceInfo         string //设备号
	FeeType            string //货币类型
	SpBillCreateIp     string //终端IP
	SignType           string //签名类型
	MachCert           []byte //商户证书
	UseSandbox         bool   //是否是测试环境
	NotifyUrl          string //通知地址
	TradeType          string //交易类型
	HttpClient         *http.Client
	HttpClientWithCert *http.Client
}

//统一下单-请求信息
type unifiedOrderRequest struct {
	XMLName        struct{} `xml:"xml"`
	AppId          string   `xml:"appid"`            //小程序ID
	MchId          string   `xml:"mch_id"`           //商户号
	DeviceInfo     string   `xml:"device_info"`      //设备号
	NonceStr       string   `xml:"nonce_str"`        //随机字符串
	Sign           string   `xml:"sign"`             //签名
	SignType       string   `xml:"sign_type"`        //签名类型
	Body           string   `xml:"body"`             //商品描述
	Detail         string   `xml:"detail"`           //商品详情
	Attach         string   `xml:"attach"`           //附加数据
	OutTradeNo     string   `xml:"out_trade_no"`     //商户订单号
	FeeType        string   `xml:"fee_type"`         //标价币种
	TotalFee       int64    `xml:"total_fee"`        //标价金额
	SpBillCreateIp string   `xml:"spbill_create_ip"` //终端IP
	TimeStart      string   `xml:"time_start"`       //交易起始时间
	TimeExpire     string   `xml:"time_expire"`      //交易结束时间
	GoodsTag       string   `xml:"goods_tag"`        //订单优惠标记
	NotifyUrl      string   `xml:"notify_url"`       //通知地址
	TradeType      string   `xml:"trade_type"`       //交易类型
	ProductId      string   `xml:"product_id"`       //商品ID
	LimitPay       string   `xml:"limit_pay"`        //指定支付方式
	OpenId         string   `xml:"openid"`           //用户标识
	SceneInfo      string   `xml:"scene_info"`       //场景信息
}

//统一下单-响应信息
type unifiedOrderRespond struct {
	XMLName    struct{} `xml:"xml"`
	ReturnCode string   `xml:"return_code"`  //返回状态码
	ReturnMsg  string   `xml:"return_msg"`   //返回信息
	AppId      string   `xml:"appid"`        //小程序ID
	MchId      string   `xml:"mch_id"`       //商户号
	DeviceInfo string   `xml:"device_info"`  //设备号
	NonceStr   string   `xml:"nonce_str"`    //随机字符串
	Sign       string   `xml:"sign"`         //签名
	ResultCode string   `xml:"result_code"`  //业务结果
	ErrCode    string   `xml:"err_code"`     //错误代码
	ErrCodeDes string   `xml:"err_code_des"` //错误代码描述
	TradeType  string   `xml:"trade_type"`   //交易类型
	PrepayId   string   `xml:"prepay_id"`    //预支付交易会话标识
	CodeUrl    string   `xml:"code_url"`     //二维码链接（小程序、APP、公众号支付返回）
	MwebUrl    string   `xml:"mweb_url"`     //支付跳转链接（H5支付返回）
}

//统一下单-返回结果
type UnifiedOrderResult struct {
	ReturnCode     string            //返回状态码
	ReturnMsg      string            //返回信息
	ResultCode     string            //业务结果
	ErrCode        string            //错误代码
	ErrCodeDes     string            //错误代码描述
	PrepayId       string            //预支付交易会话标识
	PrepayUrl      string            //二维码链接或者支付跳转链接
	RequestPayment *WXRequestPayment //小程序支付
}

//小程序调起支付信息
type WXRequestPayment struct {
	AppId     string `xml:"appId"`     //小程序ID
	TimeStamp string `xml:"timeStamp"` //时间戳
	NonceStr  string `xml:"nonceStr"`  //随机字符串
	Package   string `xml:"package"`   //数据包
	SignType  string `xml:"signType"`  //签名方式
	Sign      string `xml:"sign"`      //签名
}

//统一下单-支付结果通知
type UnifiedOrderNotifyResult struct {
	XMLName            struct{} `xml:"xml"`
	ReturnCode         string   `xml:"return_code"`          //返回状态码
	ReturnMsg          string   `xml:"return_msg"`           //返回信息
	AppId              string   `xml:"appid"`                //小程序ID
	MchId              string   `xml:"mch_id"`               //商户号
	DeviceInfo         string   `xml:"device_info"`          //设备号
	NonceStr           string   `xml:"nonce_str"`            //随机字符串
	Sign               string   `xml:"sign"`                 //签名
	ResultCode         string   `xml:"result_code"`          //业务结果
	ErrCode            string   `xml:"err_code"`             //错误代码
	ErrCodeDes         string   `xml:"err_code_des"`         //错误代码描述
	Openid             string   `xml:"openid"`               //用户标识
	IsSubscribe        string   `xml:"is_subscribe"`         //是否关注公众账号
	TradeType          string   `xml:"trade_type"`           //交易类型
	BankType           string   `xml:"bank_type"`            //付款银行
	TotalFee           string   `xml:"total_fee"`            //订单金额
	SettlementTotalFee string   `xml:"settlement_total_fee"` //应结订单金额
	FeeType            string   `xml:"fee_type"`             //标价币种
	CashFee            string   `xml:"cash_fee"`             //现金支付金额
	CashFeeType        string   `xml:"cash_fee_type"`        //现金支付币种
	CouponFee          string   `xml:"coupon_fee"`           //代金券金额
	CouponCount        string   `xml:"coupon_count"`         //代金券使用数量
	TransactionId      string   `xml:"transaction_id"`       //微信支付订单号
	OutTradeNo         string   `xml:"out_trade_no"`         //商户订单号
	Attach             string   `xml:"attach"`               //附加数据
	TimeEnd            string   `xml:"time_end"`             //支付完成时间
}

//提交刷卡支付-请求信息
type micropayRequest struct {
	XMLName        struct{} `xml:"xml"`
	AppId          string   `xml:"appid"`            //小程序ID
	SubAppId       string   `xml:"sub_appid"`        //子商户公众账号ID,否
	MchId          string   `xml:"mch_id"`           //商户号
	SubMchId       string   `xml:"sub_mch_id"`       //子商户号，微信支付分配的子商户号，开发者模式下必填
	DeviceInfo     string   `xml:"device_info"`      //设备号
	NonceStr       string   `xml:"nonce_str"`        //随机字符串
	Sign           string   `xml:"sign"`             //签名
	SignType       string   `xml:"sign_type"`        //签名类型
	Body           string   `xml:"body"`             //商品描述
	Detail         string   `xml:"detail"`           //商品详情
	Attach         string   `xml:"attach"`           //附加数据
	OutTradeNo     string   `xml:"out_trade_no"`     //商户订单号
	FeeType        string   `xml:"fee_type"`         //标价币种
	TotalFee       int64    `xml:"total_fee"`        //标价金额
	SpBillCreateIp string   `xml:"spbill_create_ip"` //终端IP
	TimeStart      string   `xml:"time_start"`       //交易起始时间
	TimeExpire     string   `xml:"time_expire"`      //交易结束时间
	GoodsTag       string   `xml:"goods_tag"`        //订单优惠标记
	LimitPay       string   `xml:"limit_pay"`        //指定支付方式
	AuthCode       string   `xml:"auth_code"`        //授权码
	SceneInfo      string   `xml:"scene_info"`       //场景信息
}

//提交刷卡支付-返回结果
type MicropayResult struct {
	XMLName            struct{} `xml:"xml"`
	ReturnCode         string   `xml:"return_code"`          //返回状态码
	ReturnMsg          string   `xml:"return_msg"`           //返回信息
	AppId              string   `xml:"appid"`                //小程序ID
	SubAppId           string   `xml:"sub_appid"`            //子商户公众账号ID,否
	MchId              string   `xml:"mch_id"`               //商户号
	SubMchId           string   `xml:"sub_mch_id"`           //子商户号，微信支付分配的子商户号，开发者模式下必填
	DeviceInfo         string   `xml:"device_info"`          //设备号
	NonceStr           string   `xml:"nonce_str"`            //随机字符串
	Sign               string   `xml:"sign"`                 //签名
	ResultCode         string   `xml:"result_code"`          //业务结果
	ErrCode            string   `xml:"err_code"`             //错误代码
	ErrCodeDes         string   `xml:"err_code_des"`         //错误代码描述
	Openid             string   `xml:"openid"`               //用户标识
	IsSubscribe        string   `xml:"is_subscribe"`         //是否关注公众账号
	SubOpenid          string   `xml:"sub_openid"`           //子商户appid下用户唯一标识，如需返回则请求时需要传sub_appid
	TradeType          string   `xml:"trade_type"`           //交易类型
	BankType           string   `xml:"bank_type"`            //付款银行
	FeeType            string   `xml:"fee_type"`             //货币类型
	TotalFee           string   `xml:"total_fee"`            //订单金额
	SettlementTotalFee string   `xml:"settlement_total_fee"` //应结订单金额
	CouponFee          string   `xml:"coupon_fee"`           //代金券金额
	CashFeeType        string   `xml:"cash_fee_type"`        //现金支付币种
	CashFee            string   `xml:"cash_fee"`             //现金支付金额
	TransactionId      string   `xml:"transaction_id"`       //微信支付订单号
	OutTradeNo         string   `xml:"out_trade_no"`         //商户订单号
	Attach             string   `xml:"attach"`               //附加数据
	TimeEnd            string   `xml:"time_end"`             //支付完成时间
	PromotionDetail    string   `xml:"promotion_detail"`     //营销详情
}

//账单查询-请求信息
type orderQueryRequest struct {
	XMLName       struct{} `xml:"xml"`
	AppId         string   `xml:"appid"`          //小程序ID
	MchId         string   `xml:"mch_id"`         //商户号
	TransactionId string   `xml:"transaction_id"` //微信支付订单号
	OutTradeNo    string   `xml:"out_trade_no"`   //商户订单号
	NonceStr      string   `xml:"nonce_str"`      //随机字符串
	Sign          string   `xml:"sign"`           //签名
	SignType      string   `xml:"sign_type"`      //签名类型
}

//查询订单-返回信息
type OrderqueryResult struct {
	XMLName            struct{} `xml:"xml"`
	ReturnCode         string   `xml:"return_code"`          //返回状态码
	ReturnMsg          string   `xml:"return_msg"`           //返回信息
	AppId              string   `xml:"appid"`                //小程序ID
	MchId              string   `xml:"mch_id"`               //商户号
	DeviceInfo         string   `xml:"device_info"`          //设备号
	NonceStr           string   `xml:"nonce_str"`            //随机字符串
	Sign               string   `xml:"sign"`                 //签名
	ResultCode         string   `xml:"result_code"`          //业务结果
	ErrCode            string   `xml:"err_code"`             //错误代码
	ErrCodeDes         string   `xml:"err_code_des"`         //错误代码描述
	Openid             string   `xml:"openid"`               //用户标识
	IsSubscribe        string   `xml:"is_subscribe"`         //是否关注公众账号
	TradeType          string   `xml:"trade_type"`           //交易类型
	TradeState         string   `xml:"trade_state"`          //交易状态
	BankType           string   `xml:"bank_type"`            //付款银行
	TotalFee           string   `xml:"total_fee"`            //订单金额
	SettlementTotalFee string   `xml:"settlement_total_fee"` //应结订单金额
	FeeType            string   `xml:"fee_type"`             //货币类型
	CashFeeType        string   `xml:"cash_fee_type"`        //现金支付币种
	CashFee            string   `xml:"cash_fee"`             //现金支付金额
	CouponFee          string   `xml:"coupon_fee"`           //代金券金额
	CouponCount        string   `xml:"coupon_count"`         //代金券使用数量
	TransactionId      string   `xml:"transaction_id"`       //微信支付订单号
	OutTradeNo         string   `xml:"out_trade_no"`         //商户订单号
	Attach             string   `xml:"attach"`               //附加数据
	TimeEnd            string   `xml:"time_end"`             //支付完成时间
	TradeStateDesc     string   `xml:"trade_state_desc"`     //交易状态描述
}

//关闭订单-请求信息
type CloseOrder struct {
	XMLName       struct{} `xml:"xml"`
	AppId         string   `xml:"appid"`          //小程序ID
	MchId         string   `xml:"mch_id"`         //商户号
	TransactionId string   `xml:"transaction_id"` //微信支付订单号
	OutTradeNo    string   `xml:"out_trade_no"`   //商户订单号
	NonceStr      string   `xml:"nonce_str"`      //随机字符串
	Sign          string   `xml:"sign"`           //签名
	SignType      string   `xml:"sign_type"`      //签名类型
}

//通用-请求返回信息
type CommonReturn struct {
	ReturnCode string `xml:"return_code"` //返回状态码
	ReturnMsg  string `xml:"return_msg"`  //返回信息
}