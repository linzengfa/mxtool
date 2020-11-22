// Copyright (c) 2020.
// ALL Rights reserved.
// @Description mxsha_test.go
// @Author moxiao
// @Date 2020/11/22 10:19

package mxsha

import "testing"

func TestSha512(t *testing.T) {
	src := "test2"
	t.Log(Sha512([]byte(src)))
}
