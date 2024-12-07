package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/sjxiang/oms/common/config"
	"github.com/sjxiang/oms/common/pb"
	"github.com/sjxiang/oms/common/server"
	"github.com/sjxiang/oms/order/ports"
	"github.com/sjxiang/oms/order/service"
)

// 初始化
func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}

	log.Printf("全局配置, %v\n", viper.Get("order"))
}


func main() {
	serviceName := viper.GetString("order.service-name")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := service.NewApplication(ctx)

	go server.RunGrpcServer(serviceName, func(server *grpc.Server) {
		svc := ports.NewGrpcServer(application)
		pb.RegisterOrderServiceServer(server, svc)
	})

	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, nil, ports.GinServerOptions{
			BaseURL:    "/api",
			Middleware: nil,
		})
	})
}


