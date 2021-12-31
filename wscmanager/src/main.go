package main

import (
	"github.com/gin-gonic/gin"
	wsc_apis "wscmanager.com/apis"
	wsc_middleware "wscmanager.com/middleware" // 미들웨어 모듈
)

func main() {
	eng := gin.Default()

	eng.Use(wsc_middleware.Middleware()) // 미들웨어
	wsc_apis.APIs(eng)                   // APIs

	eng.Run(":3111")
}
