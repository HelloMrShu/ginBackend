package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求前获取当前时间
		beginTime := time.Now()

		//请求处理
		c.Next()

		//处理后获取消耗时间
		interval := time.Since(beginTime)
		url := c.Request.URL.String()
		fmt.Printf("Request URL: %s Time Interval: %v\n", url, interval)
	}
}