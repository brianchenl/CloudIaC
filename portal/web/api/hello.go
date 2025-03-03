// Copyright 2021 CloudJ Company Limited. All rights reserved.

package api

import (
	"cloudiac/common"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"

	"cloudiac/portal/libs/ctx"
)

// 简单 handler 函数
func Hello(c *ctx.GinRequest) {
	c.JSON(http.StatusOK, "", gin.H{
		"hello": "world",
		"goos":  runtime.GOOS,
	})
}

func Health(c *ctx.GinRequest) {
	c.JSON(http.StatusOK, nil, gin.H{
		"ok":      true,
		"version": common.VERSION,
		"build":   common.BUILD,
	})
}
