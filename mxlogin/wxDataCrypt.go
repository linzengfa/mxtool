// Copyright (c) 2020.
// ALL Rights reserved.
// @Description wxDataCrypt.go
// @Author moxiao
// @Date 2020/11/22 10:19

package mxlogin

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/linzengfa/mxtool/mxaes"
)

type wxDataCrypt struct {
	appId      string
	sessionKey string
}
type watermark struct {
	Appid     string `json:"appid"`
	timestamp int64  `json:"timestamp"`
}

type wxEncryptedData struct {
	OpenId    string `json:"openId"`
	NickName  string `json:"nickName"`
	Gender    byte   `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
	UnionId   string `json:"unionId"`
	watermark `json:"watermark"`
}

type wxEncryptedPhoneNumberData struct {
	PhoneNumber     string `json:"phoneNumber"`     //用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string `json:"purePhoneNumber"` //没有区号的手机号
	CountryCode     string `json:"countryCode"`     //区号
	watermark       `json:"watermark"`
}

func NewWXDataCrypt(appId string, sessionKey string) *wxDataCrypt {
	return &wxDataCrypt{appId, sessionKey}
}

func (wx *wxDataCrypt) decryptUserInfoData(encryptedData string, iv string) (wxed *wxEncryptedData, err error) {
	origData, err := decryptData(wx.sessionKey, encryptedData, iv)
	if err != nil {
		return
	}
	err = json.Unmarshal(origData, &wxed)
	return
}

func (wx *wxDataCrypt) decryptPhoneNumberData(encryptedData string, iv string) (wxed *wxEncryptedPhoneNumberData, err error) {
	origData, err := decryptData(wx.sessionKey, encryptedData, iv)
	if err != nil {
		return
	}
	fmt.Println("decryptData success222", string(origData[:]))
	//wxed = new(wxEncryptedPhoneNumberData)
	err = json.Unmarshal(origData[:], &wxed)
	fmt.Println("decryptData success55", err)
	return
}

func decryptData(sessionKey, encryptedData, iv string) (decryptData []byte, err error) {
	sessionKeyByte, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return
	}
	encryptedDataByte, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return
	}

	ivByte, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return
	}

	decryptData, err = mxaes.Decrypt(encryptedDataByte, sessionKeyByte, ivByte)
	return
}
