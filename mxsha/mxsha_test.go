/**********************************************
** @Des: mxsha_test.go
** @Author: moxiao
** @Date:   2020/6/20 22:44
** @Last Modified by:  moxiao
** @Last Modified time: 2020/6/20 22:44
***********************************************/
package mxsha

import "testing"

func TestSha512(t *testing.T) {
	src := "test2"
	t.Log(Sha512([]byte(src)))
}
