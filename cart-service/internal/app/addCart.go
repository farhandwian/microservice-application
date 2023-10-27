package app

import (
	"context"

	"github.com/farhandwian/microservice/pb"

	"github.com/farhandwian/microservice/internal/dto"
	"github.com/farhandwian/microservice/internal/util"
)

func (m *MicroserviceServer) CreateCart(ctx context.Context, req *pb.AddCartRequest) (*pb.AddCartResponse, error) {
	transformedItemsDTO := util.TransformCartItemsFromPBtoDTO(req.GetItems())

	cart, err := m.cartService.CreateCart(ctx, dto.Cart{
		UserID: int32(req.GetUserId()),
		Items:  transformedItemsDTO,
	})
	transformedItemsPB := util.TransformCartItemsFromDataStructToPB(cart.Items)

	if err != nil {
		return nil, err
	}

	return &pb.AddCartResponse{UserId: int32(cart.UserID), Item: transformedItemsPB}, nil
}
