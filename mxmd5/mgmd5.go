// Copyright (c) 2020.
// ALL Rights reserved.
// @Description mxmd5.go
// @Author moxiao
// @Date 2020/11/22 10:19

package mxmd5

import (
	"crypto/md5"
	"encoding/hex"
)

//字符窜MD5计算
func Md5(src []byte) string {
	md5Install := md5.New()
	md5Install.Write(src)
	return hex.EncodeToString(md5Install.Sum(nil))
}
