package ports

import (
	"context"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	"github.com/sjxiang/oms/common/pb"
	"github.com/sjxiang/oms/stock/app"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(app app.Application) *GrpcServer {
	return &GrpcServer{app: app}
}

func (s GrpcServer) CheckIfItemsInStock(ctx context.Context, request *pb.CheckIfItemsInStockRequest) (*pb.CheckIfItemsInStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIfItemsInStock not implemented")
}
func (s GrpcServer) GetItems(ctx context.Context, request *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItems not implemented")
}