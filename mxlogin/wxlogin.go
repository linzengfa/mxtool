// Copyright (c) 2020.
// ALL Rights reserved.
// @Description mxlogin.go
// @Author moxiao
// @Date 2020/11/22 10:19

package mxlogin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/linzengfa/mgtool/mxsha"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	WXLOGIN_RETURN_ERRCODE_KEY    = "errcode"
	WXLOGIN_RETURN_ERRMSG_KEY     = "errmsg"
	WXLOGIN_RETURN_OPENID_KEY     = "openid"
	WXLOGIN_RETURN_SESSIONKEY_KEY = "session_key"
)

var (
	wxlogin_param_error           = errors.New("wxlogin_param_error")
	wxlogin_watermark_check_error = errors.New("encryptedData errors")
)

type WXLogin struct {
	AppId     string
	AppSecret string
}

//微信用户信息
type UserInfo struct {
	NickName  string `json:"nickName"`
	Gender    byte   `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	Language  string `json:"language"`
	AvatarUrl string `json:"avatarUrl"`
}

//加密数据
type EncryptedData struct {
	OpenId    string `json:"openId"`
	NickName  string `json:"nickName"`
	Gender    byte   `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
	UnionId   string `json:"unionId"`
}
type WXUserInfo struct {
	Userinfo   EncryptedData `json:"userinfo"`
	SessionKey string        `json:"skey"` //会话密钥
}

type WXAuthorization struct {
	Openid     string `json:"openid"`      //用户唯一标识
	SessionKey string `json:"session_key"` //会话密钥
	Unionid    string `json:"unionid"`     //用户在开放平台的唯一标识符。
	Errcode    int    `json:"errcode"`     //错误码
	ErrMsg     string `json:"errmsg"`      //错误信息
}

type WXPhoneNumber struct {
	PhoneNumber     string `json:"phoneNumber"`     //用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string `json:"purePhoneNumber"` //没有区号的手机号
	CountryCode     string `json:"countryCode"`     //区号
}

func New(appId, appSecret string) (wx *WXLogin, err error) {
	if appId == "" || appSecret == "" {
		err = wxlogin_param_error
		return
	}
	return &WXLogin{appId, appSecret}, nil
}

//微信数据签名校验
func (wx *WXLogin) Check(rawData, signature, session_key string) bool {
	var buffer bytes.Buffer
	buffer.WriteString(rawData)
	buffer.WriteString(session_key)
	comresult := strings.Compare(mxsha.Sha1(buffer.Bytes()), signature)
	return comresult == 0
}

//临时登录凭证code
func (wx *WXLogin) Code2Session(code string) (wxa *WXAuthorization, err error) {
	wxa, err = codeExchangeSessionkey(wx.AppId, wx.AppSecret, code)
	return
}

//微信登录
//加密数据( encryptedData )进行解密
func (wx *WXLogin) Login(encryptedData, iv, code string) (wxu *WXUserInfo, err error) {
	wxa, err := codeExchangeSessionkey(wx.AppId, wx.AppSecret, code)
	if wxa == nil || err != nil {
		return
	}
	if len(encryptedData) != 0 && len(iv) != 0 {
		wxDataCrypt := NewWXDataCrypt(wx.AppId, wxa.SessionKey)
		wxData, tmperr := wxDataCrypt.decryptUserInfoData(encryptedData, iv)
		err = tmperr
		if wxData == nil || err != nil {
			return
		}

		if len(wxData.Appid) == 0 || wxData.Appid != wx.AppId {
			err = wxlogin_watermark_check_error
			return
		}

		wxu = &WXUserInfo{
			Userinfo: EncryptedData{
				wxa.Openid,
				wxData.NickName,
				wxData.Gender,
				wxData.City,
				wxData.Province,
				wxData.Country,
				wxData.AvatarUrl,
				wxData.UnionId,
			},
			SessionKey: wxa.SessionKey,
		}
	} else {
		wxu = &WXUserInfo{
			Userinfo: EncryptedData{
				wxa.Openid,
				"",
				0,
				"",
				"",
				"",
				"",
				wxa.Unionid,
			},
			SessionKey: wxa.SessionKey,
		}
	}
	return
}

