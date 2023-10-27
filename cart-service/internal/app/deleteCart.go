package app

import (
	"context"

	"github.com/farhandwian/microservice/pb"
)

func (m *MicroserviceServer) DeleteCart(ctx context.Context, req *pb.DeleteCartRequest) (*pb.DeleteCartResponse, error) {

	return &pb.DeleteCartResponse{}, nil
}
