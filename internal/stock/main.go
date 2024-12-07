package main

import (
	"log"

	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/sjxiang/oms/common/config"
	"github.com/sjxiang/oms/common/pb"
	"github.com/sjxiang/oms/common/server"
	"github.com/sjxiang/oms/stock/service"
	"github.com/sjxiang/oms/stock/ports"
)

// 初始化
func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}

	log.Printf("全局配置, %v\n", viper.Get("stock"))
}

func main() {
	serviceName := viper.GetString("stock.service-name")
	serviceType := viper.GetString("stock.server-to-run")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := service.NewApplication(ctx)
	switch serviceType {
	case "grpc":
		server.RunGrpcServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGrpcServer(application)
			pb.RegisterStockServiceServer(server, svc)
		})
	case "http":
		// 暂时不用
	default:
		panic("unexpected server type")
	}
}