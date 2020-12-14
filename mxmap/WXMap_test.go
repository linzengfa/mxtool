// Copyright (c) 2020.
// ALL Rights reserved.
// @Description WXMap_test.go
// @Author moxiao
// @Date 2020/12/14 20:24

package mxmap

import "testing"

func TestFormatRequestURL(t *testing.T) {
	t.Log(formatRequestURL("123", 39.984154, 116.307490, "", ""))
	t.Log(formatRequestURL("123", 39.984154, 116.307490, "1", ""))
	t.Log(formatRequestURL("123", 39.984154, 116.307490, "", "poi_options=address_format=short"))
	t.Log(formatRequestURL("123", 39.984154, 116.307490, "0", "address_format=short;radius=5000;\npage_size=20;page_index=1;policy=2"))
}

func TestGeocoder(t *testing.T) {
	wx := New("填写腾讯位置服务密钥")
	result, err := wx.Geocoder(24.571792, 118.073502, "", "")
	if err != nil {
		t.Error(err)
		return
	}
	if result == nil {
		t.Error("转换失败")
		return
	}
	t.Log(result.Message)
	t.Log(result.Result.AdInfo.AdCode)
	t.Log(result.Result.AdInfo.Name)
}
