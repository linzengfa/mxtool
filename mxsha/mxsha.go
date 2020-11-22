// Copyright (c) 2020.
// ALL Rights reserved.
// @Description mxsha.go
// @Author moxiao
// @Date 2020/11/22 10:19

package mxsha

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/hex"
	"errors"
)

func Sha1(src []byte) string {
	h := sha1.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write(src)
	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

func Sha512(src []byte) string {
	h := sha512.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write(src)
	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

func SessionID() (string, error) {
	b := make([]byte, 32)
	n, err := rand.Read(b)
	if n != len(b) || err != nil {
		return "", errors.New("Could not successfully read from the system CSPRNG")
	}
	return hex.EncodeToString(b), nil
}
