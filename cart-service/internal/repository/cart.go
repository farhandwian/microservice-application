package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/farhandwian/microservice/internal/datastruct"
)

type CartQuery interface {
	GetCart(ctx context.Context, key int32) (*datastruct.Cart, error)
	CreateCart(ctx context.Context, cart datastruct.Cart) (*datastruct.Cart, error)
	UpdateCart(ctx context.Context, cart datastruct.Cart) (*datastruct.Cart, error)
	DeleteCart(ctx context.Context, key int32) error
}

type cartQuery struct{}

func (c *cartQuery) GetCart(ctx context.Context, key int32) (*datastruct.Cart, error) {
	data, err := rdb.Get(ctx, strconv.Itoa(int(key))).Bytes()
	if err != nil {
		return nil, err
	}

	var cart datastruct.Cart
	err = json.Unmarshal(data, &cart)
	if err != nil {
		return nil, err
	}

	return &cart, nil
}
func (c *cartQuery) CreateCart(ctx context.Context, cart datastruct.Cart) (*datastruct.Cart, error) {
	key := cart.UserID
	data, err := c.GetCart(ctx, key)

	if data == nil || err != nil {
		data = &cart
	}

	//buat kode untuk append data setelah ngelakuin GetCart
	cart_data, err := json.Marshal(data)

	if err != nil {
		log.Printf("Error marshalling cart: %v", err)
		return nil, err
	}
	rdb, err := RDS() // Get the Redis client and check for errors.
	if err != nil {
		log.Printf("Error initializing Redis: %v", err)
		return nil, err
	}

	rds := rdb.Set(ctx, strconv.Itoa(int(key)), cart_data, 0) // Use the Redis client to set the data.
	if rds.Err() != nil {
		log.Printf("Error adding cart to Redis: %v", rds.Err())
		return nil, rds.Err()
	}
	//-------------------------------
	newData, err := c.GetCart(ctx, key)
	if newData == nil || err != nil {
		newData = &cart
	}
	return newData, nil
}

func (c *cartQuery) UpdateCart(ctx context.Context, cart datastruct.Cart) (*datastruct.Cart, error) {
	key := cart.UserID

	// Check if the cart exists first
	existingCart, err := c.GetCart(ctx, key)
	if err != nil {
		log.Printf("Error fetching cart: %v", err)
		return nil, err
	}

	// If no existing cart, return an error (or you can decide to create one if that's your use-case)
	if existingCart == nil {
		return nil, fmt.Errorf("cart not found")
	}

	// Update with new cart data
	cartData, err := json.Marshal(cart)
	if err != nil {
		log.Printf("Error marshalling cart: %v", err)
		return nil, err
	}

	rdb, err := RDS() // Get the Redis client and check for errors.
	if err != nil {
		log.Printf("Error initializing Redis: %v", err)
		return nil, err
	}
	rds := rdb.Set(ctx, strconv.Itoa(int(key)), cartData, 0)
	if rds.Err() != nil {
		log.Printf("Error updating cart in Redis: %v", rds.Err())
		return nil, rds.Err()
	}

	return &cart, nil
}

func (c *cartQuery) DeleteCart(ctx context.Context, key int32) error {
	rdb, err := RDS() // Get the Redis client and check for errors.
	if err != nil {
		log.Printf("Error initializing Redis: %v", err)
		return err
	}
	// Delete the cart from Redis
	rds := rdb.Del(ctx, strconv.Itoa(int(key)))
	if rds.Err() != nil {
		log.Printf("Error deleting cart from Redis: %v", rds.Err())
		return rds.Err()
	}

	// Optionally, check if the cart was deleted (using the result from the Del command)
	if rds.Val() == 0 {
		return fmt.Errorf("no cart found to delete")
	}

	return nil
}
