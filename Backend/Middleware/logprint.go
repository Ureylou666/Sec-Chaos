package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func LogPrint() gin.HandlerFunc {
	logger := logrus.New()
	filePath := "/Users/ureylou/Downloads/center/Backend/Log/Server.log"
	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err:", err)
	}
	logger.Out = src
	return func(c *gin.Context) {
		stratTime := time.Now()
		c.Next()
		stopTime := time.Since(stratTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "Unknown"
		}
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"RequestTime": stratTime.Format("20060102150405"),
			"Method":      method,
			"Path":        path,
			"Status":      statusCode,
			"SpendTime":   spendTime,
			"ClientIP":    clientIP,
			"HostName":    hostName,
			"UserAgent":   userAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
