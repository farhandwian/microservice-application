package service

import (
	"context"

	"github.com/farhandwian/microservice/internal/dto"
	"github.com/farhandwian/microservice/internal/repository"

	"github.com/farhandwian/microservice/internal/datastruct"
	"github.com/farhandwian/microservice/internal/util"
)

type CartService interface {
	GetCart(ctx context.Context, user_id int32) (*datastruct.Cart, error)
	CreateCart(ctx context.Context, cart dto.Cart) (*datastruct.Cart, error)
	UpdateCart(ctx context.Context, cart dto.Cart) (*datastruct.Cart, error)
	DeleteCart(ctx context.Context, user_id int32) error
}

type cartService struct {
	dao repository.DAO
}

func NewCartService(dao repository.DAO) *cartService {
	return &cartService{dao: dao}
}

func (c *cartService) GetCart(ctx context.Context, user_id int32) (*datastruct.Cart, error) {
	carts, err := c.dao.NewCartQuery().GetCart(ctx, user_id)
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (c *cartService) CreateCart(ctx context.Context, cart dto.Cart) (*datastruct.Cart, error) {
	transformedItem := util.TransformCartItemsFromDtoToDataStruct(cart.Items)

	data, err := c.dao.NewCartQuery().CreateCart(ctx, datastruct.Cart{
		UserID: cart.UserID,
		Items:  transformedItem,
	})

	if err != nil {
		return nil, err
	}

	return data, nil

}

func (c *cartService) UpdateCart(ctx context.Context, cart dto.Cart) (*datastruct.Cart, error) {
	mockData := &datastruct.Cart{
		UserID: 1,
		Items: []datastruct.CartItem{
			{
				ProductId:   1,
				ProductName: "Laptop",
				Price:       1000,
				Description: "Dell XPS 15",
				Amounts:     1,
				Image:       "laptop.jpg",
				Status:      "Available",
			},
			{
				ProductId:   2,
				ProductName: "Phone",
				Price:       500,
				Description: "iPhone 12",
				Amounts:     1,
				Image:       "phone.jpg",
				Status:      "Out of Stock",
			},
			{
				ProductId:   3,
				ProductName: "Monitor",
				Price:       200,
				Description: "Samsung 27 inch",
				Amounts:     2,
				Image:       "monitor.jpg",
				Status:      "Available",
			},
		},
	}
	return mockData, nil
}

func (c *cartService) DeleteCart(ctx context.Context, user_id int32) error {
	return nil
}
