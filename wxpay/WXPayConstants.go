// Copyright (c) 2020.
// ALL Rights reserved.
// @Description WXPayConstants.go
// @Author moxiao
// @Date 2020/11/22 10:19

package wxpay

//接口URL地址
const (
	WXPAY_UNIFIEDORDER_URL = "https://api.mch.weixin.qq.com/pay/unifiedorder" //统一下单
	WXPAY_MICROPAY_URL     = "https://api.mch.weixin.qq.com/pay/micropay"     //提交刷卡支付
	WXPAY_ORDERQUER_URL     = "https://api.mch.weixin.qq.com/pay/orderquery"     //查询订单
)

//接口URL地址（沙箱）
const (
	WXPAY_SANDBOX_UNIFIEDORDER_URL = "https://api.mch.weixin.qq.com/pay/unifiedorder" //统一下单
	WXPAY_SANDBOX_MICROPAY_URL     = "https://api.mch.weixin.qq.com/pay/micropay"     //提交刷卡支付
	WXPAY_SANDBOX_ORDERQUER_URL     = "https://api.mch.weixin.qq.com/pay/orderquery"     //查询订单
)

//返回结果
const (
	SUCCESS = "SUCCESS"
	FAIL    = "FAIL"
	OK      = "OK"
)

//签名类型
const (
	SIGN_TYPE_MD5         = "MD5"
	SIGN_TYPE_HMAC_SHA256 = "HMAC-SHA256"
)

//交易类型
const (
	TRADE_TYPE_MWEB     = "MWEB"     //H5支付
	TRADE_TYPE_JSAPI    = "JSAPI"    //小程序/公众号支付
	TRADE_TYPE_APP      = "APP"      //APP支付
	TRADE_TYPE_NATIVE   = "NATIVE"   //扫码支付
	TRADE_TYPE_MICROPAY = "MICROPAY" //刷卡支付
)

//设备号
const (
	DEVICEINFO_XXC = "XXC" //默认设备类型
	DEVICEINFO_WEB = "WEB" //H5支付或公众号支付
)
const (
	DEFAULT_ORDER_BODY  = "充值"
	DEFAULT_FEETYPE     = "CNY" //默认货币类型
	K_CONTENT_TYPE_FORM = "application/x-www-form-urlencoded;charset=utf-8"
	DEFAULT_TAG         = "xml"
	DEFAULT_EXPIRE      = "1h" //默认失效时间间隔1小时
)