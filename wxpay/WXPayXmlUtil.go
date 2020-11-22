/**********************************************
** @Des: WXPayXmlUtil.go
** @Author: MoXiao
** @Date:   2018/10/16 9:44
** @Last Modified by:  MoXiao
** @Last Modified time: 2018/10/16 9:44
***********************************************/
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
