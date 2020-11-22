// Copyright (c) 2020.
// ALL Rights reserved.
// @Description mxcache.go
// @Author moxiao
// @Date 2020/11/22 10:19

package mxcache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var Che *cache.Cache

func Init(defaultExpiration, cleanupInterval time.Duration) {
	Che = cache.New(defaultExpiration, cleanupInterval)
}
