package server

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_tags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
)

func init() {
	logger := logrus.New()
	logger.SetLevel(logrus.WarnLevel)
	grpc_logrus.ReplaceGrpcLogger(logrus.NewEntry(logger))
}


// 启动 grpc 服务
func RunGrpcServer(serviceName string, registerServer func(server *grpc.Server)) {
	addr := viper.Sub(serviceName).GetString("grpc-addr")
	if addr == "" {
		// TODO: Warning log
		addr = viper.GetString("fallback-grpc-addr")
	}
	RunGrpcServerOnAddr(addr, registerServer)
}

func RunGrpcServerOnAddr(addr string, registerServer func(server *grpc.Server)) {
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			grpc_tags.UnaryServerInterceptor(grpc_tags.WithFieldExtractor(grpc_tags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
		),
		grpc.ChainStreamInterceptor(
			grpc_tags.StreamServerInterceptor(grpc_tags.WithFieldExtractor(grpc_tags.CodeGenRequestFieldExtractor)),
	 		grpc_logrus.StreamServerInterceptor(logrusEntry),
		),
	)

	// 注册服务
	registerServer(grpcServer)

	// 监听端口
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Panic(err)
	}

	logrus.Infof("Starting gRPC server, Listening: %s", addr)
	
	if err := grpcServer.Serve(ln); err != nil {
		logrus.Panic(err)
	}
}