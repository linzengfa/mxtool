/**********************************************
** @Des: WXPay_test.go
** @Author: MoXiao
** @Date:   2018/10/18 17:52
** @Last Modified by:  MoXiao
** @Last Modified time: 2018/10/18 17:52
***********************************************/
package wxpay

import (
	"testing"
	"time"
)

var wp *WxPay
var wp2 *WxPay

func init() {
	wp = New("wx0e8962a228d01824", "", "1481013262", "",
		"96e79218965eb72c92a549dd5a330112", "https://busapi.linzengfa.cn/api/wxPayNotify",
		"192.168.1.100", TRADE_TYPE_JSAPI, SIGN_TYPE_MD5, DEVICEINFO_XXC, false, nil)
	wp2 = New("wx6b4390fb32d88375", "", "1514758171", "1516776181",
		"c777d01eadf0a6909ebd24f6b3b33832", "https://busapi.linzengfa.cn/api/wxPayNotify",
		"192.168.1.100", TRADE_TYPE_JSAPI, SIGN_TYPE_MD5, DEVICEINFO_XXC, false, nil)
}

func TestWePay_UnifiedOrderNotify(t *testing.T) {
	reqBody := "<xml><appid><![CDATA[wx2421b1c4370ec43b]]></appid><attach><![CDATA[支付测试]]></attach><bank_type><![CDATA[CFT]]></bank_type><fee_type><![CDATA[CNY]]></fee_type><is_subscribe><![CDATA[Y]]></is_subscribe><mch_id><![CDATA[10000100]]></mch_id><nonce_str><![CDATA[5d2b6c2a8db53831f7eda20af46e531c]]></nonce_str><openid><![CDATA[oUpF8uMEb4qRXf22hE3X68TekukE]]></openid><out_trade_no><![CDATA[1409811653]]></out_trade_no><result_code><![CDATA[SUCCESS]]></result_code><return_code><![CDATA[SUCCESS]]></return_code><sign><![CDATA[B552ED6B279343CB493C5DD0D78AB241]]></sign><sub_mch_id><![CDATA[10000100]]></sub_mch_id><time_end><![CDATA[20140903131540]]></time_end><total_fee>1</total_fee><coupon_fee><![CDATA[10]]></coupon_fee><coupon_count><![CDATA[1]]></coupon_count><coupon_type><![CDATA[CASH]]></coupon_type><coupon_id><![CDATA[10000]]></coupon_id><coupon_fee><![CDATA[100]]></coupon_fee><trade_type><![CDATA[JSAPI]]></trade_type><transaction_id><![CDATA[1004400740201409030005092168]]></transaction_id></xml>"
	uor, err := wp.UnifiedOrderNotify([]byte(reqBody))

	if err != nil {
		t.Fatalf("TestWePay_UnifiedOrderNotify error", err.Error())
	}
	if uor == nil {
		t.Fatalf("TestWePay_UnifiedOrderNotify result is empty")
	}
}

func TestWePay_OrderQuery(t *testing.T) {
	transactionId := "20181019145309Vqt8sjw6oVHKJwVS2W"
	result, err := wp.OrderQuery("", transactionId)
	if err != nil {
		t.Fatalf("TestWePay_OrderQuery error", err.Error())
	}
	if result == nil {
		t.Fatalf("TestWePay_OrderQuery result is nil")
	}
	if result.OutTradeNo != transactionId {
		t.Fatalf("TestWePay_OrderQuery result transactionId got[%v],want[%v]", result.TransactionId, transactionId)
	}
}

func TestWxPay_MicroPay(t *testing.T) {
	authCode := "1350 4921 9592 5483 50"
	start := time.Now().Unix()
	outTradeNo := CreateOutTradeNo()
	result, err := wp.MicroPayWithPos(1, CreateOutTradeNo(), "芒果公交-刷卡消费", authCode, "", "", "", "", "")
	if err != nil {
		t.Fatalf("TestWxPay_MicroPay error", err.Error())
	}
	if result == nil {
		t.Fatalf("TestWxPay_MicroPay result is nil")
	}
	t.Log(result)
	end:=time.Now().Unix()
	t.Logf("耗时：%d",end-start)

	if result.OutTradeNo != outTradeNo {
		t.Fatalf("TestWxPay_MicroPay result OutTradeNo got[%v],want[%v]", result.OutTradeNo, outTradeNo)
	}
	}

func TestWxPay_MicroPay2(t *testing.T) {
	authCode := "134888058813918081"
	outTradeNo := CreateOutTradeNo()
	result, err := wp2.MicroPayWithPos(1, CreateOutTradeNo(), "芒果公交-刷卡消费", authCode, "", "", "", "", "")
	if err != nil {
		t.Fatalf("TestWxPay_MicroPay error", err.Error())
	}
	if result == nil {
		t.Fatalf("TestWxPay_MicroPay result is nil")
	}
	t.Log(result)
	if result.OutTradeNo != outTradeNo {
		t.Fatalf("TestWxPay_MicroPay result OutTradeNo got[%v],want[%v]", result.OutTradeNo, outTradeNo)
	}
}


