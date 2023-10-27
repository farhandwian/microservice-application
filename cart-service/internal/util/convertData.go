package util

import (
	"github.com/farhandwian/microservice/pb"

	"github.com/farhandwian/microservice/internal/datastruct"
	"github.com/farhandwian/microservice/internal/dto"
)

func TransformCartItemsFromPBtoDTO(items []*pb.CartItem) []*dto.CartItem {
	var dtoItems []*dto.CartItem
	for _, item := range items {
		dtoItem := &dto.CartItem{
			ProductId:   item.GetProductId(),
			ProductName: item.GetProductName(),
			Price:       item.GetPrice(),
			Description: item.GetDescription(),
			Amounts:     item.GetAmounts(),
			Image:       item.GetImage(),
			Status:      item.GetStatus(),
		}
		dtoItems = append(dtoItems, dtoItem)
	}
	return dtoItems
}

func TransformCartItemsFromDataStructToPB(items []datastruct.CartItem) []*pb.CartItem {
	var pbItems []*pb.CartItem
	for _, item := range items {
		pbItem := &pb.CartItem{
			ProductId:   item.ProductId,
			ProductName: item.ProductName,
			Price:       item.Price,
			Description: item.Description,
			Amounts:     item.Amounts,
			Image:       item.Image,
			Status:      item.Status,
		}
		pbItems = append(pbItems, pbItem)
	}
	return pbItems

}

func TransformCartItemsFromDtoToDataStruct(items []*dto.CartItem) []datastruct.CartItem {
	var datastructItems []datastruct.CartItem
	for _, item := range items {
		datastructItem := datastruct.CartItem{
			ProductId:   item.ProductId,
			ProductName: item.ProductName,
			Price:       item.Price,
			Description: item.Description,
			Amounts:     item.Amounts,
			Image:       item.Image,
			Status:      item.Status,
		}
		datastructItems = append(datastructItems, datastructItem)
	}
	return datastructItems
}
