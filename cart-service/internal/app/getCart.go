package app

import (
	"context"
	"fmt"

	"github.com/farhandwian/microservice/pb"

	"github.com/farhandwian/microservice/internal/util"
)

func (m *MicroserviceServer) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	UserId := req.GetUserId()
	cart, err := m.cartService.GetCart(ctx, UserId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	// return &pb.GetCartResponse{}, nil
	transformedItemsPB := util.TransformCartItemsFromDataStructToPB(cart.Items)

	return &pb.GetCartResponse{
		UserId: int32(cart.UserID),
		Items:  transformedItemsPB,
	}, nil
}
