// Copyright (c) 2020.
// ALL Rights reserved.
// @Description WXPayErrors.go
// @Author moxiao
// @Date 2020/11/22 10:19

package wxpay

import "errors"

//请求参数错误
var (
	WXPAY_ERROR               = errors.New("请求失败")
	WXPAY_REQ_PARAMT_ERROR    = errors.New("缺少参数")
	WXPAY_JSON_MARSHA_ERROR   = errors.New("JSON生成错误")
	WXPAY_JSON_UNMARSHA_ERROR = errors.New("JSON解析错误")
	WXPAY_XML_MARSHA_ERROR    = errors.New("XML生成错误")
	WXPAY_XML_UNMARSHA_ERROR  = errors.New("XML解析错误")
	WXPAY_SIGN_CAL_ERROR      = errors.New("签名计算失败")
	WXPAY_SIGN_VERIFY_ERROR   = errors.New("签名验证失败")
)
//微信支付返回错误码
const (
	NOAUTH                = "NOAUTH"                //商户无此接口权限
	NOTENOUGH             = "NOTENOUGH"             //用户帐号余额不足，请用户充值或更换支付卡后再支付
	ORDERPAID             = "ORDERPAID"             //商户订单已支付，无需更多操作
	ORDERCLOSED           = "ORDERCLOSED"           //当前订单已关闭，请重新下单
	SYSTEMERROR           = "SYSTEMERROR"           //系统异常，请用相同参数重新调用
	APPID_NOT_EXIST       = "APPID_NOT_EXIST"       //请检查APPID是否正确
	MCHID_NOT_EXIST       = "MCHID_NOT_EXIST"       //请检查MCHID是否正确
	APPID_MCHID_NOT_MATCH = "APPID_MCHID_NOT_MATCH" //请确认appid和mch_id是否匹配
	LACK_PARAMS           = "LACK_PARAMS"           //请检查参数是否齐全
	OUT_TRADE_NO_USED     = "OUT_TRADE_NO_USED"     //请核实商户订单号是否重复提交
	SIGNERROR             = "SIGNERROR"             //签名错误
	XML_FORMAT_ERROR      = "XML_FORMAT_ERROR"      //XML格式错误
	REQUIRE_POST_METHOD   = "REQUIRE_POST_METHOD"   //未使用post传递参数
	POST_DATA_EMPTY       = "POST_DATA_EMPTY"       //post数据不能为空
	NOT_UTF8              = "NOT_UTF8"              //未使用指定编码格式
	PARAM_ERROR           = "PARAM_ERROR"           //参数错误
	AUTHCODEEXPIRE        = "AUTHCODEEXPIRE"        //请收银员提示用户，请用户在微信上刷新条码
	NOTSUPORTCARD         = "NOTSUPORTCARD"         //该卡不支持当前支付，提示用户换卡支付或绑新卡支付
	ORDERREVERSED         = "ORDERREVERSED"         //订单已撤销
	BANKERROR             = "BANKERROR"             //请立即调用被扫订单结果查询API，查询当前订单的不同状态，决定下一步的操作。
	USERPAYING            = "USERPAYING"            //等待5秒，然后调用被扫订单结果查询API，查询当前订单的不同状态，决定下一步的操作。
	AUTH_CODE_ERROR       = "AUTH_CODE_ERROR"       //每个二维码仅限使用一次，请刷新再试
	AUTH_CODE_INVALID     = "AUTH_CODE_INVALID"     //请扫描微信支付被扫条码/二维码
	BUYER_MISMATCH        = "BUYER_MISMATCH"        //请确认支付方是否相同
	INVALID_REQUEST       = "INVALID_REQUEST"       //请确认商户系统是否正常，是否具有相应支付权限
	TRADE_ERROR           = "TRADE_ERROR"           //请确认帐号是否存在异常
)
