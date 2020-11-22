// Copyright (c) 2020.
// ALL Rights reserved.
// @Description WXPayXmlUtil.go
// @Author moxiao
// @Date 2020/11/22 10:19

package wxpay

import "encoding/xml"

//转化成XML
func buildXML(result interface{}) (xmlData []byte, err error) {
	if result == nil {
		return nil, WXPAY_ERROR
	}
	xmlData, err = xml.Marshal(result)
	//headerBytes := []byte(xml2.Header)
	//拼接XML头和实际XML内容
	//xml = append(headerBytes, xml...)
	return
}

//解析XML
func parseXML(xmlData []byte, result interface{}) (err error) {
	if xmlData == nil || result == nil {
		return WXPAY_ERROR
	}
	err = xml.Unmarshal(xmlData, result)
	return
}
