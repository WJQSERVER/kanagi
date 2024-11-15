package record

import (
	"io"
	"kanagi/config"
	"kanagi/logger"

	"github.com/gin-gonic/gin"
)

var (
	logw       = logger.Logw
	logInfo    = logger.LogInfo
	logWarning = logger.LogWarning
	logError   = logger.LogError
)

func Record(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求信息
		// IP METHOD FULLURL(包括参数) PROTOCOL UA
		logInfo("%s %s %s %s %s", c.ClientIP(), c.Request.Method, c.Request.URL.String(), c.Request.Proto, c.Request.Header.Get("User-Agent"))
		// 获取请求完整header
		logInfo("Header: %v", c.Request.Header)
		// 获取请求body
		body, err := io.ReadAll(c.Request.Body)
		logInfo("Body: %s", string(body))
		if err != nil {
			logError("Read request body error: %v", err)
		}
		// 返回200 OK
		c.String(200, "OK")
	}
}
