package ports

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions)() {
	router.Use(options.Middleware...)

	router.POST(options.BaseURL+"/customer/:custmer_id/orders", si.CreateOrder)
	router.GET(options.BaseURL+"/customer/:customer_id/orders/:order_id", si.GetOrder)
}

type GinServerOptions struct {
	BaseURL    string
	Middleware []gin.HandlerFunc
}

type ServerInterface interface {
	CreateOrder(c *gin.Context)
	GetOrder(c *gin.Context)	
}