func (wx *WXLogin) GetUserInfo(accessToken, openid string) (wxu *EncryptedData, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s",
		accessToken, openid)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var resultMap map[string]interface{}
	err = json.Unmarshal(body, &resultMap)
	if err != nil {
		return
	}
	if _, ok := resultMap[WXLOGIN_RETURN_ERRCODE_KEY]; ok {
		errmsg := fmt.Sprintf("GetUserInfo errors，errcode=%v,errmsg=%v", resultMap[WXLOGIN_RETURN_ERRCODE_KEY], resultMap[WXLOGIN_RETURN_ERRMSG_KEY])
		err = errors.New(errmsg)
		return
	}
	wxu = new(EncryptedData)
	if openid, ok := resultMap["openid"].(string); ok {
		wxu.OpenId = openid
	}
	if nickName, ok := resultMap["nickname"].(string); ok {
		wxu.NickName = nickName
	}
	if sex, ok := resultMap["sex"].(byte); ok {
		wxu.Gender = sex
	}
	if province, ok := resultMap["province"].(string); ok {
		wxu.Province = province
	}
	if city, ok := resultMap["city"].(string); ok {
		wxu.City = city
	}
	if country, ok := resultMap["country"].(string); ok {
		wxu.Country = country
	}
	if headimgurl, ok := resultMap["headimgurl"].(string); ok {
		wxu.AvatarUrl = headimgurl
	}
	if unionid, ok := resultMap["unionid"].(string); ok {
		wxu.UnionId = unionid
	}

	return
}

//code换取session_key
func codeExchangeSessionkey(appId string, appSecret string, code string) (wxa *WXAuthorization, err error) {
	url := buildExchangeUrl(appId, appSecret, code)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var resultMap map[string]interface{}
	err = json.Unmarshal(body, &resultMap)
	if err != nil {
		return
	}

	//code换取session_key失败
	if _, ok := resultMap[WXLOGIN_RETURN_ERRCODE_KEY]; ok {
		errmsg := fmt.Sprintf("codeExchangeSessionkey errors，errcode=%v,errmsg=%v", resultMap[WXLOGIN_RETURN_ERRCODE_KEY], resultMap[WXLOGIN_RETURN_ERRMSG_KEY])
		err = errors.New(errmsg)
		return
	}
	if _, ok := resultMap[WXLOGIN_RETURN_SESSIONKEY_KEY]; !ok {
		err = errors.New("codeExchangeSessionkey errors,session_key does not exist")
		return
	}
	if _, ok := resultMap[WXLOGIN_RETURN_OPENID_KEY]; !ok {
		err = errors.New("codeExchangeSessionkey errors,openid does not exist")
		return
	}
	err = json.Unmarshal(body, &wxa)
	return
}

//组织code换取session_key请求url
func buildExchangeUrl(appId string, appSecret string, code string) string {
	return fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appId, appSecret, code)
}

//加密数据( encryptedData )进行解密
func (wx *WXLogin) DecryptPhoneNumber(encryptedData, iv, sessionKey string) (wpn *WXPhoneNumber, err error) {
	if len(encryptedData) == 0 || len(iv) == 0 || len(sessionKey) == 0 {
		err = wxlogin_param_error
		return
	}
	fmt.Println("DecryptPhoneNumber sessionKey", sessionKey)
	wxDataCrypt := NewWXDataCrypt(wx.AppId, sessionKey)
	wxData, err := wxDataCrypt.decryptPhoneNumberData(encryptedData, iv)
	if err != nil || wxData == nil {
		fmt.Println("DecryptPhoneNumber error", err)
		return
	}

	if len(wxData.PhoneNumber) == 0 {
		err = wxlogin_watermark_check_error
		return
	}
	fmt.Println("DecryptPhoneNumber success", wxData)
	wpn = &WXPhoneNumber{
		wxData.PhoneNumber,
		wxData.PurePhoneNumber,
		wxData.CountryCode,
	}
	return
}

//用户信息加密数据( encryptedData )进行解密
func (wx *WXLogin) DecryptUserInfo(encryptedData, iv, sessionKey string) (wxu *EncryptedData, err error) {
	if len(encryptedData) == 0 || len(iv) == 0 || len(sessionKey) == 0 {
		err = wxlogin_param_error
		return
	}

	wxDataCrypt := NewWXDataCrypt(wx.AppId, sessionKey)
	wxData, err := wxDataCrypt.decryptUserInfoData(encryptedData, iv)
	if err != nil || wxData == nil {
		fmt.Println("decryptUserInfoData error", err)
		return
	}
	if len(wxData.Appid) == 0 || wxData.Appid != wx.AppId {
		err = wxlogin_watermark_check_error
		return
	}
	wxu = &EncryptedData{
		wxData.OpenId,
		wxData.NickName,
		wxData.Gender,
		wxData.City,
		wxData.Province,
		wxData.Country,
		wxData.AvatarUrl,
		wxData.UnionId,
	}
	return
}
