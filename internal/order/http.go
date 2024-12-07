package main

import (
	// "github.com/gin-gonic/gin"
	"github.com/sjxiang/oms/order/app"
)


type HTTPServer struct {
	app app.Application
}

func NewHTTPServer(app app.Application) HTTPServer {
	return HTTPServer{
		app: app,
	}
}
// // 创建订单
// func (s *HTTPServer) CreateOrder(c *gin.Context) {
// }

// // 订单列表
// func (s *HTTPServer) CreateOrderList(c *gin.Context) {
// }


// type CreateOrderRequest struct {
// 	OrderID string `uri:"order_id" binding:"required"`
// }

// type CreateOrderResponse struct {
// 	OrderID string `json:"order_id"`
// }

// // 获取订单
// func (s *HTTPServer) CreateOrder(c *gin.Context, customID string, orderID string) {

// }

// type GetOrderRequest struct {
// 	OrderID string `uri:"order_id" binding:"required"`
// }

// type GetOrderResponse struct {
// 	OrderID string `json:"order_id"`
// }

// // 创建订单
// func (s *HTTPServer) GetOrder(c *gin.Context, customID string) {

// }
