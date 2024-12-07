package ports

import (
	"context"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"github.com/sjxiang/oms/common/pb"
	"github.com/sjxiang/oms/order/app"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(app app.Application) *GrpcServer {
	return &GrpcServer{app: app}
}

func (s GrpcServer) CreateOrder(ctx context.Context, request *pb.CreateOrderRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}

func (s GrpcServer) GetOrder(ctx context.Context, request *pb.GetOrderRequest) (*pb.Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}

func (s GrpcServer) UpdateOrder(ctx context.Context, request *pb.Order) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
