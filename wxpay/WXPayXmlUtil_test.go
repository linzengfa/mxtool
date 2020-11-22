/**********************************************
** @Des: WXPayXmlUtil_test.go
** @Author: MoXiao
** @Date:   2018/10/16 17:21
** @Last Modified by:  MoXiao
** @Last Modified time: 2018/10/16 17:21
***********************************************/
package wxpay

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_buildXML(t *testing.T) {
	data1 := unifiedOrderRespond{}
	data1.AppId = "wxd930ea5d5a258f4f"
	data1.MchId = "10000100"
	data1.DeviceInfo = "1000"
	data1.NonceStr = "ibuaiVcKdpRxkhJA"
	data1.ResultCode = "SUCCESS"
	data1.ReturnCode = "SUCCESS"
	x, err := buildXML(data1)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(string(x))
	data2 := unifiedOrderRespond{}
	err = parseXML(x, &data2)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(data2.AppId)
}
func Test_parseXML(t *testing.T) {
	xmlstr := "<xml><return_code><![CDATA[FAIL]]></return_code><return_msg><![CDATA[签名错误]]></return_msg></xml>"
	data1 := unifiedOrderRespond{}

	err := parseXML([]byte(xmlstr), &data1)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(data1)

	if data1.ReturnCode != "FAIL" {
		t.Fatalf("UnifiedOrderRes test error,appid want[%v],goto[%v]", "wxd930ea5d5a258f4f", data1.AppId)
	}
	t.Log(data1.ReturnMsg)
}

func Test_reflect(t *testing.T)  {
	data1 := unifiedOrderRespond{}
	data1.Sign = "111"
	t.Log(reflect.ValueOf(&data1).Kind())
	data := unifiedOrderRespond{}
	data.Sign = "111"
	t.Log("type: ",reflect.TypeOf(data))
	t.Log("value: ",reflect.ValueOf(data))
	DoFiledAndMethod(data)
}

// 通过接口来获取任意参数，然后一一揭晓
func DoFiledAndMethod(input interface{}) {

	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}