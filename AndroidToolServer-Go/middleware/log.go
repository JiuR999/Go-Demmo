package middleware

import (
	"AndroidToolServer-Go/common/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		path := context.Request.URL.Path
		raw := context.Request.URL.RawQuery
		context.Next()
		end := time.Now()
		//耗费时间
		latency := end.Sub(start)
		clientIp := context.ClientIP()
		method := context.Request.Method
		statusCode := context.Writer.Status()
		if raw != "" {
			path = path + "?" + raw
		}
		log.WithFieldsMsg(logrus.Fields{
			"status_code":    statusCode,
			"latency":        latency,
			"client_ip":      clientIp,
			"request_method": method,
			"request_url":    path,
		}, "")
		/*log.Default().Println("latency=", latency, "\n", "clientIp=", clientIp,
		"method=", method, "\n", "statuscode=", statusCode)*/
	}

}
