/**********************************************
** @Des: mxmd5.go
** @Author: MoXiao
** @Date:   2018/9/30 10:24
** @Last Modified by:  MoXiao
** @Last Modified time: 2018/9/30 10:24
***********************************************/
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
