/**********************************************
** @Des: mgmd5_test.go
** @Author: MoXiao
** @Date:   2018/10/2 15:24
** @Last Modified by:  MoXiao
** @Last Modified time: 2018/10/2 15:24
***********************************************/
package mxmd5

import (
	"testing"
)

func TestMd5(t *testing.T) {
	str := "123456"
	md5 := Md5([]byte(str))
	sign := "e10adc3949ba59abbe56e057f20f883e"
	if md5 != sign {
		t.Fatalf("want=%s,got=%s", sign, md5)
	}
	t.Log(md5)
}
