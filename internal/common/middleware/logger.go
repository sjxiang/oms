package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


func StandardLog(l *logrus.Entry) gin.HandlerFunc {
	return func(c *gin.Context) {		
		c.Next()
	}
}