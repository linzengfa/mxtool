// Copyright (c) 2020.
// ALL Rights reserved.
// @Description WXPayUtil_test.go
// @Author moxiao
// @Date 2020/11/22 10:19

package wxpay

import (
	"github.com/linzengfa/mxtool/mxconv"
	"testing"
)

type Test struct {
	X int
	Y string
}

func TestCreateOutTradeNo(t *testing.T) {
	t.Log(CreateOutTradeNo())
}

func TestStructToMap(t *testing.T) {
	po := unifiedOrderRespond{}
	m := mxconv.StructToMap(po)

	t.Log(m["Sign"])

}

func Test_createSignString(t *testing.T) {

	data1 := unifiedOrderRespond{}
	data1.AppId = "wxd930ea5d5a258f4f"
	data1.MchId = "10000100"
	data1.DeviceInfo = "1000"
	data1.NonceStr = "ibuaiVcKdpRxkhJA"
	data1.ErrCode = "success"
	gotSign := createSignString(mxconv.StructToMap(data1))
	wantSign := "appid=wxd930ea5d5a258f4f&body=test&device_info=1000&mch_id=10000100&nonce_str=ibuaiVcKdpRxkhJA&key=192006250b4c09247ec02edce69f6a2d"
	if gotSign != wantSign {
		t.Fatalf("createSignString error,want[%s],got[%s]", wantSign, gotSign)
	}
	t.Log(gotSign)
}

func Test_sign(t *testing.T) {
	data1 := unifiedOrderRequest{}
	data1.AppId = "wxd930ea5d5a258f4f"
	data1.MchId = "10000100"
	data1.DeviceInfo = "1000"
	data1.NonceStr = "ibuaiVcKdpRxkhJA"
	data1.Body = "test"
	gotSign, err := sign(mxconv.StructToMapWithTag(data1, DEFAULT_TAG), "192006250b4c09247ec02edce69f6a2d")
	if err != nil {
		t.Fatal(err.Error())
	}
	wantSign := "9A0A8659F005D6984697E2CA0A9CF3B7"
	if gotSign != wantSign {
		t.Fatalf("sign error,want[%s],got[%s]", wantSign, gotSign)
	}
	t.Log(gotSign)
}

func Test_sign_2(t *testing.T) {
	xmlstr := "<xml><return_code><![CDATA[SUCCESS]]></return_code><return_msg><![CDATA[OK]]></return_msg><appid><![CDATA[wx0e8962a228d01824]]></appid><mch_id><![CDATA[1481013262]]></mch_id><device_info><![CDATA[XXC]]></device_info><nonce_str><![CDATA[EcLLNdiojeT82REj]]></nonce_str><sign><![CDATA[33598B2D3C73817983D682F8489E5F2D]]></sign><result_code><![CDATA[SUCCESS]]></result_code><prepay_id><![CDATA[wx291557057086967b14530bdd3395937153]]></prepay_id><trade_type><![CDATA[JSAPI]]></trade_type></xml>"
	data1 := unifiedOrderRespond{}
	err := parseXML([]byte(xmlstr), &data1)
	if err != nil {
		t.Fatal(err.Error())
	}

	gotSign, err := sign(mxconv.StructToMapWithTag(data1, DEFAULT_TAG), "96e79218965eb72c92a549dd5a330112")
	if err != nil {
		t.Fatal(err.Error())
	}
	wantSign := data1.Sign
	if gotSign != wantSign {
		t.Fatalf("sign error,want[%s],got[%s]", wantSign, gotSign)
	}
	t.Log(gotSign)
}
