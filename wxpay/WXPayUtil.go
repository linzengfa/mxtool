// Copyright (c) 2020.
// ALL Rights reserved.
// @Description WXPayUtil.go
// @Author moxiao
// @Date 2020/11/22 10:19

package wxpay

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/linzengfa/mxtool/mxconv"
	"github.com/linzengfa/mxtool/mxmd5"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"time"
)

//初始化证书
func InitTransport(certFilePath, keyFilePath, rootCaPath string) (httpTransport *http.Transport, err error) {
	if len(certFilePath) == 0 || len(keyFilePath) == 0 || len(rootCaPath) == 0 {
		err = errors.New("请求参数不能为空")
	}
	certs, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		return
	}
	rootCa, err := ioutil.ReadFile(rootCaPath)
	if err != nil {
		return
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(rootCa)

	httpTransport = &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{certs},
		},
	}
	return
}

//生成随机字符串
//lens:生成字符串长度
func GetRandomString(lens int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func signStruct(data interface{}, key string) (signResult string, err error) {
	if data == nil {
		err = WXPAY_REQ_PARAMT_ERROR
		return
	}
	mapData := mxconv.StructToMapWithTag(data, DEFAULT_TAG)
	return sign(mapData, key)
}

//统一下单MD5签名
func sign(data map[string]interface{}, key string) (signResult string, err error) {
	if data == nil {
		err = WXPAY_REQ_PARAMT_ERROR
		return
	}
	signStr := createSignString(data)
	var buf bytes.Buffer
	buf.WriteString(signStr)

	if len(key) != 0 {
		buf.WriteString("&key=")
		buf.WriteString(key)
	}
	fmt.Println("签名字符串：")
	fmt.Println(buf.String())
	md5Value := mxmd5.Md5(buf.Bytes())
	signResult = strings.ToUpper(md5Value)
	return
}

//组织签名信息
func createSignString(data map[string]interface{}) (query string) {
	//不参与签名变量名
	//excludeKeys := [4]string{"pfx", "apiKey", "sign", "key"}
	keys := make([]string, 0, len(data))
	for key, _ := range data {
		if !IsEmpty(data[key]) && "key" != key && "sign" != key && "apiKey" != key {
			keys = append(keys, key)
		}
	}
	//排序,ASCII码从小到大排序
	sort.Strings(keys)
	fmt.Println(keys)
	query = ""
	for i, k := range keys {
		if i == 0 {
			query = fmt.Sprintf("%v=%v", k, data[k])
		} else {
			query = fmt.Sprintf("%v&%v=%v", query, k, data[k])
		}
	}
	return
}

//创建唯一订单号
//yyyymmddhhssmm_18位随机序列
func CreateOutTradeNo() string {
	return fmt.Sprintf("%v%v", time.Now().Format("20060102150405"), GetRandomString(18))
}

//判断是否为空
func IsEmpty(a interface{}) bool {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}
